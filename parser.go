package filter

import "fmt"

type QueryExpression struct {
	Field    string `json:"field" yaml:"field"`
	Operator string `json:"op" mapstructure:"op" yaml:"op"`
	Value    any    `json:"value" yaml:"value"`
}

type ExpressionParser func(exp *QueryExpression) (Condition, error)

func ParseQueryExpressions(expressions []*QueryExpression) (Condition, error) {
	return parseQueryExpressions(expressions)
}

func parseQueryExpressions(expressions []*QueryExpression) (Condition, error) {
	if len(expressions) == 0 {
		return Where(nil), nil
	}

	var conditions []Condition
	for _, exp := range expressions {
		c, err := parseQueryExpression(exp)
		if err != nil {
			return nil, err
		}
		conditions = append(conditions, c)
	}
	if len(conditions) == 0 {
		return Where(nil), nil
	}
	if len(conditions) == 1 {
		return Where(conditions[0]), nil
	}
	return Where(
		And(conditions...),
	), nil
}

func parseQueryExpression(exp *QueryExpression) (Condition, error) {
	parse, found := parsers[exp.Operator]
	if !found {
		return nil, fmt.Errorf("unknown operator: %s", exp.Operator)
	}
	return parse(exp)
}
