package filter

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFieldPathsHelper_OneOf(t *testing.T) {
	var paths []string
	qph := NewFieldPathsHelper(paths)
	require.False(t, qph.OneOf("test"))

	paths = []string{"test"}
	qph = NewFieldPathsHelper(paths)
	require.True(t, qph.OneOf("test", "foo"))

	paths = []string{"foo", "bar", "baz"}
	qph = NewFieldPathsHelper(paths)
	require.False(t, qph.OneOf("test"))

	paths = []string{"foo", "bar", "baz"}
	qph = NewFieldPathsHelper(paths)
	require.True(t, qph.OneOf("bar", "foo"))
}

func TestFieldPathsHelper_All(t *testing.T) {
	var paths []string
	qph := NewFieldPathsHelper(paths)
	require.False(t, qph.All())

	paths = []string{"foo"}
	qph = NewFieldPathsHelper(paths)
	require.True(t, qph.All("foo"))

	paths = []string{"foo"}
	qph = NewFieldPathsHelper(paths)
	require.False(t, qph.All("bar"))

	paths = []string{"foo", "bar"}
	qph = NewFieldPathsHelper(paths)
	require.True(t, qph.All("foo", "bar"))

	paths = []string{"foo", "bar", "baz"}
	qph = NewFieldPathsHelper(paths)
	require.True(t, qph.All("foo", "bar"))

	paths = []string{"foo", "bar", "baz"}
	qph = NewFieldPathsHelper(paths)
	require.False(t, qph.All("foo", "bar", "test"))

	paths = []string{"foo", "bar", "baz"}
	qph = NewFieldPathsHelper(paths)
	require.True(t, qph.All("foo", "bar", "baz"))
}

func TestPrependPath(t *testing.T) {
	var fields []string

	require.Empty(t, PrependPath(fields, "test"))

	fields = []string{"a"}
	require.Equal(t, []string{"a"}, PrependPath(fields, ""))

	fields = []string{"a"}
	require.Equal(t, []string{"x.a"}, PrependPath(fields, "x"))

	fields = []string{"a"}
	require.Equal(t, []string{"x.a"}, PrependPath(fields, "x."))

	fields = []string{"a"}
	require.Equal(t, []string{"x.y.z.a"}, PrependPath(fields, "x.y.z"))

	fields = []string{"a"}
	require.Equal(t, []string{"x.y.z.a"}, PrependPath(fields, "x.y.z."))

	fields = []string{"a", "b"}
	require.Equal(t, []string{"z.a", "z.b"}, PrependPath(fields, "z"))

	fields = []string{"a", "b"}
	require.Equal(t, []string{"z.a", "z.b"}, PrependPath(fields, "z."))

	fields = []string{"a", "b"}
	require.Equal(t, []string{"y.z.a", "y.z.b"}, PrependPath(fields, "y.z"))

	fields = []string{"a", "b"}
	require.Equal(t, []string{"y.z.a", "y.z.b"}, PrependPath(fields, "y.z."))
}

func TestUniqueFields(t *testing.T) {
	var fields []string

	require.Empty(t, UniqueFields(fields))

	fields = []string{"a"}
	require.Equal(t, []string{"a"}, UniqueFields(fields))

	fields = []string{"a", "b"}
	require.Equal(t, []string{"a", "b"}, UniqueFields(fields))

	fields = []string{"a"}
	require.Equal(t, []string{"a", "b"}, UniqueFields(fields, "b"))

	fields = []string{"b"}
	require.Equal(t, []string{"a", "b"}, UniqueFields(fields, "a"))

	fields = []string{}
	require.Equal(t, []string{"d", "e"}, UniqueFields(fields, "d", "e"))
}

func TestJoinFields(t *testing.T) {
	tests := []struct {
		name     string
		fields   []string
		expected string
	}{
		{
			name:     "no fields",
			fields:   []string{},
			expected: "",
		},
		{
			name:     "one field",
			fields:   []string{"a"},
			expected: "a",
		},
		{
			name:     "two fields",
			fields:   []string{"z", "y"},
			expected: "z.y",
		},
		{
			name:     "three fields",
			fields:   []string{"foo", "bar", "baz"},
			expected: "foo.bar.baz",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expected, JoinFields(test.fields...))
		})
	}
}
