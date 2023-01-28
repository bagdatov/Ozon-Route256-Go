package anagram

// toLowerLetter is a constant difference between uppercase and lowercase letter.
// According to ASCII table, lowercase letter locates after uppercase.
const toLowerLetter = 'a' - 'A'

// FindAnagrams works with ASCII symbols, and is not meant to use with
// chars that weights more than one byte.
func FindAnagrams(dictionary []string, word string) (result []string) {

	w := toBytes(word)
	if len(w) == 0 || len(dictionary) == 0 {
		return
	}

	for i := 0; i < len(dictionary); i++ {
		if dictionary[i] == word {
			continue
		}

		d := toBytes(dictionary[i])

		if len(w) != len(d) {
			continue
		}

		if anagram(w, d) {
			result = append(result, dictionary[i])
		}

	}

	return
}

// toBytes converts letters to lowercase and removes spaces
func toBytes(s string) []byte {
	bytes := make([]byte, 0, len(s))

	for i := range s {

		if s[i] == ' ' {
			continue
		}

		if s[i] >= 'A' && s[i] <= 'Z' {
			bytes = append(bytes, s[i]+toLowerLetter)
			continue
		}

		bytes = append(bytes, s[i])

	}
	return bytes
}

// anagram is comparing two slices of byte.
// Original word is not considered as anagram.
func anagram(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	if string(a) == string(b) {
		return false
	}

	foo := make(map[byte]int64, len(a))

	// increase and decrease each entrance of letter,
	// so values for each key should be zero
	for i := 0; i < len(a); i++ {
		foo[a[i]]++
		foo[b[i]]--
	}

	for key := range foo {
		if foo[key] != 0 {
			return false
		}
	}

	return true
}
