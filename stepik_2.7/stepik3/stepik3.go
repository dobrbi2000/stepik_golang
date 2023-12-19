package main

import "fmt"

func main() {

	var text string
	_, err := fmt.Scan(&text)
	if err != nil {
		fmt.Println("invalid value")
	}
	rtext := []rune(text)

	maxR := rtext[0]

	for i := 0; i < len(rtext); i++ {
		r := rtext[i]
		if r > maxR {
			maxR = r
		}
	}
	fmt.Println(string(maxR))

}

/*
Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков и строка содержит только арабские цифры.
Выходные данные
Выведите максимальную цифру, которая встречается во введенной строке.
Sample Input:
1112221112
*/
