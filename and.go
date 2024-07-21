package filter

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

const AndConditionType = "AndCondition"

func init() {
	registerConditionType(AndConditionType)
	registerExpressionParser("and", parseAnd)
}

func parseAnd(exp *QueryExpression) (Condition, error) {
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
			return nil, fmt.Errorf("unexpected 'and' expression value of type %T", exp.Value)
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
	return And(conditions...), nil
}

// AndCondition concatenates multiple Conditions by using the logical "AND".
type AndCondition struct {
	Conditions []Condition
}

// And creates a new Condition by concatenating the conditions with "AND".
func And(c ...Condition) Condition {
	return &AndCondition{Conditions: c}
}

// String returns the string representation of the condition.
func (c *AndCondition) String() string {
	var str string
	for _, condition := range c.Conditions {
		if str != "" {
			str += " and "
		}
		str += condition.String()
	}
	return str
}

// Type returns the name of the condition.
func (c *AndCondition) Type() string {
	return AndConditionType
}
