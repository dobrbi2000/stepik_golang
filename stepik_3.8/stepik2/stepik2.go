package main

import "fmt"

func removeDuplicates(inputStream <-chan string, outputStream chan<- string) {
	defer close(outputStream)

	prevValue := <-inputStream //чтение из канала
	outputStream <- prevValue  //передаем в канал

	for value := range inputStream {
		if prevValue == value {
			continue
		}
		outputStream <- value
		prevValue = value
	}
}

func main() {

	inputStream := make(chan string)
	outputStream := make(chan string)
	go removeDuplicates(inputStream, outputStream)

	go func() {
		defer close(inputStream)

		for _, r := range "112334456" {
			inputStream <- string(r)
		}
	}()

	for x := range outputStream {
		fmt.Print(x)
	}

}
