// package main

// import "fmt"

// func fiboNum(A int) int {
// 	F1 := 1
// 	F2 := 1
// 	n := 2

// 	for F2 < A {
// 		temp := F2
// 		F2 = F1 + F2
// 		F1 = temp
// 		n++

// 	}

// 	if F2 == A {
// 		return n
// 	} else {
// 		return -1
// 	}
// }

// func main() {

// 	var A int
// 	fmt.Scan(&A)

// 	result := fiboNum(A)
// 	fmt.Print(result)

// }

package main

import "fmt"

func main() {
	var n int
	f1 := 1
	f2 := 1
	res := -1
	i := 3
	fmt.Scan(&n)
	for ; f2 < n; i++ {
		f1, f2 = f2, f1+f2
		if f2 == n {
			res = i
		}
	}
	fmt.Print(res)
}

/*

Номер числа Фибоначчи
Дано натуральное число A > 1.
Определите, каким по счету числом Фибоначчи оно является, то есть выведите такое число n, что φn=A.
Если А не является числом Фибоначчи, выведите число -1.




*/
