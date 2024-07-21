package filter

// Filter is a Condition.
type Filter Condition

// Condition contains methods to filter objects.
type Condition interface {
	String() string
	Type() string
}

var (
	conditionTypes []string
	parsers        map[string]ExpressionParser
)

func registerExpressionParser(op string, parser ExpressionParser) {
	if parsers == nil {
		parsers = make(map[string]ExpressionParser)
	}
	parsers[op] = parser
}

func registerConditionType(t string) {
	conditionTypes = append(conditionTypes, t)
}

func AllConditionTypes() []string {
	var types []string
	for _, t := range conditionTypes {
		types = append(types, t)
	}
	return types
}

func UnwrapWhere(c Condition) Condition {
	wc, ok := c.(*WhereCondition)
	if ok {
		return wc.Condition
	}
	return c
}
