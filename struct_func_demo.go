package main

import "fmt"

// Student 结构体方法
type Student struct {
	name string
	age  int
	sex  int
}

// (方法接受者变量，接收者类型)方法名(参数列表)(返回参数)

// PrintInfo 接收者类型是结构体
func (s Student) PrintInfo() {
	fmt.Printf("Student值是%v,全部信息是%#v,类型是%T,地址是%p\n", s, s, s, &s)

}

// PrintPointInfo 接收者类型是结构体指针
func (s *Student) PrintPointInfo() {
	fmt.Printf("Student值是%v,全部信息是%#v,类型是%T,地址是%p\n", s, s, s, &s)

}

func (s *Student) SetPointInfo(name string, age int) {
	s.name = name
	s.age = age

}
func main() {
	//实例化一个person对象person1
	var stu = Student{
		name: "小李",
		age:  1,
		sex:  1,
	}
	//使用person1调用对象的方法
	stu.PrintPointInfo()
	stu.PrintInfo()

	//实例化一个person指针变量对象person2
	var stu1 = &Student{
		name: "小张",
		age:  1,
		sex:  1,
	}
	//使用person2调用对象的方法
	stu1.PrintPointInfo()
	stu1.PrintInfo()

	stu1.SetPointInfo("小明", 555)

	stu1.PrintPointInfo()
	stu1.PrintInfo()
}
