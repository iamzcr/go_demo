package main

import "fmt"

func main() {
	//创建
	//ch := make(chan int, 3)
	//fmt.Printf("值%v,容量%v,长度%v，类型%T\n", ch, cap(ch), len(ch), ch)
	//
	////给管道存储数据
	//ch <- 10
	//ch <- 51
	//ch <- 60
	//fmt.Printf("值%v,容量%v,长度%v，类型%T\n", ch, cap(ch), len(ch), ch)
	//
	////取出管道的值赋值给num
	//num := <-ch
	//fmt.Printf("%v,%T\n", num, num)
	//fmt.Printf("值%v,容量%v,长度%v，类型%T\n", ch, cap(ch), len(ch), ch)
	//
	////取出管道的值
	//<-ch
	//num1 := <-ch
	//fmt.Printf("%v,%T\n", num1, num1)
	//fmt.Printf("值%v,容量%v,长度%v，类型%T\n", ch, cap(ch), len(ch), ch)
	//
	////再给管道存储数据
	//ch <- 51
	//ch <- 60
	//fmt.Printf("值%v,容量%v,长度%v，类型%T\n", ch, cap(ch), len(ch), ch)

	//管道阻塞(死锁)：出现死锁的原因是因为在一个容量为 1 的缓冲型通道中尝试存储了2个数据项。
	//缓冲型通道的容量决定了它能够存储的数据项数量，当通道已满时，任何尝试向通道发送数据的操作都会被阻塞。
	//chTest := make(chan int, 1)
	//chTest <- 22
	//chTest <- 55
	//管道阻塞(死锁):发生死锁的原因是因为在一个容量为 1 的缓冲型通道中进行了多次接收操作而没有进行对应的发送操作。
	//chTest := make(chan int, 1)
	//chTest <- 22
	//<-chTest
	//<-chTest
	//总结，再没有协程的情况下，管道无论是接受，还是发送，如果超出了条件，都会发生死锁（阻塞）

	//遍历管道
	var ch1 = make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch1 <- i
	}
	l := len(ch1)
	for i := 0; i < l; i++ {
		n := <-ch1
		fmt.Printf("值%v,类型%T\n", n, n)

	}
	//在 for range 循环中，从通道 ch1 中接收数据。但是，由于通道没有被关闭，for range 循环会一直等待新的数据项到达通道，而不会主动退出。
	//因此，当所有数据项都被发送到通道 ch1 后，for range 循环仍然在等待新的数据项，而没有其他 goroutine 会向通道发送数据，导致程序陷入死锁状态。
	//close(ch1)
	//for item := range ch1 {
	//	fmt.Printf("值%v,类型%T\n", item, item)
	//}
}
