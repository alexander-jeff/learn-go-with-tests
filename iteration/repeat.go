package iteration

import "strings"

// Repeat repeats a given string times times.
//
//
// Unlike strings.Repeat,
// this Repeat() will not panic if a negative times value or
// if the value would result in an overflow. Instead
// it will just return an empty string.
func Repeat(in string, times int) string {
	if !isValid(in, times) {
		return ""
	}
	return strings.Repeat(in, times)
}

func isValid(in string, times int) bool {
	negativeOrNone := times <= 0
	if negativeOrNone {
		return false
	}
	willOverflow := (len(in)*times/times != len(in))
	return !(willOverflow)
}
