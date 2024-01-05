package main

import (
	"fmt"
	"time"
)

func main() {

	var dateStr string

	_, err := fmt.Scan(&dateStr)

	if err != nil {
		panic(err)
	}

	date, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		panic(err)
	}

	fmt.Println(date.Format(time.UnixDate))

}
