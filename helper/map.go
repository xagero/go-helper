package helper

// FetchKeys return extracted keys in random order
func FetchKeys[key comparable, value any](input map[key]value) []key {
	data := make([]key, 0, len(input))

	for k := range input {
		data = append(data, k)
	}
	return data
}

// FetchValues return extracted values in random order
func FetchValues[key comparable, value any](input map[key]value) []value {
	data := make([]value, 0, len(input))

	for _, v := range input {
		data = append(data, v)
	}
	return data
}
