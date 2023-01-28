package missingnumbers

// Missing using map to detect rather there was
// a break during iteration from 1 to N.
// Returns two missing numbers of sequence.
func Missing(numbers []int) []int {
	if len(numbers) == 0 {
		return nil
	}

	var (
		// empty struct occupies 0 bytes in storage
		existing = make(map[int]struct{}, len(numbers))
		// always consists of two numbers
		missing = make([]int, 0, 2)
		max     = 0
	)

	for i := 0; i < len(numbers); i++ {

		existing[numbers[i]] = struct{}{}

		if numbers[i] > max {
			max = numbers[i]
		}
	}

	for i := 1; i <= max; i++ {
		if _, found := existing[i]; !found {
			missing = append(missing, i)
		}
	}

	// if there were no breaks in sequence,
	// missing numbers are located after maximum
	for len(missing) < 2 {
		missing = append(missing, max+1)
		max++
	}

	return missing
}
