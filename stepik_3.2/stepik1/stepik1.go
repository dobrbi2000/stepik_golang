// package main

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// 	"unicode"
// )

// func ParseNumber(numStr string) int64 {
// 	builder := strings.Builder{}
// 	for _, r := range numStr {
// 		if unicode.Is(unicode.Digit, r) {
// 			builder.WriteRune(r)
// 		}
// 	}

// 	number, err := strconv.ParseInt(builder.String(), 10, 64) //строку в число
// 	if err != nil {
// 		panic(err)
// 	}

// 	return number

// }

// func adding(a, b string) int64 {
// 	return ParseNumber(a) + ParseNumber(b)
// }

// func main() {
// 	fmt.Println(adding("%^80", "hhhhh20&&&&nd"))

// }

// package main

// import (
// 	"fmt"
// 	"unicode/utf8"
// )

// type BigWord interface {
// 	isBig() bool
// }

// type MySuperString string

// func (mss MySuperString) isBig() bool {
// 	if utf8.RuneCountInString(string(mss)) > 10 {
// 		return true
// 	}
// 	return false
// }

// func main() {

// 	sample := MySuperString("jsjhvfjvfsvfvdfvdsvfs")
// 	var interfaceSample BigWord // объявление переменной типа интерфейса BigWord
// 	interfaceSample = sample

// 	fmt.Println("IsBig?:", interfaceSample.isBig())

// }

package main

import "fmt"

type FigureBuilder interface {
	Area() int
	Perimeter() int
}

type Rectangle struct {
	lenght, width int
}

func (r Rectangle) Area() int {
	return r.width * r.lenght
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.lenght + r.width)
}

type Circle struct {
	radius int
}

func (c Circle) Area() int {
	return 3 * c.radius * c.radius
}
func (c Circle) Perimeter() int {
	return 2 * 3 * c.radius
}

func main() {

	r1 := Rectangle{10, 20}
	r2 := Rectangle{25, 50}
	r3 := Rectangle{100, 200}

	c1 := Circle{25}
	c2 := Circle{12}
	c3 := Circle{36}

	figures := []FigureBuilder{r1, r2, r3, c1, c2, c3}

	total := 0
	for _, fig := range figures {
		total += fig.Area()
	}
	fmt.Println("Total Area:", total)

}
