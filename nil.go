package filter

import "fmt"

const (
	IsNilConditionType = "IsNilCondition"

	NotNilConditionType = "NotNilCondition"
)

func init() {
	registerConditionType(IsNilConditionType)
	registerExpressionParser("isNull", parseIsNil)
	registerExpressionParser("null", parseIsNil)
	registerExpressionParser("isNil", parseIsNil)
	registerExpressionParser("nil", parseIsNil)
	registerConditionType(NotNilConditionType)
	registerExpressionParser("isNotNull", parseNotNil)
	registerExpressionParser("notNull", parseNotNil)
	registerExpressionParser("isNotNil", parseNotNil)
	registerExpressionParser("notNil", parseNotNil)
}

func parseNotNil(exp *QueryExpression) (Condition, error) {
	return NotNil(exp.Field), nil
}

func parseIsNil(exp *QueryExpression) (Condition, error) {
	return IsNil(exp.Field), nil
}

type IsNilCondition struct {
	Field string
}

func (c *IsNilCondition) String() string {
	return fmt.Sprintf("%s IS NULL", c.Field)
}

func (c *IsNilCondition) Type() string {
	return IsNilConditionType
}

func IsNil(field string) Condition {
	return &IsNilCondition{
		Field: field,
	}
}

type NotNilCondition struct {
	Field string
}

func (c *NotNilCondition) String() string {
	return fmt.Sprintf("%s IS NOT NULL", c.Field)
}

func (c *NotNilCondition) Type() string {
	return NotNilConditionType
}

func NotNil(field string) Condition {
	return &NotNilCondition{
		Field: field,
	}
}
