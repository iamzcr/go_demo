package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func testCal() {
	for i := 2; i < 500000; i++ {
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

// 1-100000
// 100001-100000
// 100001-200000
// 200001-300000
// 300001-400000
// 400001-500000
func testCalGo(n int) {
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

// 计算素数，除了被1和被自己整除
func main() {
	start := time.Now().Unix()
	/*
		for i := 1; i <= 6; i++ {
			wg.Add(1)
			go testCalGo(i)
		}
		wg.Wait()
	*/
	testCal()
	fmt.Println("main end")
	end := time.Now().Unix()
	fmt.Println("消耗时间：", end-start)

}
