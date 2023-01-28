package jaro

import (
	"strings"
	"unicode/utf8"
)

func Distance(word1 string, word2 string) float64 {
	word1 = strings.ToLower(word1)
	word2 = strings.ToLower(word2)

	if word1 == word2 {
		return 1
	}

	l1 := utf8.RuneCountInString(word1)
	l2 := utf8.RuneCountInString(word2)

	matchLen := matchLength(l1, l2)

	// m - matches, t - transpositions
	m, t := count([]rune(word1), []rune(word2), matchLen)

	if m <= 0 {
		return 0
	}

	return float64((l1+l2)*m*m+l1*l2*(m-t)) / float64(3*l1*l2*m)
}

// matchLength defines maximum allowed distance for letters
// that considered to be matching.
func matchLength(l1, l2 int) int {
	max := l1
	if l2 > max {
		max = l2
	}
	return (max / 2) - 1
}

// count returns number of matches and half of the transpositions
func count(word1, word2 []rune, matchLen int) (matches int, transpositions int) {

	word1Matches := make([]bool, len(word1))
	word2Matches := make([]bool, len(word2))

	// count matches and mark them
	for i := range word1 {
		start := i - matchLen

		if start < 0 {
			start = 0
		}

		end := i + matchLen + 1
		if end > len(word2) {
			end = len(word2)
		}

		for j := start; j < end; j++ {
			// skip if used already
			if word2Matches[j] {
				continue
			}

			if word1[i] != word2[j] {
				continue
			}
			word1Matches[i] = true
			word2Matches[j] = true
			matches++
			break
		}
	}

	// count transpositions
	j := 0
	for i := range word1 {
		if !word1Matches[i] {
			continue
		}

		// move to the closest letter
		// that has a pair
		for !word2Matches[j] {
			j++
		}

		// it shouldn't be exact match
		if word1[i] != word2[j] {
			transpositions++
		}
		j++
	}

	return matches, transpositions / 2
}
