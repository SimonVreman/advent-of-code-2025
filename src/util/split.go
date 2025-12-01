package util

func Split[T comparable](slice []T, separator T) [][]T {
	slices := [][]T{}
	index := 0

	for i, v := range slice {
		if v == separator {
			slices = append(slices, slice[index:i])
			index = i + 1
		}
	}

	if index < (len(slice)) {
		slices = append(slices, slice[index:])
	}

	return slices
}
