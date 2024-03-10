package main

import "fmt"

// 给自定义类型添加方法
// 自定义两个类型
type Celsius float64
type Fahrenheit float64

// 给Celsius添加ToFahrenheit方法
func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func main() {
	//实例化Celsius类型
	c := Celsius(25)
	//调用ToFahrenheit方法
	f := c.ToFahrenheit()
	fmt.Println(f) // 输出: 77
}
