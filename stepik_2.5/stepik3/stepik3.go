package main

import (
	"fmt"
	"strings"
)

func main() {

	var text string = "ihgewlqlkot"
	//_, _ = fmt.Scan(&text)

	var b strings.Builder
	for i, r := range text {
		if i%2 != 0 {
			b.WriteRune(r)
		}
	}
	fmt.Println(b.String())

}
