package main
import "jikeshijian/L3Upgrade/lib"

import (
	"flag"
)

var name string

func init() {
	flag.StringVar(&name, "name", "every one", "The greeting object")
}

func main() {
	flag.Parse()
	lib.Hello(name)
}