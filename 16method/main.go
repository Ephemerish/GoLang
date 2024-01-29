package main

import "fmt"

func main() {
	fmt.Println("Methods")
	// no inheretance in golang; no super or parent
	kenneth := User{"Kenneth", "kennethpanis24@gmail.com", true, 21}
	fmt.Println(kenneth)
	fmt.Printf("kenneth details are: %+v\n", kenneth)
	fmt.Printf("Name is: %v \nand email is:  %v\n", kenneth.Name, kenneth.Email)

	kenneth.GetStatus()
	kenneth.NewEmail("Kenneth@go.dev")
	fmt.Println(kenneth)
}

type User struct {
	Name   string
	Email  string
	Status bool
	age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user avtive: ", u.Status)
}

func (u User) NewEmail(newEmail string) {
	u.Email = newEmail
	fmt.Println("Email of this user is", u.Email)
}
