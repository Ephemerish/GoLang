package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")
	var fruitlist = []string{"apple", "tomato", "peach"}
	fmt.Println("Type of fruit list", fruitlist)

	fruitlist = append(fruitlist, "mango", "banana")
	fmt.Println("Type of fruit list", fruitlist)

	fruitlist = append(fruitlist[1:3])
	fmt.Println("Type of fruit list", fruitlist)

	highScore := make([]int, 4)
	highScore[0] = 234
	highScore[1] = 328
	highScore[2] = 472
	highScore[3] = 593

	highScore = append(highScore, 555, 666, 321)

	fmt.Println(highScore)

	sort.Ints(highScore)

	fmt.Println(highScore)

	index := 4
	highScore = append(highScore[:index], highScore[index+1:]...)
	fmt.Println(highScore)
}
