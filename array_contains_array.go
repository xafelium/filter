package filter

import "fmt"

const ArrayContainsArrayConditionType = "ArrayContainsArrayCondition"

func init() {
	registerConditionType(ArrayContainsArrayConditionType)
	registerExpressionParser("arrayContainsArray", parseArrayContainsArray)
}

func parseArrayContainsArray(exp *QueryExpression) (Condition, error) {
	return ArrayContainsArray(exp.Field, exp.Value), nil
}

// ArrayContainsArrayCondition is an expression to filter values where one array contains another array.
type ArrayContainsArrayCondition struct {
	Field string
	Value any
}

func (c *ArrayContainsArrayCondition) String() string {
	return fmt.Sprintf("%s contains %v", c.Field, c.Value)
}

func (c *ArrayContainsArrayCondition) Type() string {
	return ArrayContainsArrayConditionType
}

func ArrayContainsArray(field string, value any) Condition {
	return &ArrayContainsArrayCondition{
		Field: field,
		Value: value,
	}
}
