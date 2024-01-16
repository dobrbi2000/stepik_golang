// package main

// import "fmt"

// func greet(c chan string) {
// 	fmt.Println("Hello " + <-c + "!")
// }

// func main() {

// 	fmt.Println("main() started")

// 	c := make(chan string)

// 	go greet(c)

// 	c <- "Jack"

// 	fmt.Println("main() stopped")

// }

package main

import "fmt"

func squares(c chan int) {
	for i := 0; i <= 3; i++ {
		num := <-c
		fmt.Println(num * num)
	}

}

func main() {

	fmt.Println("main() started")

	c := make(chan int, 3)

	go squares(c) // запуск горутины

	c <- 1
	c <- 2
	c <- 3
	c <- 4

	fmt.Println("main() stopped")

}
