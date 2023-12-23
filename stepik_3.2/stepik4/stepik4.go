// package main

// import (
// 	"fmt"
// )

// /*
// Внутри функции main (объявлять ее не нужно) вы должны объявить функцию вида
// func(uint) uint, которую внутри функции main в дальнейшем можно будет вызвать
// по имени fn.

// Функция на вход получает целое положительное число (uint), возвращает число
// того же типа в котором отсутствуют нечетные цифры или цифра 0, если же получившееся
// число равно 0, то возвращается число 100. Цифры в возвращаемом числе имеют тот же
// порядок, что и в исходном числе.

// Пакет main объявлять не нужно!
// Вводить и выводить что-либо не нужно!
// Уже импортированы - "fmt" и "strconv", другие пакеты импортировать нельзя.

// Sample Input:
// 727178

// Sample Output:
// 28
// */

// // func main() {
// // 	fn := func(num uint) uint {
// // 		numStr := strconv.FormatUint(uint64(num), 10)
// // 		rs := make([]rune, 0, len(numStr))

// // 		for _, r := range numStr {
// // 			if r%2 == 0 && r != 48 {
// // 				rs = append(rs, r)
// // 			}
// // 		}
// // 		numStr = string(rs)

// // 		x, _ := strconv.Atoi(numStr)
// // 		if x == 0 {
// // 			x = 100
// // 		}

// // 		return uint(x)
// // 	}

// // 	fmt.Println(fn(7271780))
// // 	// 28
// // }

// type animal interface {
// 	makeSound()
// }

// type cat struct{}
// type dog struct{}

// func (c *cat) makeSound() {
// 	fmt.Println("meow!")
// }

// func (c *dog) makeSound() {
// 	fmt.Println("wowf!")
// }

// func main() {

// 	var c, d animal = &cat{}, &dog{}

// 	c.makeSound()
// 	d.makeSound()

// }

package main

import (
	"fmt"
	"unicode/utf8"
)

type BigWord interface {
	isBig() bool
}

type MySuperString string

func (mss MySuperString) isBig() bool {
	if utf8.RuneCountInString(string(mss)) > 10 {
		return true
	}
	return false
}

func main() {

	sample := MySuperString("jsjhvfjvfsvfvdfvdsvfs")
	var interfaceSample BigWord // объявление переменной типа интерфейса BigWord
	interfaceSample = sample

	fmt.Println("IsBig?:", interfaceSample.isBig())

}
