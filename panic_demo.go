package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// 异常处理
// panic可以在任何地方引发，recover只有在defer调用的函数中有效
func catchPanic() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("this is panic")
}

func testFunc() {
	fmt.Println("this is test func")
}
func panicFunc() {
	panic("this is panic")
}
func testFn(n, m int) (cal int, err error) {
	defer func() {
		r := recover()
		if r != nil {
			//TODO可以执行一些监控接口，通知异常
			err = fmt.Errorf("Panic error: %v", r)
		}
	}()
	cal = n / m
	return
}

func checkName(name string) (err error) {
	if name == "iamzcr" {
		return nil
	} else {
		return errors.New("name error")
	}
}
func checkOrder(orderNum string) (err error) {
	if orderNum == "123456" {
		return nil
	} else {
		return errors.New("order num error")
	}
}
func check(orderData map[string]string) (err error) {
	defer func() {
		r := recover()
		if r != nil {
			//TODO可以执行一些监控接口，通知异常
			err = fmt.Errorf("Panic error: %v", r)
		}
	}()
	err = checkName(orderData["name"])
	if err != nil {
		panic(err)
	}
	err = checkOrder(orderData["order_num"])
	if err != nil {
		panic(err)
	}
	return
}

func testGoFn(n, m int) {
	defer func() {
		r := recover()
		if r != nil {
			//TODO可以执行一些监控接口，通知异常
			fmt.Printf("Panic error: %v\n", r)
		}
	}()
	fmt.Printf("cal: %v\n", n/m)
}

func main() {
	//var orderData = make(map[string]string)
	//orderData["name"] = "iamzc1r"
	//orderData["order_num"] = "123456"
	//err := check(orderData)
	//if err == nil {
	//	fmt.Println("order check success")
	//} else {
	//	fmt.Println("order check fail msg:", err)
	//}
	fmt.Println("main start")

	for i := 0; i < 5; i++ {
		n := rand.Intn(3) // 生成0到2之间的随机整数
		m := rand.Intn(3) // 生成0到2之间的随机整数
		go testGoFn(n, m)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("main end")
}
