package main

import (
	"fmt"
	"time"
)

func timer(duration time.Duration) chan bool {
	ch := make(chan bool)
	 
	go func() {
		time.Sleep(duration)
		ch <- true
	}()
	return ch
}

func main() {
	timeout := timer(time.Second * 10)

	for {
		select {
		case <- timeout:
			fmt.Println("Already ls !")
			return
		}
	}
}