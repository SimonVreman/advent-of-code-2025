package util

func Map[T any, R any](slice []T, lambda func(T) R) []R {
	mapped := []R{}

	for _, item := range slice {
		mapped = append(mapped, lambda(item))
	}

	return mapped
}
