package main

import (
	"fmt"
	"sync"
	"time"
)

//goroutine和channel协同工作案例

var wggo sync.WaitGroup

// 写入管道
func testWrite(ch chan int) {
	defer close(ch) //写入完成，必须关闭管道

	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Printf("写人数据到管道，%v-%T\n", i, i)
		time.Sleep(time.Millisecond * 10) //写的时候比较慢，读比较快，并不会发生阻塞，因为go会等待管道里面的数据发过来
	}
	wggo.Done()
}

func testRead(ch chan int) {
	for item := range ch {
		fmt.Printf("读取管道的数据，%v-%T\n", item, item)
		time.Sleep(time.Millisecond * 1000)
	}
	wggo.Done()
}

func main() {

	ch := make(chan int, 10)
	wggo.Add(1)
	go testWrite(ch)
	wggo.Add(1)
	go testRead(ch)
	wggo.Wait()
	fmt.Println("main end")

}
