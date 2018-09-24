package main
import "fmt"

func main() {
	var messages chan string = make(chan string)
	go func(message string) {
		messages <- message
	}("ping ping ")
	
	fmt.Println(<-messages)
}