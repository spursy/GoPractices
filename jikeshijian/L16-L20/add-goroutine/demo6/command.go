package main
import "fmt"

var ch chan int = make(chan int)

func foo(id int) {
	ch <- id
}

func main() {
	for i := 0; i < 5; i++ {
		go foo(i)
	}

	for i := 0; i < 5; i ++ {
		fmt.Println(<- ch)
	}
}