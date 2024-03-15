package main

import (
	"fmt"
	"reflect"
)

type Stu struct {
	Name string
}

// 发射获取变量值
func reflectValueFunc(x interface{}) {
	xValue := reflect.ValueOf(x)
	//反射获取变量的原始值::xValue.Int()
	//fmt.Printf("值：%v，原始值：%v\n", xValue, xValue.Int())

	kind := xValue.Kind()
	switch kind {
	case reflect.Int:
		fmt.Printf("值：%v，原始值：%v\n", xValue, xValue.Int())
	case reflect.Float64:
		fmt.Printf("值：%v，原始值：%v\n", xValue, xValue.Float())
	case reflect.String:
		fmt.Printf("值：%v，原始值：%v\n", xValue, xValue.String())
	default:
		fmt.Println("no type")
	}
}
func main() {
	test := "test"
	reflectValueFunc(test)
	person := Stu{
		Name: "test",
	}
	reflectValueFunc(person)
	arr := [2]int{1, 2}
	reflectValueFunc(arr)
	slice := []int{1, 2}
	reflectValueFunc(slice)

}
