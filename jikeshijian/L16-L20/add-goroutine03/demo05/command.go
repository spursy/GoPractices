package main
import "fmt"

func xrange() chan int {
	var ch chan int = make(chan int)

	go func() {
		for i := 2; ; i ++ {
			ch <- i
		}
	}()
	return ch
}

func filter(in chan int, number int) chan int {
	out := make(chan int)
	go func() {
		for {
			i := <- in
			 if i % number != 0 {
				 out <- i
			 }
		}
	}()
	return out
}

func main() {
	const max = 100
	nums := xrange()
	number := <- nums

	for number <= max {
		fmt.Println(number)
		nums = filter(nums, number)
		number =<- nums
	}
}