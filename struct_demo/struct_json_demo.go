package main

import (
	"encoding/json"
	"fmt"
)

//结构体对象转换为json字符串，属性必须是公有的
//结构体对象转换为json字符串，属性必须是公有的

type Stu struct {
	ID     int      `json:"id"`
	Sex    int      `json:"sex"`
	Name   string   `json:"name"`
	Number string   `json:"number"`
	Hobby  []string `json:"hobby"`
	Addr   Address  `json:"address"`
}
type Address struct {
	AddrName string
	Phone    string
}

func main() {
	//结构体转换为json字符串
	var stu = Stu{
		ID:     1,
		Sex:    2,
		Name:   "小明",
		Number: "N123456",
		Hobby:  []string{"写代码", "足球"},
	}
	stu.Addr = Address{
		AddrName: "test",
		Phone:    "123456789",
	}
	//返回的是一个byte类型的切片
	stuByte, _ := json.Marshal(stu)
	stuJson := string(stuByte)
	fmt.Printf("%v\n", stuJson)

	//json字符串转为结构体
	var stuStr = `{"id":1,"sex":2,"name":"小明","number":"N123456","hobby":["写代码","足球"],"address":{"AddrName":"test","Phone":"123456789"}}`
	var stuStruct Stu
	//应该要修改stuStruct的值，所以需要传入地址
	err := json.Unmarshal([]byte(stuStr), &stuStruct)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("%#v\n", stuStruct)

}
