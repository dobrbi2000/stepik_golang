package main

import "fmt"

func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	maxNum := array[0] // записываем первый элемент массива, чтобы дальше перебрать на максимум все
	for i := 1; i < len(array); i++ {
		x := array[i]
		if x > maxNum {
			maxNum = x
		}
		fmt.Println(maxNum)
	}

}

/*
На ввод подаются пять целых чисел, которые записываются в массив. Однако эта часть программы уже написана.
Вам нужно написать фрагмент кода, с помощью которого можно найти и вывести максимальное число в этом массиве.
*/
