package main

func main() {
	// demo 1
	ch1 := make(chan int, 1)
	ch1 <- 1
	ch1 <- 2 // 通道已满，因此这里会阻塞

	// demo 2
	ch2 := make(chan int, 1)
	elem, ok := <-ch2 // 通道已空，因此这里会阻塞
	_,_ = elem, ok

	// demo 3
	var ch3 chan int
	ch3 <- 1 // 通道的值为nil，发送操作和接受操作都会永久处于阻塞状态
}