package main

import "fmt"

//接口和类型断言使用细节

type Address struct {
	Name string
}

func main() {

	var user = make(map[string]interface{})
	user["name"] = "小明"
	user["age"] = 20
	user["hobby"] = []string{"足球", "口琴"}
	var addr = Address{
		Name: "test",
	}
	user["addr"] = addr
	//user的value是空接口，所以没法直接拿到addr结构体，和hobby切片的值，需要类型断言
	userAddress, ok := user["addr"].(Address)
	if ok {
		fmt.Println(userAddress.Name)
	}

	userHobby, ok := user["hobby"].([]string)
	if ok {
		fmt.Println(userHobby[0])
	}
}
