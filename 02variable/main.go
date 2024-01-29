package main

import "fmt"

func main() {
	var UserName string = "Kenneth"
	fmt.Println(UserName)
	fmt.Printf("Variable is of type: %T \n", UserName)

	var isLogIn bool = false
	fmt.Println(isLogIn)
	fmt.Printf("Variable is of type: %T \n", isLogIn)

	var smallVal int = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.255252525525252525
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implecit
	var website = "google.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	// no var style
	numberOfUser := 99
	fmt.Println(numberOfUser)
	fmt.Printf("Variable is of type: %T \n", numberOfUser)

}
