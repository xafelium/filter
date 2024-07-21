package filter

import "fmt"

const OverlapsConditionType = "OverlapsCondition"

func init() {
	registerConditionType(OverlapsConditionType)
	registerExpressionParser("overlaps", parseOverlaps)
}

func parseOverlaps(exp *QueryExpression) (Condition, error) {
	return Overlaps(exp.Field, exp.Value), nil
}

// OverlapsCondition checks if two time ranges overlap.
type OverlapsCondition struct {
	Field string
	Value interface{}
}

// String returns the string representation of the condition.
func (c *OverlapsCondition) String() string {
	return fmt.Sprintf("%s overlap %v", c.Field, c.Value)
}

// Type returns the name of the condition.
func (c *OverlapsCondition) Type() string {
	return OverlapsConditionType
}

// Overlaps creates a new OverlapsCondition.
func Overlaps(field string, value interface{}) Condition {
	return &OverlapsCondition{
		Field: field,
		Value: value,
	}
}
