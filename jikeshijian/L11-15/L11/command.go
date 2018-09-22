package main

import (
	"fmt"
	"math/rand"
)

func main() {
	intChannels := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}	

	index := rand.Intn(3)
	fmt.Printf("The index: %d\n", index)
	
	select {
	case <- intChannels[0]:
		fmt.Printf("The first candidate case is selected")
	case <- intChannels[1]:
		fmt.Printf("The second candidate case is selected")
	case <- intChannels[2]:
		fmt.Printf("The third candidate case is selected")
	default:
		fmt.Printf("No candidate case is selected")
	}
}