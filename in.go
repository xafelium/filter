package filter

import "fmt"

const InConditionType = "InCondition"

func init() {
	registerConditionType(InConditionType)
	registerExpressionParser("in", parseIn)
}

func parseIn(exp *QueryExpression) (Condition, error) {
	return In(exp.Field, exp.Value), nil
}

// InCondition checks if a list of values is in a field.
type InCondition struct {
	Field string
	Value interface{}
}

// String returns the string representation of the condition.
func (c *InCondition) String() string {
	return fmt.Sprintf("%s IN (%v)", c.Field, c.Value)
}

// Type returns the name of the condition.
func (c *InCondition) Type() string {
	return InConditionType
}

// In creates a new InCondition.
func In(field string, value interface{}) Condition {
	return &InCondition{
		Field: field,
		Value: value,
	}
}
