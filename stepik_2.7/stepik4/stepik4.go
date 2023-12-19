package main

import (
	"fmt"
	"strconv"
)

func main() {

	var text string
	_, err := fmt.Scan(&text)
	if err != nil {
		panic("invalid type")
	}

	for _, r := range text {
		n, err := strconv.Atoi(string(r))
		if err != nil {
			panic(fmt.Sprintf("invalid int value: %c", r))
		}
		fmt.Print(n * n)
	}

}

// package main

// import "fmt"

// func main() {
// 	v := 5
// 	p := &v //5
// 	fmt.Print(*p, " ") // 5
// 	changePointer(p)
// 	fmt.Print(*p)
// }

// func changePointer(p *int) { // нет возращаемого параметра
// 	v := 3
// 	p = &v
// }
