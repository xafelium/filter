package filter

import "fmt"

const ContainsConditionType = "ContainsCondition"

func init() {
	registerConditionType(ContainsConditionType)
	registerExpressionParser("contains", parseContains)
}

func parseContains(exp *QueryExpression) (Condition, error) {
	v, ok := exp.Value.(string)
	if !ok {
		return nil, fmt.Errorf("contains must have string value but was of type %T", exp.Value)
	}
	return Contains(exp.Field, v), nil
}

// ContainsCondition filters strings containing a value.
type ContainsCondition struct {
	Field string
	Value string
}

// Contains creates a new ContainsCondition.
func Contains(field string, value string) Condition {
	return &ContainsCondition{
		Field: field,
		Value: value,
	}
}

// String returns the string representation of the condition.
func (c *ContainsCondition) String() string {
	return fmt.Sprintf("%s contains %s", c.Field, c.Value)
}

// Type returns the name of the condition.
func (c *ContainsCondition) Type() string {
	return ContainsConditionType
}
