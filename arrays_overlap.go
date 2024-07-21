package filter

import "fmt"

const ArraysOverlapConditionType = "ArraysOverlapCondition"

func init() {
	registerConditionType(ArraysOverlapConditionType)
	registerExpressionParser("arraysOverlap", parseArraysOverlap)
}

func parseArraysOverlap(exp *QueryExpression) (Condition, error) {
	return ArraysOverlap(exp.Field, exp.Value), nil
}

// ArraysOverlapCondition is an expression to filter values where arrays overlaps, that is, have any elements in common
type ArraysOverlapCondition struct {
	Field string
	Value any
}

func (c *ArraysOverlapCondition) String() string {
	return fmt.Sprintf("%s overlaps array %v", c.Field, c.Value)
}

func (c *ArraysOverlapCondition) Type() string {
	return ArraysOverlapConditionType
}

func ArraysOverlap(field string, value any) Condition {
	return &ArraysOverlapCondition{
		Field: field,
		Value: value,
	}
}
