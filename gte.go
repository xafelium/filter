package filter

import "fmt"

const GreaterThanOrEqualConditionType = "GreaterThanOrEqualCondition"

func init() {
	registerConditionType(GreaterThanOrEqualConditionType)
	registerExpressionParser("greaterOrEqual", parseGreaterThanOrEqual)
	registerExpressionParser("gte", parseGreaterThanOrEqual)
	registerExpressionParser(">=", parseGreaterThanOrEqual)
}

func parseGreaterThanOrEqual(exp *QueryExpression) (Condition, error) {
	return GreaterThanOrEqual(exp.Field, exp.Value), nil
}

// GreaterThanOrEqualCondition checks if a field is greater or equal a value.
type GreaterThanOrEqualCondition struct {
	Field string
	Value interface{}
}

// String returns the string representation of the condition.
func (c *GreaterThanOrEqualCondition) String() string {
	return fmt.Sprintf("%s >= %v", c.Field, c.Value)
}

// Type returns the name of the condition.
func (c *GreaterThanOrEqualCondition) Type() string {
	return GreaterThanOrEqualConditionType
}

// GreaterThanOrEqual creates a new GreaterThanOrEqualCondition.
func GreaterThanOrEqual(field string, value interface{}) Condition {
	return &GreaterThanOrEqualCondition{
		Field: field,
		Value: value,
	}
}
