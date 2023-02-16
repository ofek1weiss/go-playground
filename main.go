package main

import (
	"fmt"
	"linkparser/linkparser"
	"os"
)

func main() {
	file, err := os.Open("examples/4.html")
	if err != nil {
		panic(err)
	}
	links, err := linkparser.Parse(file)
	if err != nil {
		panic(err)
	}
	for _, l := range links {
		fmt.Println(l)
	}
}
