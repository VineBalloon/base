package base

import (
	"math"
	"strings"
)

var (
	basechar = map[rune]int{'0': 0,
		'1': 1, '2': 2, '3': 3, '4': 4, '5': 5,
		'6': 6, '7': 7, '8': 8, '9': 9, 'a': 10,
		'b': 11, 'c': 12, 'd': 13, 'e': 14, 'f': 15,
		'g': 16, 'h': 17, 'i': 18, 'j': 19, 'k': 20,
		'l': 21, 'm': 22, 'n': 23, 'o': 24, 'p': 25,
		'q': 26, 'r': 27, 's': 28, 't': 29, 'u': 30,
		'v': 31, 'w': 32, 'x': 33, 'y': 34, 'z': 35,
	}
	charbase = []string{"0",
		"1", "2", "3", "4", "5",
		"6", "7", "8", "9", "A",
		"B", "C", "D", "E", "F",
		"G", "H", "I", "J", "K",
		"L", "M", "N", "O", "P",
		"Q", "R", "S", "T", "U",
		"V", "W", "X", "Y", "Z",
	}
)

// Frombase converts a string representation of a number, given it's base and converts it to decimal.
// Works from bases of 2 to 35
func Frombase(num string, from int) int {
	if from < 2 || from > 35 {
		return -1
	}

	ret := 0
	num = strings.ToLower(num)

	// Get decimal rep
	pow := 0
	chars := []rune(num)
	for i := len(num) - 1; i >= 0; i-- {
		factor := int(math.Pow(float64(from), float64(pow)))

		digit := basechar[chars[i]]

		ret += digit * factor
		pow++
	}

	return ret
}

// Tobase converts an integer to a string representation given the desired base
// Works from bases of 2 to 35
func Tobase(num int, to int) string {
	if to < 2 || to > 35 {
		return "read the doc, baka"
	}

	ret := ""

	// Convert to string rep
	for num > 0 {
		ret += charbase[num%to]
		num /= to
	}

	return Reverse(ret)
}

// Reverse reverses a string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
