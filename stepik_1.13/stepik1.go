package main

import "fmt"

func main() {

	var a, a1, a2, a3 int
	fmt.Scan(&a)

	a1 = a % 10       //3
	a2 = a / 10 % 10  //2
	a3 = a / 100 % 10 //1

	sum := a1 + a2 + a3
	fmt.Print(sum)

}

/*
Дано трехзначное число. Найдите сумму его цифр.
Формат входных данных
На вход дается трехзначное число.
Формат выходных данных
Выведите одно целое число - сумму цифр введенного числа.
*/
