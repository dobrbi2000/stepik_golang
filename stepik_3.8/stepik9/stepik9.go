package main

func main() {

	func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
		
		result := make(chan int)

		go func() {
			defer close(result)

			select {
			case value := <-firstChan:
				result <- value * value
			case value := <-secondtChan:
				result <- value * 3
			case <- stopChan:
				return
			}
		}()
		return result
	}

	// firstChan := make(chan int)
	// secondChan := make(chan int)
	// stopChan := make(chan int)

}
















/*
Вам необходимо написать функцию calculator следующего вида:
    func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int

Функция получает в качестве аргументов 3 канала, и возвращает канал типа <-chan int.
 * в случае, если аргумент будет получен из канала firstChan, в выходной (возвращенный) канал
   вы должны отправить квадрат аргумента.
 * в случае, если аргумент будет получен из канала secondChan, в выходной (возвращенный) канал
   вы должны отправить результат умножения аргумента на 3.
 * в случае, если аргумент будет получен из канала stopChan, нужно просто завершить работу функции.

Функция calculator должна быть неблокирующей, сразу возвращая управление. Ваша функция получит
всего одно значение в один из каналов - получили значение, обработали его, завершили работу.

После завершения работы необходимо освободить ресурсы, закрыв выходной канал, если вы этого
не сделаете, то превысите предельное время работы.
*/
