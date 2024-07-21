package filter

import (
	"fmt"
)

const NotEqualsConditionType = "NotEqualsCondition"

func init() {
	registerConditionType(NotEqualsConditionType)
	registerExpressionParser("notEqual", parseNotEqual)
	registerExpressionParser("ne", parseNotEqual)
	registerExpressionParser("!=", parseNotEqual)
	registerExpressionParser("<>", parseNotEqual)
}

func parseNotEqual(exp *QueryExpression) (Condition, error) {
	return NotEquals(exp.Field, exp.Value), nil
}

func NotEquals(field string, value any) Condition {
	return &NotEqualsCondition{
		Field: field,
		Value: value,
	}
}

type NotEqualsCondition struct {
	Field string
	Value any
}

func (c *NotEqualsCondition) String() string {
	return fmt.Sprintf("%s != %v", c.Field, c.Value)
}

func (c *NotEqualsCondition) Type() string {
	return NotEqualsConditionType
}
