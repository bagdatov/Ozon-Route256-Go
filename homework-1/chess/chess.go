package chess

import (
	"errors"
	"fmt"
)

var (
	errLength    = errors.New("invalid length of argument")
	errLetter    = errors.New("impossible letter in position")
	errNumber    = errors.New("incorrect number in position")
	errDuplicate = errors.New("duplicated position")
)

func CanKnightAttack(white, black string) (bool, error) {
	if err := validatePosition(white); err != nil {
		return false, fmt.Errorf("white: %w", err)
	}

	if err := validatePosition(black); err != nil {
		return false, fmt.Errorf("black: %w", err)
	}

	if white == black {
		return false, errDuplicate
	}

	diff := abs(int(white[0]) - int(black[0]))
	diff2 := abs(int(white[1]) - int(black[1]))

	if (diff == 1 && diff2 == 2) || (diff == 2 && diff2 == 1) {
		return true, nil
	}

	return false, nil
}

func validatePosition(s string) error {
	if len(s) != 2 {
		return errLength
	}

	if !(s[0] >= 'a' && s[0] <= 'h') {
		return errLetter
	}

	if !(s[1] >= '1' && s[1] <= '8') {
		return errNumber
	}

	return nil
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
