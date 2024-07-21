package filter

import "fmt"

const GreaterThanConditionType = "GreaterThanCondition"

func init() {
	registerConditionType(GreaterThanConditionType)
	registerExpressionParser("greaterThan", parseGreaterThan)
	registerExpressionParser("gt", parseGreaterThan)
	registerExpressionParser(">", parseGreaterThan)
}

func parseGreaterThan(exp *QueryExpression) (Condition, error) {
	return GreaterThan(exp.Field, exp.Value), nil
}

// GreaterThanCondition checks if a field is greater a value.
type GreaterThanCondition struct {
	Field string
	Value interface{}
}

// String returns the string representation of the condition.
func (c *GreaterThanCondition) String() string {
	return fmt.Sprintf("%s > %v", c.Field, c.Value)
}

// Type returns the name of the condition.
func (c *GreaterThanCondition) Type() string {
	return GreaterThanConditionType
}

// GreaterThan creates a new GreaterThanCondition.
func GreaterThan(field string, value interface{}) Condition {
	return &GreaterThanCondition{
		Field: field,
		Value: value,
	}
}
