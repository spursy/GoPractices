package main

import "fmt"

func main() {
	ch := make(chan int)
	fmt.Print(<- ch)
}