package slicemanip

func Reverse[T any](input []T) []T {
	output := make([]T, len(input))

	for i, val := range input {
		output[len(input)-1-i] = val
	}

	return output
}