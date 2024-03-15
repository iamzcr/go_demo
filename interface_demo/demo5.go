package main

import "fmt"

type Aminal interface {
	SetName(string)
	GetName() string
}

type Cat struct {
	Name string
}

func (c *Cat) SetName(name string) {
	c.Name = name
}
func (c *Cat) GetName() (name string) {
	name = c.Name
	return
}
func main() {
	var aminal Aminal = &Cat{}
	aminal.SetName("big cat")
	fmt.Println("cat的名字是：", aminal.GetName())
}
