package main

import (
	"fmt"
	"sync"
	"time"
)

func service(wg *sync.WaitGroup, instance int) {
	time.Sleep(2 * time.Second)
	fmt.Println("Service called on instance", instance)
	wg.Done() //уменьшает счетчик -1
}

func main() {

	fmt.Println("main() started")

	var wg sync.WaitGroup //пустой WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go service(&wg, i)
	}

	wg.Wait() //заблокирует main

	fmt.Println("main() stopped")

}
