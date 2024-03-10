package main

import (
	"encoding/json"
	"fmt"
)

//结构体对象转换为json字符串，属性必须是公有的
//结构体对象转换为json字符串，属性必须是公有的

type Stu struct {
	ID     int    `json:"id"`
	Sex    int    `json:"sex"`
	Name   string `json:"name"`
	Number string `json:"number"`
}
type Class struct {
	ClassName string
	Students  []Student
}

func main() {
	//结构体转换为json字符串
	var stu = Stu{
		ID:     1,
		Sex:    2,
		Name:   "小明",
		Number: "N123456",
	}
	stuByte, _ := json.Marshal(stu)
	stuJson := string(stuByte)
	fmt.Printf("%v\n", stuJson)

	//json字符串转为结构体
	var stuStr = `{"ID":1,"Sex":2,"Name":"小明","Number":"N123456"}`
	var stuStruct Stu
	//应该要修改stuStruct的值，所以需要传入地址
	err := json.Unmarshal([]byte(stuStr), &stuStruct)
	if err != nil {
		fmt.Printf("%v\n", err)

	}
	fmt.Printf("%#v\n", stuStruct)

}
