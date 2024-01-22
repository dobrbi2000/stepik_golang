package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("http://golang.org")

	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("%s", data)

}
