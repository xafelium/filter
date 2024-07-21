package filter

import "fmt"

const LowerThanConditionType = "LowerThanCondition"

func init() {
	registerConditionType(LowerThanConditionType)
	registerExpressionParser("lowerThan", parseLowerThan)
	registerExpressionParser("lt", parseLowerThan)
	registerExpressionParser("<", parseLowerThan)
}

func parseLowerThan(exp *QueryExpression) (Condition, error) {
	return LowerThan(exp.Field, exp.Value), nil
}

// LowerThanCondition checks if a field is lower than a values.
type LowerThanCondition struct {
	Field string
	Value interface{}
}

// String returns the string representation of the condition.
func (c *LowerThanCondition) String() string {
	return fmt.Sprintf("%s < %v", c.Field, c.Value)
}

// Type returns the name of the condition.
func (c *LowerThanCondition) Type() string {
	return LowerThanConditionType
}

// LowerThan creates a new LowerThanCondition.
func LowerThan(field string, value interface{}) Condition {
	return &LowerThanCondition{
		Field: field,
		Value: value,
	}
}
