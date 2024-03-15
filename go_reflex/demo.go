package main

import (
	"fmt"
	"reflect"
)

type myInt int

// 有时候我们需要写一个函数，这个函数有能力统一处理各种值类型，而这些类型可能无法共享同一个接口，也有可能布局位置，也有可能这个类型在我们设计函数的时候还不存在，这个时候就可以用反射
// 例如空接口，空接口可以存储任意类型的遍历，我们如何知道这个空接口保存数据的类型是什么，值是什么呢
// 1可以通过类型断言
// 2可以通过反射实现，也就是在程序运行时动态获取一个变量的类型信息和值信息，在结构体转换成json的时候，tag就用到了反射的技术
type Person struct {
	Name string
}

func reflectFunc(x interface{}) {
	xType := reflect.TypeOf(x)
	// xType.Kind()底层的具体类型
	fmt.Printf("类型：%v，类型名称：%v,类型种类：%v\n", xType, xType.Name(), xType.Kind())
}
func main() {
	test := "test"
	reflectFunc(test)

	var num myInt = 10
	reflectFunc(num)
	reflectFunc(&num)
	person := Person{
		Name: "test",
	}
	reflectFunc(person)
	arr := [2]int{1, 2}
	reflectFunc(arr)

	slice := []int{1, 2}
	reflectFunc(slice)

}
