package main

import "fmt"
import "jikeshijian/L5/lib"

var block = "package"
func main() {
	block := "function"
	{
		block := "inner"
		fmt.Printf("The block is %s.\n", block)
	}
	fmt.Printf("The block is %s.\n", block)
	lib.Print();
}