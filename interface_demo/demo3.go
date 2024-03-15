package main

// 类型断言
// 一个接口的值是由一个具体类型和具体类型值两部分组成，这两部分分别称为动态类型和动态值
// 入股我们想要判断空接口值的类型，这个时候就可以使用类型断言
// 使用也非常多
//value.(type)只能用在switch
import "fmt"

func PrintType(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("Type: int")
	case string:
		fmt.Println("Type: string")
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	var test interface{}
	test = "test"
	_, ok := test.(int)
	if ok {
		fmt.Printf("test是int类型，值：%v，类型%T\n", test, test)
	} else {
		fmt.Printf("test不是int类型\n")

	}
	var test1 interface{}
	test1 = 10
	_, test1Ok := test1.(int)
	if test1Ok {
		fmt.Printf("test1是int类型，值：%v，类型%T\n", test1, test1)
	} else {
		fmt.Printf("test1不是int类型\n")

	}

	PrintType("iamzcr")
}
