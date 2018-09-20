package main

import (
	"flag"
	"fmt"
)

func main() {
	// method 1
	// var name string
	// flag.StringVar(&name, "name", "everyone", "The greeting object")

	// method 2
	var name = *flag.String("name", "everyone", "The greeting object!")

	// method 3
	// name := *flag.String("name", "everyone", "The greeting object!")

	flag.Parse()
	fmt.Printf("Hello, %v!\n", name)
}