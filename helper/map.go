package helper

// FetchKeys return extracted keys
func FetchKeys[key comparable, value any](input map[key]value) []key {
	data := make([]key, 0, len(input))

	for k := range input {
		data = append(data, k)
	}
	return data
}
