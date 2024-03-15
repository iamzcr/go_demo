package main

//一个结构体实现多个接口
import "fmt"

type Aminal1 interface {
	SetName(string)
}
type Aminal2 interface {
	GetName() string
}

type Cat1 struct {
	Name string
}

func (c *Cat1) SetName(name string) {
	c.Name = name
}
func (c *Cat1) GetName() (name string) {
	name = c.Name
	return
}
func main() {
	var c = &Cat1{}
	var aminal1 Aminal1 = c
	var aminal2 Aminal2 = c
	aminal1.SetName("big cat")
	fmt.Println("cat的名字是：", aminal2.GetName())
}
