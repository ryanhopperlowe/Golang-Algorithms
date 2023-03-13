package module01

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {
	res := 0

	for i, digit := range value {
		val := convertToDecimal(digit)
		convertedVal := int(math.Pow(float64(base), float64(len(value)-i-1))) * val
		fmt.Print(string(digit), " => ", val, " * ", strconv.Itoa(base)+"^"+strconv.Itoa(len(value)-i-1), " = ", convertedVal, " => ")
		fmt.Print(convertedVal, " + ", res, " = ")
		res += convertedVal
		fmt.Println(res)
	}

	return res
}

func convertToDecimal(digit rune) int {
	if strings.Contains("FEDCBA", string(digit)) {
		return int(digit-rune('F')) + 15
	}
	val, _ := strconv.Atoi(string(digit))
	return val
}
