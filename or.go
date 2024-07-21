package filter

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

const OrConditionType = "OrCondition"

func init() {
	registerConditionType(OrConditionType)
	registerExpressionParser("or", parseOr)
}

func parseOr(exp *QueryExpression) (Condition, error) {
	expressions, ok := exp.Value.([]*QueryExpression)
	if !ok {
		anyExpressions, ok := exp.Value.([]any)
		if ok {
			for _, ae := range anyExpressions {
				var e QueryExpression
				err := mapstructure.Decode(ae, &e)
				if err != nil {
					return nil, err
				}
				expressions = append(expressions, &e)
			}
		} else {
			return nil, fmt.Errorf("unexpected 'or' expression value of type %T", exp.Value)
		}
	}
	var conditions []Condition
	for _, e := range expressions {
		c, err := parseQueryExpression(e)
		if err != nil {
			return nil, err
		}
		conditions = append(conditions, c)
	}
	return Or(conditions...), nil
}

// OrCondition concatenates multiple Conditions by using the logical "OR".
type OrCondition struct {
	Conditions []Condition
}

// Or creates a new OrCondition.
func Or(c ...Condition) Condition {
	return &OrCondition{
		Conditions: c,
	}
}

// String returns the string representation of the condition.
func (c *OrCondition) String() string {
	var str string
	for _, condition := range c.Conditions {
		if str != "" {
			str += " or "
		}
		str += condition.String()
	}
	return str
}

// Type returns the name of the condition.
func (c *OrCondition) Type() string {
	return OrConditionType
}
