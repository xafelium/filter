package filter

import "fmt"

const RegexConditionType = "RegexCondition"

func init() {
	registerConditionType(RegexConditionType)
	registerExpressionParser("regex", parseRegex)
}

func parseRegex(exp *QueryExpression) (Condition, error) {
	v, ok := exp.Value.(string)
	if !ok {
		return nil, fmt.Errorf("regex must have string value but was of type %T", exp.Value)
	}
	return Regex(exp.Field, v), nil
}

func Regex(field string, expression string) Condition {
	return &RegexCondition{
		Field:      field,
		Expression: expression,
	}
}

type RegexCondition struct {
	Field      string
	Expression string
}

func (c *RegexCondition) String() string {
	return fmt.Sprintf("%s matchesRegex(%s)", c.Field, c.Expression)
}

func (c *RegexCondition) Type() string {
	return RegexConditionType
}
