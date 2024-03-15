package main

import (
	"fmt"
	"time"
)

// 管道(channel) ,管道是go在语言级别提供的goroutine间的通讯方式，
// 使用channel在多个goroutine之间传递信息，channel是可以让一个goroutine发送特定值到另外一个goroutine的通信机制。
// go语言的并发模型是CSP，提倡通过通信共享内存，而不是通过共享内存实现通信
// go中管道是一种特殊类型，管道像一个传送带或者队列，总是遵循先入先出的规则，保证收发数据的顺序，每个管道都是一个具体的类型的导管，也就是声明channel的时候需要为其指定元素类型
// 管道是一种引用类型

// 实现计算和打印
func testCalChannel(n int) {
	for i := (n - 1) * 100000; i < n*100000; i++ {
		if i > 1 {
			var flag = true
			for j := 2; j < i; j++ {
				if i%j == 0 {
					flag = false
					break
				}
			}
			if flag {
				//fmt.Printf("%v是素数\n", i)
			}
		}
	}
	wg.Done()
}

func main() {
	start := time.Now().Unix()
	/*
		for i := 1; i <= 6; i++ {
			wg.Add(1)
			go testCalGo(i)
		}
		wg.Wait()
	*/
	fmt.Println("main end")
	end := time.Now().Unix()
	fmt.Println("消耗时间：", end-start)
}
