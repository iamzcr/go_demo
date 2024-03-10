package main

import "fmt"

func deferFunc() int {
	var n int
	defer func() {
		n++
	}()
	return n
}
func deferFunc1() (n int) {
	defer func() {
		n++
	}()
	return
}
func main() {
	fmt.Println("main start")
	defer fmt.Println("main defer")
	fmt.Println("main middle")
	defer fmt.Println("main defer middle")
	fmt.Println("main end")

}
