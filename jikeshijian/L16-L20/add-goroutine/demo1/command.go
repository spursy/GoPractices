package main

import "fmt"

func loop() {
	for i:=0; i < 10; i ++ {
		fmt.Println(i)
	}
}
func main() {
	loop()
	loop()
}