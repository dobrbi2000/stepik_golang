package main

import "fmt"

func work(x int) int {
	return x * x
}

func main() {

	const N = 10
	cache := map[int]int{}

	var n int
	for i := 0; i < N; i++ {
		_, err := fmt.Scan(&n)
		fmt.Println(n)
		if err != nil {
			panic(fmt.Sprintf("invalid int value:%d", n))
		}
	}

	//если значения нет в словаре
	if _, ok := cache[n]; !ok {
		cache[n] = work(n)
	}

	n = cache[n]
	fmt.Printf("%d", n)

}

/*
cache := make(map[int]int, 10)
for n, i := 0, 0; i < 10; i++ {
	fmt.Scan(&n)
	if _, exists := cache[n]; !exists {
		cache[n] = work(n)
	}
	fmt.Print(cache[n], " ")
}


*/

/*
Внутри функции main (объявлять функцию не нужно) необходимо написать программу:
На стандартный ввод подается 10 целых чисел, разделенных пробелами (числа могут
повторяться). Для чтения из стандартного ввода импортирован пакет fmt.
Вам необходимо вычислить результат выполнения функции work для каждого из
полученных чисел. Функция work имеет следующий вид:
func work(x int) int

Результаты вычислений , разделенные пробелами, должны быть напечатаны в строку.
Однако работа функции work занимает слишком много времени. Выделенного вам времени
выполнения не хватит на последовательную обработку каждого числа, поэтому необходимо
реализовать кэширование уже готовых результатов и использовать их в работе.
После завершения работы программы результат выполнения будет дополнен информацией
о соблюдении установленного лимита времени выполнения.

Sample Input:
3 1 5 2 3 5 3 0 3 4

Sample Output:
2 0 6 1 2 6 2 -1 2 5 time limit ok
*/
