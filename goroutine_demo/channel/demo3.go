package main

import "fmt"

// 单向管道，只读，只写
func main() {
	//双向管道
	ch := make(chan int, 2)
	ch <- 20
	num := <-ch
	fmt.Printf("%v\n", num)

	//只读单向管道
	//chRead := make(<-chan int, 2)

	//只写单向管道
	//chWrite := make(chan<- int, 2)

}
