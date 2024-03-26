package helper

// FetchKeys return extracted keys in random order
func FetchKeys[key comparable, value any](input map[key]value) []key {
	result := make([]key, 0, len(input))

	for k := range input {
		result = append(result, k)
	}
	return result
}

// FetchValues return extracted values in random order
func FetchValues[key comparable, value any](input map[key]value) []value {
	result := make([]value, 0, len(input))

	for _, v := range input {
		result = append(result, v)
	}
	return result
}

// FilterByFunc returns extracted values filtered by callback function.
func FilterByFunc[value any](input []value, fn func(value) bool) []value {
	result := make([]value, 0)

	for _, v := range input {
		if fn(v) {
			result = append(result, v)
		}
	}

	return result
}
