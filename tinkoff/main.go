// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func nameTinkoff(letters string) bool {
// 	// Создаем словарь для подсчета количества каждой буквы
// 	lettersDictin := make(map[rune]int)

// 	// Подсчитываем количество каждой буквы в привезенном наборе
// 	for _, letter := range letters {
// 		lettersDictin[letter]++
// 	}

// 	// Проверяем, присутствует ли каждая буква из "TINKOFF" в словаре и является ли количество ее присутствий достаточным
// 	tinkoffWord := "TINKOFF"
// 	for _, letter := range tinkoffWord {
// 		if lettersDictin[letter] == 0 {
// 			return false
// 		}
// 		lettersDictin[letter]--
// 	}

// 	return true
// }

// func main() {
// 	var s int
// 	fmt.Scan(&s)

// 	for i := 0; i < s; i++ {
// 		var letters string
// 		fmt.Scan(&letters)

//			if nameTinkoff(strings.ToUpper(letters)) {
//				fmt.Println("YES")
//			} else {
//				fmt.Println("NO")
//			}
//		}
//	}
package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	prices := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&prices[i])
	}
	sort.Ints(prices)

	maxMoney := 0
	for i := 0; i < n && prices[i] <= m; i++ {
		maxMoney = m - prices[i]
	}
	fmt.Println(maxMoney)
}
