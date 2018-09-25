package main
import "fmt"

var ch chan int = make(chan int)

func loop() {
	for i := 0; i < 10; i ++ {
		fmt.Printf("-> %d", i)
	}
	fmt.Println()
	ch <- 10
}

func main() {
	go loop()
	
	fmt.Println(<- ch)
}