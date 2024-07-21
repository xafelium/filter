package filter_test

import (
	"github.com/stretchr/testify/require"
	"github.com/xafelium/filter"
	"testing"
)

func TestParseQueryExpressions(t *testing.T) {
	tests := []struct {
		name           string
		expressions    []*filter.QueryExpression
		expectedFilter filter.Condition
		errorContains  string
	}{
		{
			name:           "nil as filter",
			expressions:    nil,
			expectedFilter: filter.Where(nil),
		},
		{
			name:           "empty filter",
			expressions:    nil,
			expectedFilter: filter.Where(nil),
		},
		{
			name: "single filter element",
			expressions: []*filter.QueryExpression{
				{
					Field:    "id",
					Operator: "equals",
					Value:    123,
				},
			},
			expectedFilter: filter.Where(
				filter.Equals("id", 123),
			),
		},
		{
			name: "multiple filter elements on top level",
			expressions: []*filter.QueryExpression{
				{
					Field:    "name",
					Operator: "equals",
					Value:    "foo",
				},
				{
					Field:    "isActive",
					Operator: "isNull",
				},
				{
					Field:    "age",
					Operator: "greaterThan",
					Value:    17,
				},
			},
			expectedFilter: filter.Where(
				filter.And(
					filter.Equals("name", "foo"),
					filter.IsNil("isActive"),
					filter.GreaterThan("age", 17),
				),
			),
		},
		{
			name: "single top level and",
			expressions: []*filter.QueryExpression{
				{
					Operator: "and",
					Value: []*filter.QueryExpression{
						{
							Field:    "name",
							Operator: "regex",
							Value:    "foo",
						},
						{
							Field:    "id",
							Operator: "gt",
							Value:    4711,
						},
					},
				},
			},
			expectedFilter: filter.Where(
				filter.And(
					filter.Regex("name", "foo"),
					filter.GreaterThan("id", 4711),
				),
			),
		},
		{
			name: "single top level or",
			expressions: []*filter.QueryExpression{
				{
					Operator: "or",
					Value: []*filter.QueryExpression{
						{
							Field:    "lastName",
							Operator: "notNil",
						},
						{
							Field:    "firstName",
							Operator: "notRegex",
							Value:    "foo",
						},
					},
				},
			},
			expectedFilter: filter.Where(
				filter.Or(
					filter.NotNil("lastName"),
					filter.NotRegex("firstName", "foo"),
				),
			),
		},
		{
			name: "single top level not",
			expressions: []*filter.QueryExpression{
				{
					Operator: "not",
					Value: []*filter.QueryExpression{
						{
							Field:    "name",
							Operator: "==",
							Value:    "tester",
						},
					},
				},
			},
			expectedFilter: filter.Where(
				filter.Not(
					filter.Equals("name", "tester"),
				),
			),
		},
		{
			name: "single top level group",
			expressions: []*filter.QueryExpression{
				{
					Operator: "group",
					Value: []*filter.QueryExpression{
						{
							Field:    "name",
							Operator: ">=",
							Value:    "Test",
						},
					},
				},
			},
			expectedFilter: filter.Where(
				filter.Group(
					filter.GreaterThanOrEqual("name", "Test"),
				),
			),
		},
		{
			name: "complex nested query object",
			expressions: []*filter.QueryExpression{
				{
					Operator: "not",
					Value: []*filter.QueryExpression{
						{
							Field:    "name",
							Operator: "==",
							Value:    "Hans",
						},
						{
							Operator: "or",
							Value: []*filter.QueryExpression{
								{
									Field:    "id",
									Operator: "isNull",
								},
								{
									Field:    "age",
									Operator: ">=",
									Value:    21,
								},
							},
						},
					},
				},
			},
			expectedFilter: filter.Where(
				filter.Not(
					filter.And(
						filter.Equals("name", "Hans"),
						filter.Or(
							filter.IsNil("id"),
							filter.GreaterThanOrEqual("age", 21),
						),
					),
				),
			),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test := test
			actual, err := filter.ParseQueryExpressions(test.expressions)
			if test.errorContains != "" {
				require.ErrorContains(t, err, test.errorContains)
				return
			}
			require.NoError(t, err)
			if test.expectedFilter == nil {
				require.Nil(t, actual)
				return
			}
			require.NotNil(t, actual)
			require.Equal(t, test.expectedFilter.String(), actual.String())
		})
	}

}
