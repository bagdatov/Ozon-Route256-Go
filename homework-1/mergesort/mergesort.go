package mergesort

// MergeSort is used to sort an array of integer
func MergeSort(input []int) []int {
	if len(input) < 2 {
		return input
	}
	a := MergeSort(input[:len(input)/2])
	b := MergeSort(input[len(input)/2:])

	return merge(a, b)
}

func merge(a, b []int) []int {
	final := make([]int, 0, len(a)+len(b))

	var i, j int

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}

	for i < len(a) {
		final = append(final, a[i])
		i++
	}
	for j < len(b) {
		final = append(final, b[j])
		j++
	}
	return final
}
