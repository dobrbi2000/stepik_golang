// package main

// import "fmt"

// func main() {
// 	var mas int = 50
// 	//fmt.Scan(&mas)
// 	for i := 1; i <= mas; i *= 2 {
// 		fmt.Print(i, " ")
// 	}
// }

// /*
// По данному числу N распечатайте все целые значения степени двойки, не превосходящие N, в порядке возрастания.
// Входные данные
// Вводится натуральное число.
// Выходные данные
// Выведите ответ на задачу.

// */

// package main

// import "fmt"

// var a, b, c, d int

// func main() {

// 	fmt.Scan(&a, &b, &c, &d)

// 	if d > b {
// 		a = a + 12*(d-b)
// 		fmt.Println(a)
// 	} else {
// 		fmt.Println(a)
// 	}

// }

/*
Вводится 4 целых положительных числа
A,B,C,D(1≤A,B,C,D≤100) — стоимость тарифа Кости, размер тарифа Кости, стоимость каждого лишнего мегабайта, размер интернет-трафика Кости в следующем месяце.
Числа во входном файле разделены пробелами.

В первом примере Костя сначала оплатит пакет интернета, после чего потратит на 5 мегабайт больше, чем разрешено по тарифу.
Следовательно, за 100+12×5=160 рублей.
Во втором примере Костя укладывается в тарифный план, поэтому платит только за него.
a	b	c	d
100 10 12 15 = 160
100 10 12 1 = 100

допустимо 10 mb

*/

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {

// 	var n, s int
// 	fmt.Scan(&n, &s)

// 	items := make([]int, n)
// 	for i := range items {
// 		fmt.Scan(&items[i])
// 	}
// 	sort.Sort(sort.Reverse(sort.IntSlice(items)))

// 	for _, price := range items {
// 		if price <= s {
// 			fmt.Println(price)
// 			return
// 		}
// 	}
// 	fmt.Println(0)

// }

// package main

// import "fmt"

// func countSheriffWords(s string) int {
// 	sheriffWord := "sheriff"
// 	sheriffLetters := make(map[rune]int)
// 	for _, c := range sheriffWord {
// 		sheriffLetters[c]++
// 	}

// 	lettersCount := make(map[rune]int) //map[int32]int [101: 2, 114: 4, 105: 2, 115: 3, 107: 1, 102: 4, 104: 2, 97: 1, 122: 2, ]
// 	for _, c := range s {              //map[int32]int [115: 1, 104: 1, 101: 1, 114: 1, 105: 1, 102: 2, ]
// 		lettersCount[c]++
// 	}

// 	minWords := len(s)

// 	for c, count := range sheriffLetters {
// 		if count != 0 {
// 			words := lettersCount[c] / count
// 			if words < minWords {
// 				minWords = words
// 			}
// 		}
// 	}
// 	return minWords
// }

// func main() {

// 	var s string = "thegorillaiswathcing"

// 	result := countSheriffWords(s)

// 	fmt.Println(result)

// }
package main

import "fmt"

func isWinnableSequence(n int, a []int, b []int) string {
start := -1
end := -1

for i := 0; i < n; i++ {
if a[i] != b[i] {
if start == -1 {
start = i
}
end = i
}
}

if start == -1 && end == -1 {
return "YES"
}

reverse := make([]int, end-start+1)
copy(reverse, a[start:end+1])
for i, j := 0, len(reverse)-1; i < j; i, j = i+1, j-1 {
reverse[i], reverse[j] = reverse[j], reverse[i]
}

for i := start; i <= end; i++ {
a[i] = reverse[i-start]
}

for i := 0; i < n; i++ {
if a[i] != b[i] {
return "NO"
}
}

return "YES"
}

func main() {
var n int
fmt.Scan(&n)

a := make([]int, n)
for i := 0; i < n; i++ {
fmt.Scan(&a[i])
}

b := make([]int, n)
for i := 0; i < n; i++ {
fmt.Scan(&b[i])
}

result := isWinnableSequence(n, a, b)
fmt.Println(result)
}
Пожалуйста, обратите внимание, что этот код проверяет, можно ли для последовательности `a` получить последовательность `b` путем однократной сортировки подпоследовательности.