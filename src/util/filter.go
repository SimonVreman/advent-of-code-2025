package util

func Filter[T any](slice []T, lambda func(T) bool) []T {
	filtered := []T{}

	for _, item := range slice {
		if lambda(item) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}
