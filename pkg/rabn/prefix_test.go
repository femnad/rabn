package rabn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommonPrefixOfTwoStrings(t *testing.T) {
	prefix := commonPrefix("bar", "baz")
	assert.Equal(t, "ba", prefix, "incorrect common prefix")
}

func TestCommonPrefix(t *testing.T) {
	prefix := findLongestCommonPrefix([]string{"bar", "baz"})
	assert.Equal(t, "ba", prefix, "incorrect common prefix")

	prefix = findLongestCommonPrefix([]string{"bar", "baz", "foo"})
	assert.Equal(t, "", prefix, "incorrect common prefix")

	prefix = findLongestCommonPrefix([]string{"bazr", "baz"})
	assert.Equal(t, "baz", prefix, "incorrect common prefix")
}
