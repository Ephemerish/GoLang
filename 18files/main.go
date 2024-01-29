package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files")
	content := "This need to go in the file"
	file, err := os.Create("./myLcoGoFile.text")
	checkNilErr(err)

	len, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("lenght is: ", len)
	defer file.Close()

	readFile("./myLcoGoFile.text")
}

func readFile(fileName string) {
	dataByte, err := os.ReadFile(fileName)
	checkNilErr(err)
	fmt.Println("Text Data inside the file is \n", string(dataByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
