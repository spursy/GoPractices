package main
import (
	"fmt"
	"errors"
)

type operate func(x, y int) int

func calculate(x int, y int, op operate) (int, error) {
	if (op == nil) {
		return 0, errors.New("invalid operation")
	}
	return op(x,y), nil
}

func add(x int, y int) int {
	return x + y
}

func main() {	
	ret, errors := calculate(3, 5, add)
	if (errors == nil) {
		fmt.Println(ret)
	} else {
		fmt.Println(errors)
	}
}