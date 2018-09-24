package main
import "fmt"
var quit chan int

func foo(id int) {
	fmt.Println(id)
	quit <- id
}

func main() {
	count := 100
	// quit = make(chan int)
	quit = make(chan int, 1000)
	for i := 0; i < count; i++ {
		go foo(i)
	}

	for i := 0; i < count; i ++ {
		fmt.Printf(">>> %d", <- quit)
	}

}