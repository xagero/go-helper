package helper

import "golang.org/x/exp/constraints"

// Min returns the smaller value
func Min[T constraints.Ordered](values ...T) T {
	var result = values[0]
	for _, v := range values {
		if v < result {
			result = v
		}
	}
	return result
}

// Max returns the larger value
func Max[T constraints.Ordered](values ...T) T {
	var result = values[0]
	for _, v := range values {
		if v > result {
			result = v
		}
	}
	return result
}
