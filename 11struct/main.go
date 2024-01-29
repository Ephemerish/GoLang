package main

import "fmt"

func main() {
	fmt.Println("Structs")
	// no inheretance in golang; no super or parent
	kenneth := User{"Kenneth", "kennethpanis24@gmail.com", true, 21}
	fmt.Println(kenneth)
	fmt.Printf("kenneth details are: %+v\n", kenneth)
	fmt.Printf("Name is: %+v and email is: \n", kenneth.Name, kenneth.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	age    int
}
