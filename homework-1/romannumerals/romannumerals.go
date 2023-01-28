package romannumerals

import "strings"

type foo struct {
	num   int
	roman string
}

var bar = []foo{
	{1, "I"},
	{4, "IV"},
	{5, "V"},
	{9, "IX"},
	{10, "X"},
	{40, "XL"},
	{50, "L"},
	{90, "XC"},
	{100, "C"},
	{400, "CD"},
	{500, "D"},
	{900, "CM"},
	{1000, "M"},
}

func Encode(n int) (string, bool) {
	var res string

	// move from the highest number to lowest
	for i := len(bar) - 1; i >= 0; i-- {

		for n >= bar[i].num {

			n -= bar[i].num
			res += bar[i].roman
		}
	}

	return res, len(res) > 0
}

func Decode(s string) (int, bool) {
	if s == "" {
		return 0, false
	}

	var res int

	// move from the highest number to lowest
	for i := len(bar) - 1; i >= 0; i-- {

		for strings.HasPrefix(s, bar[i].roman) {

			s = s[len(bar[i].roman):]
			res += bar[i].num
		}
	}

	if len(s) != 0 {
		return 0, false
	}

	return res, true
}
