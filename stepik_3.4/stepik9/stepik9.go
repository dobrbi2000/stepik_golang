package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {

	var text string //= "13.03.2018 14:00:15,12.03.2018 14:00:15"

	text, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}

	text = strings.TrimSpace(text)

	//13.03.2018 14:00:15,12.03.2018 14:00:15

	items := strings.Split(text, ",")
	if len(items) != 2 {
		panic("Invalid text format")
	}
	dateStr1, dateStr2 := strings.TrimSpace(items[0]), strings.TrimSpace(items[1])

	const dateLayot = "02.01.2006 15:04:05"

	date1, err := time.Parse(dateLayot, dateStr1)
	if err != nil {
		panic(err)
	}

	date2, err := time.Parse(dateLayot, dateStr2)
	if err != nil {
		panic(err)
	}

	if date1.After(date2) { //13 после 12 13-12
		fmt.Println(date1.Sub(date2))
	} else {
		fmt.Println(date2.Sub(date1))
	}

}

/*
На стандартный ввод подается строковое представление двух дат,
разделенных запятой (формат данных смотрите в примере).

Необходимо преобразовать полученные данные в тип Time,
а затем вывести продолжительность периода
между меньшей и большей датами.

Sample Input:
13.03.2018 14:00:15,12.03.2018 14:00:15

Sample Output:
24h0m0s
*/
