package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b float64
	_, err := fmt.Scan(&a, &b)

	if err != nil {
		panic("invalid value")
	}

	fmt.Println(math.Sqrt(a*a + b*b))
}
