package filter

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

const NotConditionType = "NotCondition"

func init() {
	registerConditionType(NotConditionType)
	registerExpressionParser("not", parseNot)
}

func parseNot(exp *QueryExpression) (Condition, error) {
	expressions, ok := exp.Value.([]*QueryExpression)
	if ok {
		var conditions []Condition
		for _, e := range expressions {
			c, err := parseQueryExpression(e)
			if err != nil {
				return nil, err
			}
			conditions = append(conditions, c)
		}
		return Not(And(conditions...)), nil
	}
	expression, ok := exp.Value.(*QueryExpression)
	if ok {
		c, err := parseQueryExpression(expression)
		if err != nil {
			return nil, err
		}
		return Not(c), nil
	}
	anyExpressions, ok := exp.Value.([]any)
	if ok {
		var conditions []Condition
		for _, ae := range anyExpressions {
			var e QueryExpression
			err := mapstructure.Decode(ae, &e)
			if err != nil {
				return nil, err
			}
			c, err := parseQueryExpression(&e)
			if err != nil {
				return nil, err
			}
			conditions = append(conditions, c)
		}
		if len(conditions) == 1 {
			return Not(conditions[0]), nil
		} else {
			return Not(And(conditions...)), nil
		}
	}

	return nil, fmt.Errorf("unexpected value type for 'not' expression: %T", exp.Value)
}

// NotCondition negates the condition.
type NotCondition struct {
	Condition Condition
}

// Not creates a new NotCondition.
func Not(c Condition) Condition {
	return &NotCondition{
		Condition: c,
	}
}

// String returns the string representation of the condition.
func (c *NotCondition) String() string {
	return fmt.Sprintf("not ( %s )", c.Condition.String())
}

// Type returns the name of the condition.
func (c *NotCondition) Type() string {
	return NotConditionType
}
