package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"unsafe"
)

func main() {

	num1, num2 := 8.2, 2.8
	sub := decimal.NewFromFloat(num1).Sub(decimal.NewFromFloat(num2))
	fmt.Printf("float9 数据类型为%T，值为%v，占用存储空间为%v字节\n", sub, sub, unsafe.Sizeof(sub))
}
