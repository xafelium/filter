package filter

import "fmt"

const EqualsConditionType = "EqualsCondition"

func init() {
	registerConditionType(EqualsConditionType)
	registerExpressionParser("equals", parseEquals)
	registerExpressionParser("eq", parseEquals)
	registerExpressionParser("==", parseEquals)
	registerExpressionParser("=", parseEquals)
}

func parseEquals(exp *QueryExpression) (Condition, error) {
	return Equals(exp.Field, exp.Value), nil
}

// EqualsCondition compares objects for equality.
type EqualsCondition struct {
	Field string
	Value interface{}
}

// Equals creates a new EqualsCondition.
func Equals(field string, value interface{}) Condition {
	return &EqualsCondition{
		Field: field,
		Value: value,
	}
}

// String returns the string representation of the condition.
func (c *EqualsCondition) String() string {
	return fmt.Sprintf("%s = %v", c.Field, c.Value)
}

// Type returns the name of the condition.
func (c *EqualsCondition) Type() string {
	return EqualsConditionType
}
