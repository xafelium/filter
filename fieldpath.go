package filter

import (
	"fmt"
	"sort"
	"strings"
)

type FieldPathsHelper interface {
	Empty() bool
	OneOfOrEmpty(values ...string) bool
	OneOf(values ...string) bool
	AllOrEmpty(values ...string) bool
	All(values ...string) bool
	GetSubPath(path string) FieldPathsHelper
	GetFields() []string
}

type fieldPathsHelper struct {
	paths map[string]struct{}
}

func (h *fieldPathsHelper) GetFields() []string {
	var fields []string
	for path := range h.paths {
		fields = append(fields, path)
	}
	return fields
}

func (h *fieldPathsHelper) Empty() bool {
	return len(h.paths) == 0
}

func (h *fieldPathsHelper) OneOfOrEmpty(values ...string) bool {
	return h.Empty() || h.OneOf(values...)
}

func (h *fieldPathsHelper) OneOf(values ...string) bool {
	if len(h.paths) == 0 {
		return false
	}
	for _, v := range values {
		_, found := h.paths[v]
		if found {
			return true
		}
		for path := range h.paths {
			if strings.HasPrefix(path, fmt.Sprintf("%s.", v)) {
				return true
			}
		}
	}
	return false
}

func (h *fieldPathsHelper) AllOrEmpty(values ...string) bool {
	return h.Empty() || h.All(values...)
}

func (h *fieldPathsHelper) All(values ...string) bool {
	if len(h.paths) == 0 {
		return false
	}
	for _, v := range values {
		_, found := h.paths[v]
		if !found {
			return false
		}
	}
	return true
}

func (h *fieldPathsHelper) GetSubPath(path string) FieldPathsHelper {
	var paths []string
	for p := range h.paths {
		prefix := fmt.Sprintf("%s.", path)
		if strings.HasPrefix(p, prefix) {
			paths = append(paths, strings.Replace(p, prefix, "", 1))
		}
	}
	return NewFieldPathsHelper(paths)
}

func NewFieldPathsHelper(paths []string) FieldPathsHelper {
	pathsMap := make(map[string]struct{}, len(paths))
	for _, path := range paths {
		pathsMap[path] = struct{}{}
	}
	return &fieldPathsHelper{
		paths: pathsMap,
	}
}

func PrependPath(fields []string, pathPrefix string) []string {
	prefixedPaths := make([]string, len(fields))
	if pathPrefix != "" && !strings.HasSuffix(pathPrefix, ".") {
		pathPrefix += "."
	}
	for i, field := range fields {
		prefixedPaths[i] = fmt.Sprintf("%s%s", pathPrefix, field)
	}
	return prefixedPaths
}

func UniqueFields(fields []string, additionalFields ...string) []string {
	fieldMap := make(map[string]struct{})
	for _, field := range fields {
		fieldMap[field] = struct{}{}
	}
	for _, field := range additionalFields {
		fieldMap[field] = struct{}{}
	}
	var uniqueFields []string
	for field := range fieldMap {
		uniqueFields = append(uniqueFields, field)
	}
	sort.Strings(uniqueFields)
	return uniqueFields
}

func UniqueSortedFields(fields []string, additionalFields ...string) []string {
	u := UniqueFields(fields, additionalFields...)
	sort.Slice(u, func(i, j int) bool {
		return u[i] < u[j]
	})
	return u
}

func JoinFields(fields ...string) string {
	return strings.Join(fields, ".")
}
