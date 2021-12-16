package genfuncs

func Any[T any](slice []T, predicate Predicate[T]) bool {
	for _, e := range slice {
		if predicate(e) {
			return true
		}
	}
	return false
}

func AssociateBy[T any, K comparable](slice []T, keySelector KeySelector[T,K]) map[K]T {
	m := make(map[K]T)
	for _, e := range slice {
		m[keySelector(e)] = e
	}
	return m
}
