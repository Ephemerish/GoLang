package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on Pointers")

	// var ptr *int
	// fmt.Println("Pointers value is: ", ptr)

	myNumber := 23
	ptr := &myNumber
	pptr := &ptr

	fmt.Println("Pointers adress is: ", ptr)
	fmt.Println("Pointers value is: ", *ptr)
	fmt.Println("Pointers Pointers adress is: ", pptr)
	fmt.Println("Pointers Pointers value is: ", **pptr)

}
