package main

import (
	"fmt"
)

func testPointFunc(num *int) {
	*num = 40
}

// 指针也是一个变量，存储的数据不是一个普通的值，而是另外一个变量的内存地址,所以实际上指针就是一个引用类型
//*point表示取出point这个指针变量对应的num的内存地址的值,point表示num对应的内存地址

func main() {

	var num int = 10
	//num地址是
	fmt.Printf("num值是%v,类型是%T,地址是%p\n", num, num, &num)
	//声明一个指针类型的变量pointNum，把num的地址赋值给pointNum
	var point *int = &num
	//指针变量point的值就是num地址，而指针变量point也有自己的内存地址
	fmt.Printf("point值是%v，类型是%T,地址是%p\n", point, point, &point)

	//通过指针变量point可以修改num的值
	*point = 20
	fmt.Printf("num值是%v,类型是%T,地址是%p\n", num, num, &num)

	testPointFunc(point)
	fmt.Printf("num值是%v,类型是%T,地址是%p\n", num, num, &num)

	//new和make函数
	//共同点：两者都是用来分配内存的,两者都有自己的内存地址
	//用途不同：
	//make 用于创建切片、映射和通道等引用类型的对象。
	//new 用于创建并分配内存空间，并且内存对应的值为类型零值，返回指向该内存空间的指针。
	//返回值类型不同：
	//make 返回的是已初始化的对象，即对应引用类型的实例。
	//new 返回的是指向分配的内存空间的指针。

	//make创建切片引用类型的对象
	var mapTest = make([]string, 0)
	fmt.Printf("num值是%v,类型是%T,num地址是%p\n", mapTest, mapTest, &mapTest)

	//new创建并分配内存空间，返回指向该内存空间的指针
	var num1 *int
	num1 = new(int)
	*num1 = 10
	fmt.Printf("num1值是%v,类型是%T,num2地址是%p\n", num1, num1, &num1)

	var num2 = new(int)
	fmt.Printf("num2值是%v,类型是%T,num2地址是%p\n", num2, num2, &num2)

}
