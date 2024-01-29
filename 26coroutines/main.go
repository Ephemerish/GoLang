package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex

// func main() {
// 	fmt.Println("Coroutine")
// 	for i := 1; i < 100000; i++ {
// 		go loadTest()
// 		wg.Add(1)
// 	}
// 	wg.Wait()
// }

// func loadTest() {
// 	mut.Lock()
// 	defer wg.Done()

// 	for j := 1; j < 1000000; j++ {
// 		x := 10 % j
// 		x += x
// 	}
// 	mut.Unlock()
// }

func main() {
	fmt.Println("Coroutine")
	for i := 1; i < 100000; i++ {
		loadTest()
	}
}

func loadTest() {
	for j := 1; j < 1000000; j++ {
		x := 10 % j
		x += x
	}
}
