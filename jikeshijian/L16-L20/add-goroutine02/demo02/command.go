package main
import "fmt"
import "time"

var quit chan int

func foo(id int) {
	fmt.Println(id)
	time.Sleep(time.Second)
	quit <- id
}

func main() {
	count := 1000
	quit = make(chan int, count)

	for i := 0; i < count; i++ {
		go foo(i)
	}
	for i := 0; i < count; i++ {
		ret := <- quit		
		fmt.Printf(">>> %d", ret)
	}
}