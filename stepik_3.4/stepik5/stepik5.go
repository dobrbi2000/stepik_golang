package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Group struct {
	ID       int
	Number   string
	Year     int
	Students []Student
}

type Student struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Phone      string
	Rating     []int
}

type Result struct {
	Average float64
}

func main() {

	var group Group

	if err := json.NewDecoder(os.Stdin).Decode(&group); err != nil {
		panic(err)
	}

	var numberRating int
	for _, student := range group.Students {
		numberRating += len(student.Rating)
	}

	result := Result{
		Average: float64(numberRating) / float64((len(group.Students))),
	}

	dataJson, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dataJson))

}
