package main

import "fmt"

func main() {

	var n uint
	fmt.Scan(&n)

	items := make([]int, n)
	for i := range items {
		fmt.Scan(&items[i])
	}
	fmt.Println(items[3])
}

/*
Напишите программу, принимающая на вход число N(N≥4), а затем
N целых чисел-элементов среза. На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.
*/
