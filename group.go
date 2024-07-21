package filter

import "fmt"

const GroupConditionType = "GroupCondition"
const GroupOp = "group"

func init() {
	registerConditionType(GroupConditionType)
	registerExpressionParser(GroupOp, parseGroup)
}

func parseGroup(exp *QueryExpression) (Condition, error) {
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
		if len(conditions) == 1 {
			return Group(conditions[0]), nil
		}
		return Group(And(conditions...)), nil
	}
	return nil, fmt.Errorf("unexpected value type for group expression: %T", exp.Value)
}

// GroupCondition groups multiple Conditions.
type GroupCondition struct {
	Condition Condition
}

// Group creates a new GroupCondition.
func Group(c Condition) Condition {
	return &GroupCondition{Condition: c}
}

// String returns the string representation of the condition.
func (c *GroupCondition) String() string {
	return "(" + c.Condition.String() + ")"
}

// Type returns the name of the condition.
func (c *GroupCondition) Type() string {
	return GroupConditionType
}
