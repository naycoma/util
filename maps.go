package util

import (
	"iter"
	"maps"
)

func AllBy[S ~[]V, K comparable, V any](s S, toKey func(v V) K) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, v := range s {
			if !yield(toKey(v), v) {
				return
			}
		}
	}
}

// Merge multiple maps into a single map. Keys in later maps override keys in earlier maps.
// If base is nil, a new map is created.
func Merge[K comparable, V any, M ~map[K]V](base M, override ...M) M {
	if base == nil {
		base = make(M)
	}
	merged := maps.Clone(base)
	for _, o := range override {
		maps.Insert(merged, maps.All(o))
	}
	return merged
}

// MergeFromSlice merges a slice of values into a map, using a keyFunc to extract keys from values.
// If base is nil, a new map is created.
// Values in the slice override existing values in base if their keys are the same.
func MergeFromSlice[M ~map[K]V, S ~[]V, K comparable, V any](base M, toKey func(v V) K, override ...S) M {
	if base == nil {
		base = make(M)
	}
	merged := maps.Clone(base)
	for _, o := range override {
		maps.Insert(merged, AllBy(o, toKey))
	}
	return merged
}

// SliceToIndexMap converts a slice into a map where keys are slice elements and values are their original indices.
// If there are duplicate elements, the last occurrence's index will be stored.
func SliceToIndexMap[T comparable, S ~[]T](slice S) map[T]int {
	result := make(map[T]int, len(slice))

	for i, v := range slice {
		result[v] = i
	}

	return result
}

// SliceToIndexMapBy converts a slice of values into a map where keys are transformed from slice elements
// using a transform function, and values are their original indices.
// If there are duplicate transformed keys, the last occurrence's index will be stored.
func SliceToIndexMapBy[K comparable, V any, S ~[]V](slice S, toKey func(item V) K) map[K]int {
	result := make(map[K]int, len(slice))

	for i, v := range slice {
		result[toKey(v)] = i
	}

	return result
}

// FilterMapToSlice filters and transforms map elements into a new slice.
// The iteratee function determines whether an element is included and how it's transformed.
func FilterMapToSlice[K comparable, V, R any, M ~map[K]V](in M, iteratee func(key K, value V) (R, bool)) []R {
	result := []R{}
	for k, v := range in {
		if r, ok := iteratee(k, v); ok {
			result = append(result, r)
		}
	}
	return result
}
