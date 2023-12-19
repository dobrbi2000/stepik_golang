package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func ParseNumber(numStr string) int64 {
	builder := strings.Builder{}
	for _, r := range numStr {
		if unicode.Is(unicode.Digit, r) {
			builder.WriteRune(r)
		}
	}

	number, err := strconv.ParseInt(builder.String(), 10, 64) //строку в число
	if err != nil {
		panic(err)
	}

	return number

}

func adding(a, b string) int64 {
	return ParseNumber(a) + ParseNumber(b)
}

func main() {
	fmt.Println(adding("%^80", "hhhhh20&&&&nd"))

}
