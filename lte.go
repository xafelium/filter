package filter

import "fmt"

const LowerThanOrEqualConditionType = "LowerThanOrEqualCondition"

func init() {
	registerConditionType(LowerThanOrEqualConditionType)
	registerExpressionParser("lowerOrEqual", parseLowerThanOrEqual)
	registerExpressionParser("lte", parseLowerThanOrEqual)
	registerExpressionParser("<=", parseLowerThanOrEqual)
}

func parseLowerThanOrEqual(exp *QueryExpression) (Condition, error) {
	return LowerThanOrEqual(exp.Field, exp.Value), nil
}

// LowerThanOrEqualCondition checks if a field is lower or equal a value.
type LowerThanOrEqualCondition struct {
	Field string
	Value interface{}
}

// String returns the string representation of the condition.
func (c *LowerThanOrEqualCondition) String() string {
	return fmt.Sprintf("%s <= %v", c.Field, c.Value)
}

// Type returns the name of the condition.
func (c *LowerThanOrEqualCondition) Type() string {
	return LowerThanOrEqualConditionType
}

// LowerThanOrEqual creates a new LowerThanOrEqualCondition.
func LowerThanOrEqual(field string, value interface{}) Condition {
	return &LowerThanOrEqualCondition{
		Field: field,
		Value: value,
	}
}
