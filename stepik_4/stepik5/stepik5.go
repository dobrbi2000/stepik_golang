package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {

	var name string = "alex"
	var age string = "10"

	//fmt.Scanln(&name, &age)

	baseUrl := ("http://127.0.0.1:8080/hello")

	queryParams := url.Values{}
	queryParams.Set("name", name)
	queryParams.Set("age", age)

	baseUrl += "?" + queryParams.Encode()

	response, err := http.Get(baseUrl)
	if err != nil {
		fmt.Println("Произошла ошибка при выполнении HTTP запроса:", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Произошла ошибка при чтении ответа сервера:", err)
		return
	}
	fmt.Println(string(body))
}

/*
Считайте с консоли две переменные, сначала name, затем age.
Сделайте HTTP запрос на http://127.0.0.1:8080/hello передав name и age в query параметрах,
затем напечатайте ответ сервера (только тело).
*/
