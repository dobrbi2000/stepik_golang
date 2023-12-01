package main

import "fmt"

func main() {

	var (
		x     int
		y     int
		p     int
		years int
	)

	fmt.Scan(&x, &p, &y)

	for ; x < y; years++ {
		x += x * p / 100
	}
	fmt.Println(years)

}

/*

Вклад в банке составляет x рублей. Ежегодно он увеличивается на p процентов, после чего дробная часть копеек отбрасывается.
Каждый год сумма вклада становится больше. Определите, через сколько лет вклад составит не менее y рублей. x > y x := x (x * (p/100))
Входные данные
Программа получает на вход три натуральных числа: x, p, y.
Выходные данные
Программа должна вывести одно целое число.
Sample Input:

100
10
200

*/
