package module01

import (
	"fmt"
	"strconv"
)

// BaseToBase takes in a number, the base it is currently
// in, and the base you want it to be converted to. It then
// returns a string representing the number in the new base.
//
// Eg:
//
//   BaseToBase("E", 16, 2) => "1110"
//
func BaseToBase(value string, base, newBase int) string {
	dec := BaseToDec(value, base)
	res := DecToBase(dec, newBase)

	fmt.Println(value+"_"+strconv.Itoa(base), "=>", dec, "=>", res+"_"+strconv.Itoa(newBase))
	return res
}
