package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	var res string
	for _, r := range word { // r is of type "rune"
		res = string(r) + res
	}
	return res
}

// func Reverse(word string) string {
// 	var sb strings.Builder
// 	for i := len(word) - 1; i >= 0; i-- {
// 		sb.WriteByte(word[i])
// 	}
// 	return sb.String()
// }

// func Reverse(word string) string {
// 	if len(word) == 0 {
// 		return ""
// 	}

// 	return Reverse(word[1:]) + string(word[0])
// }
