package main

import "fmt"

func main() {

	var s string
	_, err := fmt.Scan(&s)
	if err != nil {
		fmt.Println("invalid value")
	}
	rtext := []rune(s)

	fmt.Print(string(rtext[0]))

	for i := 1; i < len(s); i++ {
		fmt.Print("*", rtext[i])
	}
}

/*
Дана строка, содержащая только английские буквы (большие и маленькие).
Добавить символ ‘*’ (звездочка) между буквами (перед первой буквой и после последней символ ‘*’ добавлять не нужно).

*/
