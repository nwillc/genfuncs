package genfuncs

import (
	"strings"
)

// All returns true if all elements of slice match the predicate.
func All[T any](slice []T, predicate Predicate[T]) bool {
	for _, e := range slice {
		if !predicate(e) {
			return false
		}
	}
	return true
}

// Any returns true if any element of the slice matches the predicate.
func Any[T any](slice []T, predicate Predicate[T]) bool {
	for _, e := range slice {
		if predicate(e) {
			return true
		}
	}
	return false
}

// Associate returns a map containing key/values provided by transform function applied to elements of the slice.
func Associate[T any, K comparable, V any](slice []T, transform TransformKV[T, K, V]) map[K]V {
	m := make(map[K]V)
	for _, e := range slice {
		k, v := transform(e)
		m[k] = v
	}
	return m
}

// AssociateWith returns a Map where keys are elements from the given sequence and values are produced by the
// valueSelector function applied to each element.
func AssociateWith[K comparable, V any](slice []K, valueSelector ValueSelector[K, V]) map[K]V {
	m := make(map[K]V)
	for _, k := range slice {
		v := valueSelector(k)
		m[k] = v
	}
	return m
}

// Contains returns true if element is found in slice.
func Contains[T comparable](slice []T, element T) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}

// Distinct returns a slice containing only distinct elements from the given slice.
func Distinct[T comparable](slice []T) []T {
	var resultSet []T
	distinctMap := make(map[T]struct{})
	for _, e := range slice {
		if _, ok := distinctMap[e]; ok {
			continue
		}
		distinctMap[e] = struct{}{}
		resultSet = append(resultSet, e)
	}
	return resultSet
}

// Filter returns a slice containing only elements matching the given predicate.
func Filter[T any](slice []T, predicate Predicate[T]) []T {
	var results []T
	for _, t := range slice {
		if predicate(t) {
			results = append(results, t)
		}
	}
	return results
}

// Find returns the first element matching the given predicate and true, or false when no such element was found.
func Find[T any](slice []T, predicate Predicate[T]) (T, bool) {
	for _, t := range slice {
		if predicate(t) {
			return t, true
		}
	}
	var t T
	return t, false
}

// FindLast returns the last element matching the given predicate and true, or false when no such element was found.
func FindLast[T any](slice []T, predicate Predicate[T]) (T, bool) {
	var last T
	var found = false
	for _, t := range slice {
		if predicate(t) {
			found = true
			last = t
		}
	}
	return last, found
}

// FlatMap returns a slice of all elements from results of transform function being invoked on each element of
// original slice.
func FlatMap[T, R any](slice []T, transform Transform[T, R]) []R {
	var results []R
	for _, e := range slice {
		results = append(results, transform(e))
	}
	return results
}

// Fold accumulates a value starting with initial value and applying operation from left to right to current
// accumulated value and each element.
func Fold[T, R any](slice []T, initial R, operation Operation[T, R]) R {
	r := initial
	for _, t := range slice {
		r = operation(r, t)
	}
	return r
}

// GroupBy groups elements of the slice by the key returned by the given keySelector function applied to
// each element and returns a map where each group key is associated with a slice of corresponding elements.
func GroupBy[T any, K comparable](slice []T, keySelector KeySelector[T, K]) map[K][]T {
	m := make(map[K][]T)
	for _, e := range slice {
		k := keySelector(e)
		m[k] = append(m[k], e)
	}
	return m
}

// JoinToString creates a string from all the elements using the stringer on each, separating them using separator, and
// using the given prefix and postfix.
func JoinToString[T any](slice []T, stringer Stringer[T], separator string, prefix string, postfix string) string {
	var sb strings.Builder
	sb.WriteString(prefix)
	last := len(slice) - 1
	for i, e := range slice {
		sb.WriteString(stringer(e))
		if i == last {
			continue
		}
		sb.WriteString(separator)
	}
	sb.WriteString(postfix)
	return sb.String()
}

// Map returns a slice containing the results of applying the given transform function to each element in the original slice.
func Map[T, R any](slice []T, transform Transform[T, R]) []R {
	var results = make([]R, len(slice))
	for i, e := range slice {
		results[i] = transform(e)
	}
	return results
}
