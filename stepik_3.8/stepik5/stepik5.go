// package main

// import "fmt"

// func greeter(cc chan chan string) {
// 	c := make(chan string)
// 	cc <- c
// }

// func greet(c chan string) {
// 	fmt.Println("Hello " + <-c + "!")

// }

// func main() {

// 	fmt.Println("main() started")

// 	cc := make(chan chan string)

// 	go greeter(cc)

// 	c := <-cc

// 	go greet(c)

// 	c <- "Bob"

// 	fmt.Println("main() stopped")

// }

package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func service1(c chan string) {
	//time.Sleep(1 * time.Second)
	c <- "Hello from service #1"
}

func service2(c chan string) {
	//	time.Sleep(1 * time.Second)
	c <- "Hello from service #2"
}

func main() {

	fmt.Println("main() started", time.Since(start))

	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service2(chan2)

	select {
	case res := <-chan1:
		fmt.Println("Response from service 1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from service 2", res, time.Since(start))
	}
	fmt.Println("main() stopped", time.Since(start))

}
