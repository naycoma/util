package util

import "iter"

func Map[T, R any](seq iter.Seq[T], iteratee func(item T) R) iter.Seq[R] {
	return func(yield func(R) bool) {
		for item := range seq {
			if !yield(iteratee(item)) {
				return
			}
		}
	}
}

func Map2To1[K, V, R any](seq iter.Seq2[K, V], iteratee func(k K, v V) R) iter.Seq[R] {
	return func(yield func(R) bool) {
		for k, v := range seq {
			if !yield(iteratee(k, v)) {
				return
			}
		}
	}
}

func Map1To2[T, K, V any](seq iter.Seq[T], iteratee func(item T) (K, V)) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for item := range seq {
			if !yield(iteratee(item)) {
				return
			}
		}
	}
}
func Map2[K1, V1, K2, V2 any](seq iter.Seq2[K1, V1], iteratee func(k K1, v V1) (K2, V2)) iter.Seq2[K2, V2] {
	return func(yield func(K2, V2) bool) {
		for k, v := range seq {
			if !yield(iteratee(k, v)) {
				return
			}
		}
	}
}

func Filter[T any](seq iter.Seq[T], predicate func(item T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for item := range seq {
			if predicate(item) {
				if !yield(item) {
					return
				}
			}
		}
	}
}

func Filter2[K, V any](seq iter.Seq2[K, V], predicate func(k K, v V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if predicate(k, v) {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}
