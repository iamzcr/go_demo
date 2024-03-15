package main

//通过反射判断类型修改值
import (
	"fmt"
	"reflect"
)

func reflectSetValue(x interface{}) {
	xValue := reflect.ValueOf(x)
	fmt.Println(xValue.Elem().Kind())
	if xValue.Elem().Kind() == reflect.Int64 {
		xValue.Elem().SetInt(200)
	}
}
func main() {
	var num int64 = 100
	reflectSetValue(&num)
	fmt.Println("num=", num)
}
