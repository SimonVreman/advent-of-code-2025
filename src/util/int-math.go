package util

import "math"

func Max(numbers ...int) int {
	max := 0
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func Min(numbers ...int) int {
	min := math.MaxInt
	for _, number := range numbers {
		if number < min {
			min = number
		}
	}
	return min
}
