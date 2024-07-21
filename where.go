package filter

const WhereConditionType = "WhereCondition"

func init() {
	registerConditionType(WhereConditionType)
}

// WhereCondition is the "WHERE" clause. It can be used only once and must be the first Condition.
type WhereCondition struct {
	Condition Condition
}

// Where creates a new WhereCondition.
func Where(condition Condition) Condition {
	return &WhereCondition{
		Condition: condition,
	}
}

// String returns the string representation of the condition.
func (c WhereCondition) String() string {
	if c.Condition == nil {
		return ""
	}
	return "where " + c.Condition.String()
}

// Type returns the name of the condition.
func (c WhereCondition) Type() string {
	return WhereConditionType
}
