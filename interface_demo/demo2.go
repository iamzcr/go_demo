package main

import "fmt"

//空接口
//空接口可以不用定义任何方法，没有定义任何方法的接口就是空接口，空接口表示没有任何约束，任何类型都可以实现空接口，用空接口可以表示任意数据类型，实际开发中用的非常多
//空接口也可以直接当作类型使用，可以表示任意类型

type EmptyInferface interface {
}

// 空接口可以作为函数的参数
func testFunc(arg interface{}) {
	fmt.Printf("值：%v，类型%T\n", arg, arg)
}
func main() {
	var num EmptyInferface
	var str = "test"
	num = str //让字符串变量str实现EmptyInferface这个接口
	fmt.Printf("值：%v，类型%T\n", num, num)

	var test interface{}
	test = 20
	fmt.Printf("值：%v，类型%T\n", test, test)
	test = "test"
	fmt.Printf("值：%v，类型%T\n", test, test)

	testFunc(222)
	testFunc("test")

	var mapTest = make(map[string]interface{})
	mapTest["name"] = "20"
	mapTest["age"] = "20"
	fmt.Printf("值：%v，类型%T\n", mapTest, mapTest)

	var sliceTest = make([]interface{}, 0)
	sliceTest = append(sliceTest, 1, "测试")
	fmt.Printf("值：%v，类型%T\n", sliceTest, sliceTest)

}
