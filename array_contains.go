package filter

import "fmt"

const ArrayContainsConditionType = "ArrayContainsCondition"

func init() {
	registerConditionType(ArrayContainsConditionType)
	registerExpressionParser("arrayContains", parseArrayContains)
}

func parseArrayContains(exp *QueryExpression) (Condition, error) {
	return ArrayContains(exp.Field, exp.Value), nil
}

// ArrayContainsCondition is an expression to filter values contained in an array.
type ArrayContainsCondition struct {
	Field string
	Value any
}

// String returns the string representation of the condition.
func (c *ArrayContainsCondition) String() string {
	return fmt.Sprintf("%s contains element %s", c.Field, c.Value)
}

// Type returns the name of the condition.
func (c *ArrayContainsCondition) Type() string {
	return ArrayContainsConditionType
}

// ArrayContains creates a new ArrayContainsCondition.
func ArrayContains(field string, value any) Condition {
	return &ArrayContainsCondition{
		Field: field,
		Value: value,
	}
}
