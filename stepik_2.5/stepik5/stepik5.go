// package main

// import (
// 	"fmt"
// 	"unicode"
// )

// func isValidPassword(password string) bool {
// 	ps := []rune(password)
// 	if len(ps) < 5 {
// 		return false
// 	}

// 	for _, r := range password {
// 		if !unicode.Is(unicode.Latin, r) && !unicode.Is(unicode.Digit, r) { // если Не Латинская буква И Не цифра верни false !-инвертирует true в false
// 			return false
// 		}
// 	}
// 	return true
// }

// func main() {

// 	var text string

// 	_, _ = fmt.Scan(&text)

// 	if isValidPassword(text) {
// 		fmt.Println("Ok")
// 	} else {
// 		fmt.Println("Wrong password")
// 	}

// }

package main

import "fmt"

func main() {
	var a uint32 = 5
	defer specPrint(a)
	a = 6
	fmt.Print(a)
}
func specPrint(s uint32) uint32 {
	fmt.Print(s)
	s++
	return s
}
