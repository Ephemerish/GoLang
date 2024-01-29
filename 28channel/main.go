package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels")

	myChannel := make(chan int, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go func(ch <-chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		val, isChannelOpen := <-myChannel
		fmt.Println(isChannelOpen)
		fmt.Println(val)
	}(myChannel, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {
		defer wg.Done()
		myChannel <- 0
		close(myChannel)
		// myChannel <- 5
		// myChannel <- 6
	}(myChannel, wg)

	wg.Wait()
}
