package functional

func Map[T any, R any](input []T, f func(T) R) []R {
	output := make([]R, len(input))

	for i, val := range input {
		output[i] = f(val)
	}

	return output
}

func Filter[T any](input []T, predicate func(T) bool) []T {
	var output []T

	for _, val := range input {
		if predicate(val) {
			output = append(output, val)
		}
	}

	return output
}

