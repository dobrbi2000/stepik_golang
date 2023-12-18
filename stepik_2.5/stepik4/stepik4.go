package main

import (
	"fmt"
	"strings"
)

func main() {
	var text string = "zaabcbd"

	//_, _ = fmt.Scan(&text)
	var b strings.Builder

	for _, r := range text {
		if strings.Count(text, string(r)) == 1 {
			b.WriteRune(r)

		}
	}
	fmt.Println(b.String())

}

/*
package main

import (
    "fmt"
    "strings"
)

func main() {

    var a string
    fmt.Scan(&a)

    for _, ch := range a {
        if strings.Count(a, string(ch)) == 1 {
            fmt.Print(string(ch))
        }
    }
}




*/
