package secretmessage

import (
	"sort"
	"strings"
)

const (
	// alphabet and '_' symbol
	alphaSize = 27
	// represent index of '_' symbol
	divider = 26
)

// Decode works only with english alphabet.
// Passing non english letter (except '_')
// will cause panic (slice bounds out of range).
func Decode(encoded string) string {
	if encoded == "" {
		return ""
	}

	var (
		encLower = strings.ToLower(encoded)
		// how many times each letter appeared in the string
		count = make([]int, alphaSize)
	)

	// (letter - 'a') is equal to index inside count slice
	for _, letter := range encLower {
		if letter == '_' {
			count[divider]++
			continue
		}
		count[letter-'a']++
	}

	var (
		// remove '_' and letters that
		// appears rarer than '_'
		runes = make([]rune, 0, alphaSize-1)
	)

	for i, v := range count {
		if i != divider && v >= count[divider] {
			runes = append(runes, rune('a'+i))
		}
	}

	// sort by count descending
	sort.Slice(runes, func(i, j int) bool {
		return count[runes[i]-'a'] > count[runes[j]-'a']
	})

	return string(runes)
}
