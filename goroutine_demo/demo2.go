package main

import (
	"fmt"
	"sync"
	"time"
)

//sync.WaitGroup
//goroutine注意：
//主线程执行完毕，协程没有执行完毕，也会退出
//协程可以在主线程没有执行完毕前提前退出，协程是否执行完毕，不会影响主线程的执行

// 为了保证程序可以顺利执行，要让协程执行完毕后再执行主进程的退出，这个时候，可以使用sync.WaitGroup等待协程执行完毕
var wg sync.WaitGroup

func TestGo() {
	for i := 0; i < 3; i++ {
		fmt.Println("TestGo() hello", i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done() //协程计数-1
}
func TestGo1() {
	for i := 0; i < 3; i++ {
		fmt.Println("TestGo1() hello", i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done() //协程计数-1
}
func main() {
	wg.Add(1) //协程计数+1

	go TestGo() //开启一个协程
	wg.Add(1)   //协程计数+1
	go TestGo() //开启一个协程

	wg.Wait() //等待协程执行完毕

	fmt.Println("main end")

}
