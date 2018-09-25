package main
import "fmt"

func foo(i int) chan int {
	c := make(chan int)
	go func () {c <- i}()
	return c
}

func main() {
	c1, c2, c3 := foo(1), foo(2), foo(3)
	c := make(chan int )

	go func() {
		for {
			select{
			case V1 := <- c1: c <- V1
			case V2 := <- c2: c <- V2
			case V3 := <- c3: c <- V3
			}
		}
	}()

	for i := 0; i < 3; i ++ {
		fmt.Println(<-c)
	}
}
