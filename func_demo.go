package main

import "fmt"

// 函数基本使用
func sumFn(n int, m int) int {
	sum := n + m
	return sum
}

// 形参的简写，同类型的形参可以忽略前面参数的类型
func sumFn1(n, m int) int {
	sum := n + m
	return sum
}

// 可变参数
func sumMuti(n ...int) int {
	fmt.Printf("n值=%v，n类型=%T\n", n, n)
	sum := 0
	for _, i := range n {
		//sum = sum + i
		sum += i
	}
	return sum
}

// 可变参数和固定参数接收
func sumMuti1(n int, m ...int) int {
	fmt.Printf("n值=%v，n类型=%T\n", n, n)
	sum := n
	for _, i := range m {
		//sum = sum + i
		sum += i
	}
	return sum
}

// 函数多个返回值
func sumFnReturn(n int, m int) (int, int) {
	return n + m, n - m
}

// 函数返回值命名，原理是先声明要返回的值的名字，在函数体内操作该名字，最后直接return返回
func sumFnReturn1(n int, m int) (sum int) {
	sum = n + m
	return
}

// 函数返回值命名，多个同类新返回值
func calFn(n int, m int) (sum, sub int) {
	sum = n + m
	sub = n - m
	return
}

// 没有参数，没有返回值的函数
func test() {
	fmt.Println("没有参数，没有返回值的函数")
}

func main() {
	fmt.Printf("%v\n", sumFn(1, 2))
	fmt.Printf("%v\n", sumFn1(1, 2))
	fmt.Printf("%v\n", sumMuti(1, 2, 3))
	fmt.Printf("%v\n", sumMuti1(1, 2, 3))
	x, y := sumFnReturn(3, 2)
	fmt.Printf("x=%v,y=%v\n", x, y)
	fmt.Printf("%v\n", sumFnReturn1(1, 2))
	n, m := calFn(3, 2)
	fmt.Printf("x=%v,y=%v\n", n, m)

	//忽略返回
	sum, _ := calFn(3, 2)
	fmt.Printf("x=%v\n", sum)

}
