package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	res, err := http.Get("http://127.0.0.1:5555/get")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body) //читаем ответ
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}

/*
Сделайте HTTP запрос на сервер по пути http://127.0.0.1:5555/get и напечатайте ответ сервера (только тело).
*/
