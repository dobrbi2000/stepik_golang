package main

import (
	"fmt"
	"strings"
)

type Battery struct {
	capacity int
}

func NewBattery(value string) Battery {
	return Battery{
		capacity: strings.Count(value, "1"),
	}
}

func (b Battery) String() string {
	return fmt.Sprintf("[%10s]", strings.Repeat("X", b.capacity))
}

//type BatteryForTest string
// func (b BatteryForTest) String() string{
//     col := strings.Count(string(b), "1")
//     return fmt.Sprintf("[%10s]", strings.Repeat("X", col))

func main() {
	fmt.Println(NewBattery("1000010011"))

	var value string

	// _, err := fmt.Scan(&value)
	// if err != nil {
	// 	panic(err)
	// }

	batteryForTest := NewBattery(value)
	fmt.Println(batteryForTest)

}

/*
Ваш тип должен предусматривать, что на печати он будет выглядеть так:
[      XXXX]: где пробелы - "опустошенная" емкость батареи, а X - "заряженная".

Во-вторых, на стандартный ввод вы получаете строку, состоящую ровно из 10 цифр:
0 или 1 (порядок 0/1 случайный). Ваша задача считать эту строку любым возможным
способом и создать на основе этой строки объект объявленного вами на первом этапе
типа: надеюсь, вы понимаете, что строка символизирует емкость батарейки: 0 - это
"опустошенная" часть, а 1 - "заряженная".

В-третьих, созданный вами объект должен называться batteryForTest (использование
этого имени обязательно).

В вашем распоряжении фактически весь файл, НО завершающая фигурная скобка функции
main() вам не видна, но она присутствует. Перед этой скобкой присутствует функция
(которая вам тоже не видна), принимающая в качестве аргумента объект типа
fmt.Stringer - batteryForTest, и направляющая его на стандартный вывод,
поэтому вам не требуется выводить что-то на печать самостоятельно.
Удачи!

Sample Input:
1000010011

Sample Output:
[      XXXX]
*/
