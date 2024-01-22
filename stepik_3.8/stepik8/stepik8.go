// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var i int // i == 0

// func worker(wg *sync.WaitGroup, m *sync.Mutex) {
// 	m.Lock()
// 	i += 1
// 	m.Unlock()
// 	wg.Done()
// }
// func main() {

// 	var wg sync.WaitGroup
// 	var m sync.Mutex

// 	for i := 0; i < 1000; i++ {
// 		wg.Add(1)
// 		go worker(&wg, &m)
// 	}

// 	wg.Wait()

// 	fmt.Println("value of i after 1000 operations is", i)
// }

package main

import (
	"fmt"
	"sync"
)

const N = 10

func main() {
	m := make(map[int]int)

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
			fmt.Println(i)
		}()

	}
	wg.Wait()
	fmt.Println("/////////////////")
	fmt.Println(len(m))

}
