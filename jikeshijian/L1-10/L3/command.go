package main

import (
	"flag"
)

var name string

func init() {
	flag.StringVar(&name, "name", "every one", "The greeting object")
}

func main() {
	flag.Parse()
	hello(name)
}