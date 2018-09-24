package main

import (
	"fmt"
	// "time"
)

func main() {
	num := 10
	sign := make(chan struct{}, num)

	for i := 0; i < num; i ++ {
		go func() {
			fmt.Println(i)
			sign <- struct{}{}
		}()
	}

	// method 1
	// time.Sleep(time.Millisecond * 500)

	// method 2
	for j := 0; j < num; j ++ {
		<-sign
	}
}