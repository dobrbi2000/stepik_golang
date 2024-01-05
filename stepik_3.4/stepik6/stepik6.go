package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Item struct {
	GlobalId int    `json:"global_id"`
	Name     string `json:"Name"`
}

func main() {

	const URL = "https://raw.githubusercontent.com/semyon-dev/stepik-go/master/work_with_json/data-20190514T0100.json"

	rs, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	defer rs.Body.Close()

	dec := json.NewDecoder(rs.Body)

	var items []Item
	err = dec.Decode(&items)
	if err != nil {
		panic(err)
	}

	sum := 0
	for _, x := range items {
		sum += x.GlobalId
	}

	fmt.Println(sum)

}
