package filter

import "fmt"

const NotRegexConditionType = "NotRegexCondition"

func init() {
	registerConditionType(NotRegexConditionType)
	registerExpressionParser("notRegex", parseNotRegex)
}

func parseNotRegex(exp *QueryExpression) (Condition, error) {
	v, ok := exp.Value.(string)
	if !ok {
		return nil, fmt.Errorf("notRegex must have string value but was of type %T", exp.Value)
	}
	return NotRegex(exp.Field, v), nil
}

func NotRegex(field string, pattern string) Condition {
	return &NotRegexCondition{
		Field:      field,
		Expression: pattern,
	}
}

type NotRegexCondition struct {
	Field      string
	Expression string
}

func (c *NotRegexCondition) String() string {
	return fmt.Sprintf("%s matchesNotRegex(%s)", c.Field, c.Expression)
}

func (c *NotRegexCondition) Type() string {
	return NotRegexConditionType
}
