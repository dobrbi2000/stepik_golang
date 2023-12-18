package main

import "fmt"

func toReversed(text string) string {
	rs := []rune(text)

	reversed_rs := make([]rune, len(rs))

	for i := 0; i < len(reversed_rs); i++ {
		reversed_rs[i] = rs[len(reversed_rs)-i-1] // 6 - 0 - 1 = 5 в конец, 5 индекс равен нулевому
	}
	return string(reversed_rs)
}

func main() {
	var text string = "топот1"

	_, _ = fmt.Scan(&text)

	reversed := toReversed(text)

	if text == reversed {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}

}
