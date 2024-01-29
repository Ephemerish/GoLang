package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&payment=sadas2321"

func main() {
	fmt.Println("URLS")
	fmt.Println(myUrl)

	// parsing
	result, err := url.Parse(myUrl)
	checkNilErr(err)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()

	// fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println(val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/learn",
		RawQuery: "user=kenneth",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println(anotherUrl)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
