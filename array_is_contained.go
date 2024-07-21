package filter

import "fmt"

const ArrayIsContainedConditionType = "ArrayIsContainedCondition"

func init() {
	registerConditionType(ArrayIsContainedConditionType)
	registerExpressionParser("arrayIsContained", parseArrayIsContained)
}

func parseArrayIsContained(exp *QueryExpression) (Condition, error) {
	return ArrayIsContained(exp.Field, exp.Value), nil
}

// ArrayIsContainedCondition is an expression to filter values where on array is contained in another array.
type ArrayIsContainedCondition struct {
	Field string
	Value any
}

func (c *ArrayIsContainedCondition) String() string {
	return fmt.Sprintf("%s is contained by %v", c.Field, c.Value)
}

func (c *ArrayIsContainedCondition) Type() string {
	return ArrayIsContainedConditionType
}

func ArrayIsContained(field string, value any) Condition {
	return &ArrayIsContainedCondition{
		Field: field,
		Value: value,
	}
}
