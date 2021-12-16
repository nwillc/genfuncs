package genfuncs

func All[T any](slice []T, predicate Predicate[T]) bool {
	for _, e := range slice {
		if !predicate(e) {
			return false
		}
	}
	return true
}

func Any[T any](slice []T, predicate Predicate[T]) bool {
	for _, e := range slice {
		if predicate(e) {
			return true
		}
	}
	return false
}

func Associate[T any, K comparable, V any](slice []T, transform Transform[T, K, V]) map[K]V {
	m := make(map[K]V)
	for _, e := range slice {
		k,v := transform(e)
		m[k] = v
	}
	return m
}

func AssociateBy[T any, K comparable](slice []T, keySelector KeySelector[T, K]) map[K]T {
	m := make(map[K]T)
	for _, e := range slice {
		m[keySelector(e)] = e
	}
	return m
}

func Contains[T comparable](slice []T, element T) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}

func Filter[T any](slice []T, predicate Predicate[T]) []T {
	var results []T
	for _, t := range slice {
		if predicate(t) {
			results = append(results, t)
		}
	}
	return results
}

func Find[T any](slice []T, predicate Predicate[T]) (T, bool) {
	for _, t := range slice {
		if predicate(t) {
			return t, true
		}
	}
	var t T
	return t, false
}

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

func Fold[T, R any](slice []T, initial R, operation Operation[T, R]) R {
	r := initial
	for _, t := range slice {
		r = operation(r, t)
	}
	return r
}
