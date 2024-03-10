package main

import (
	"fmt"
	"unsafe"
)

// 打印输出
func main() {
	//Println,Print,Printf
	/**
	Println换行，Print不换行，Printf格式化输出,默认不换行
	*/
	fmt.Println("hello word") //
	fmt.Print("hello word")   //
	var a = "a"
	fmt.Printf("hello word%v=====%p=====%T", a, &a, a)
	//%[index]verb 的形式允许你指定一个参数索引来与特定的格式化动词配对。索引是从1开始的，而不是从0开始。所以，%[1]v 表示使用第一个参数（即 num3）与 v 动词配对。
	fmt.Printf("num3 数据类型为%T，值为%[1]v，占用存储空间为%v字节\n", a, unsafe.Sizeof(a))

}
