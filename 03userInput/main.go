package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Printf(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Your name:")

	// comma ok systax || err ok

	input, _ := reader.ReadString('\n')

	fmt.Println("Your name is,", input)
	fmt.Printf("Type of rating is %T", input)
}
