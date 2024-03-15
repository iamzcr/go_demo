package main

//接口的嵌套
import "fmt"

type AminalTop interface {
	AminalSub
	AminalSub1
}
type AminalSub interface {
	SetName(string)
}
type AminalSub1 interface {
	GetName() string
}
type CatSub struct {
	Name string
}

func (c *CatSub) SetName(name string) {
	c.Name = name
}
func (c *CatSub) GetName() (name string) {
	name = c.Name
	return
}
func main() {
	var aminalTop AminalTop = &CatSub{}
	aminalTop.SetName("big cat")
	fmt.Println("cat的名字是：", aminalTop.GetName())
}
