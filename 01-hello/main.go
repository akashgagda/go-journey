package main

import (
	"fmt"
	integers "go-journey/02-integers"
)

func main() {
	name := "akash"
	fmt.Println(Hello(name))

	fmt.Println(integers.Add(2, 2))
}

func Hello(name string) string {
	return "hello " + name
}
