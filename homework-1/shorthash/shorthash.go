package shorthash

// GenerateShortHashes recursively creates unique hash
func GenerateShortHashes(dictionary string, maxLen int) []string {
	res := []string{}

	if len(dictionary) == 0 || maxLen <= 0 {
		// using res instead of nil
		// because tester calls reflect.DeepEqual to compare result
		return res
	}

	for _, v := range dictionary {
		res = append(res, string(v))

		for _, w := range GenerateShortHashes(dictionary, maxLen-1) {
			res = append(res, w+string(v))
		}
	}

	return res
}
