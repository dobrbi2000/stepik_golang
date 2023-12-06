package main

import "fmt"

func main() {
	var n, counter int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		if x == 0 {
			counter++
		}
	}
	fmt.Print(counter)
}

/*
По данным числам, определите количество чисел, которые равны нулю.
Входные данные
Вводится натуральное число N, а затем N чисел.
Выходные данные
Выведите количество чисел, которые равны нулю.
*/
