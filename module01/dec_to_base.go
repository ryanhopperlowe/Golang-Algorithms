package module01

import (
	"fmt"
	"strconv"
)

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//
func DecToBase(dec, base int) string {
	if dec < base {
		result, _ := getNextDigit(dec, base)
		fmt.Println(result)
		return result
	}

	digit, next := getNextDigit(dec, base)

	res := DecToBase(next, base) + digit
	fmt.Println(res)
	return res

	// return DecToBase(next, base) + digit
}

func getNextDigit(dec, base int) (string, int) {
	val := dec % base
	var digit string
	switch {
	case val <= 9:
		digit = strconv.Itoa(val)
	case val == 10:
		digit = "A"
	case val == 11:
		digit = "B"
	case val == 12:
		digit = "C"
	case val == 13:
		digit = "D"
	case val == 14:
		digit = "E"
	case val == 15:
		digit = "F"
	}

	fmt.Println("val =", strconv.Itoa(val), "=>", digit)
	return digit, dec / base
}
