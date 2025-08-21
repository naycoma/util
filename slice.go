package util

import "github.com/samber/lo"

// allEqual
func UniformBy[T any, S ~[]T, C comparable](s S, predicate func(T) C) bool {
	if len(s) == 0 {
		return true
	}
	first := predicate(s[0])
	return lo.EveryBy(s[1:], func(v T) bool {
		return predicate(v) == first
	})
}

// []T -> map[T]index
func IndexMap[T comparable, S ~[]T](collection S) map[T]int {
	result := make(map[T]int, len(collection))
	for i, v := range collection {
		result[v] = i
	}
	return result
}
