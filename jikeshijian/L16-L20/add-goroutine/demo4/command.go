package main
import "fmt"

var ch chan int = make(chan int)

func foo() {
	ch <- 10
}

func main() {
	go foo()
	fmt.Println(<- ch)
}