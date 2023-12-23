// package main

// import (
// 	"fmt"
// 	"unicode/utf8"
// )

// type BigWord interface {
// 	isBig() bool
// }

// type MySuperString string

// func (mss MySuperString) isBig() bool {
// 	if utf8.RuneCountInString(string(mss)) > 10 {
// 		return true
// 	}
// 	return false
// }

// func main() {

// 	sample := MySuperString("jsjhvfjvfsvfvdfvdsvfs")
// 	var interfaceSample BigWord // объявление переменной типа интерфейса BigWord
// 	interfaceSample = sample

// 	fmt.Println("IsBig?:", interfaceSample.isBig())

// }
