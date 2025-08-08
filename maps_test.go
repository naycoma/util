package util_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/naycoma/util"
)

func TestMerge(t *testing.T) {
	a := assert.New(t)

	base := map[string]int{"a": 1, "b": 2}
	override1 := map[string]int{"b": 3, "c": 4}
	override2 := map[string]int{"c": 5, "d": 6}

	// Test basic merge
	result := util.Merge(base, override1)
	a.Equal(map[string]int{"a": 1, "b": 3, "c": 4}, result)

	// Test multiple overrides
	result = util.Merge(base, override1, override2)
	a.Equal(map[string]int{"a": 1, "b": 3, "c": 5, "d": 6}, result)

	// Test with nil base
	result = util.Merge(nil, override1)
	a.Equal(map[string]int{"b": 3, "c": 4}, result)

	// Test with empty override
	result = util.Merge(base)
	a.Equal(map[string]int{"a": 1, "b": 2}, result)
}

type Item struct {
	Name  string
	Value int
}

func TestMergeFromSlice(t *testing.T) {
	a := assert.New(t)

	base := map[string]Item{"a": {"a", 1}, "b": {"b", 2}}
	override1 := []Item{{"b", 3}, {"c", 4}}
	override2 := []Item{{"c", 5}, {"d", 6}}
	keyFunc := func(v Item) string { return v.Name }

	// Test basic merge
	result := util.MergeFromSlice(base, keyFunc, override1)
	a.Equal(map[string]Item{"a": {"a", 1}, "b": {"b", 3}, "c": {"c", 4}}, result)

	// Test multiple overrides
	result = util.MergeFromSlice(base, keyFunc, override1, override2)
	a.Equal(map[string]Item{"a": {"a", 1}, "b": {"b", 3}, "c": {"c", 5}, "d": {"d", 6}}, result)

	// Test with nil prev
	result = util.MergeFromSlice((map[string]Item)(nil), keyFunc, override1)
	a.Equal(map[string]Item{"b": {"b", 3}, "c": {"c", 4}}, result)

	// Test with empty next
	result = util.MergeFromSlice(base, keyFunc, []Item{})
	a.Equal(map[string]Item{"a": {"a", 1}, "b": {"b", 2}}, result)
}

func TestSliceToIndexMap(t *testing.T) {
	a := assert.New(t)

	slice := []string{"a", "b", "c", "a"}
	result := util.SliceToIndexMap(slice)
	a.Equal(map[string]int{"a": 3, "b": 1, "c": 2}, result)

	// Test with empty slice
	result = util.SliceToIndexMap([]string{})
	a.Equal(map[string]int{}, result)
}

func TestSliceToIndexMapBy(t *testing.T) {
	a := assert.New(t)

	type Item struct {
		ID   int
		Name string
	}
	slice := []Item{{1, "a"}, {2, "b"}, {3, "c"}, {1, "d"}}
	transform := func(item Item) int { return item.ID }
	result := util.SliceToIndexMapBy(slice, transform)
	a.Equal(map[int]int{1: 3, 2: 1, 3: 2}, result)

	// Test with empty slice
	result = util.SliceToIndexMapBy([]Item{}, transform)
	a.Equal(map[int]int{}, result)
}

func TestFilterMapToSlice(t *testing.T) {
	a := assert.New(t)

	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	iteratee := func(key string, value int) (string, bool) {
		if value%2 == 0 {
			return key + "_even", true
		}
		return "", false
	}
	result := util.FilterMapToSlice(m, iteratee)
	// Order is not guaranteed for map iteration, so sort the result
	sortedResult := result[:]
	// Sort the expected slice as well
	expected := []string{"b_even", "d_even"}
	// Sort both slices before comparison
	sort.Strings(sortedResult)
	sort.Strings(expected)
	a.ElementsMatch(expected, sortedResult)

	// Test with empty map
	result = util.FilterMapToSlice(map[string]int{}, iteratee)
	a.Equal([]string{}, result)
}
