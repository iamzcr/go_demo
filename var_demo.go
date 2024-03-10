package main

import "fmt"

// 全局变量
var globalVal = "globalVal"

func main() {
	////声明变量
	//var test int
	////赋值
	//test = 22
	//fmt.Printf("test.....%v,%T,%p\n", test, test, &test)
	//
	////声明并且赋值
	//var test1 = "test1"
	//fmt.Printf("test1...%v,%T,%p\n", test1, test1, &test1)
	//
	//var test2 string = "test2"
	//fmt.Printf("test2...%v,%T,%p\n", test2, test2, &test2)
	//
	////声明并且赋值
	//test3 := "test3"
	//fmt.Printf("test...%v,%T,%p\n", test3, test3, &test3)
	//
	////声明多个不同类型的表里
	//var (
	//	test4 string
	//	test5 int
	//)
	//fmt.Printf("test4...%v,%T,%p\n", test4, test4, &test4)
	//fmt.Printf("test5...%v,%T,%p\n", test5, test5, &test5)
	//
	////声明多个不同类型的表里
	//var (
	//	test6 string = "test"
	//	test7 int    = 100
	//)
	//fmt.Printf("test6...%v,%T,%p\n", test6, test6, &test6)
	//fmt.Printf("test7...%v,%T,%p\n", test7, test7, &test7)
	//
	////声明多个同类型变量
	//var test8, test9 string
	//
	//fmt.Printf("test8....%v,%T,%p\n", test8, test8, &test8)
	//fmt.Printf("test9...%v,%T,%p\n", test9, test9, &test9)
	////声明并赋值多个同类型变量
	//test10, test11 := 11, 12
	//fmt.Printf("test10....%v,%T,%p\n", test10, test10, &test10)
	//fmt.Printf("test11...%v,%T,%p\n", test11, test11, &test11)

	//常量
	//常量值不可变,声明一般用大写。
	//常量的值是在编译时确定的，它们在程序运行时是不可变的。所以在运行时无法获取常量的内存地址。
	const PI = "3014.158"
	fmt.Printf("PI...%v,%T\n", PI, PI)

	//多个常量一起声明。
	const (
		A = "A"
		B = 1
	)
	fmt.Printf("A...%v,%T...B...%v,%T\n", A, A, B, B)

	//多个常量一起声明，如果第一个赋值，其他不赋值，其他的值和第一个常量的值一样。
	const (
		C = "C"
		D
		E
	)
	fmt.Printf("C...%v,%T...D...%v,%T...E...%v,%T\n", C, C, D, D, E, E)

	//iota关键字
	//iota是golang的常量计数器，只能在常量的表达式中使用
	const iotaTest = iota
	fmt.Printf("iotaTest...%v,%T\n", iotaTest, iotaTest)

	//多个常量定义，如果首个常量是iota，那么默认是0，后面常量会自动加1
	const (
		iota1 = iota
		iota2
		iota3
	)
	fmt.Printf("iota常量...%v,%v,%v\n", iota1, iota2, iota3)

	//使用_跳过某些值
	const (
		iota4 = iota
		_
		iota5
	)
	fmt.Printf("iota跳过某些值...%v,%v\n", iota4, iota5)
	//使用iota声明时候，中间插入一个初始化
	const (
		iota6 = iota
		iota7 = 100
		iota8 = iota
		iota9
	)
	fmt.Printf("iota中间插入...%v,%v,%v,%v\n", iota6, iota7, iota8, iota9)

	//多个iota定义在一行，其中的递增计算，第二行的声明列数和第一行一致，否则编译报错,所以可理解为const语句块中的行索引。
	const (
		iota10, iota11 = iota + 1, iota + 3
		iota12, iota13
	)
	fmt.Printf("iota多个iota定义在一行...%v,%v,%v,%v\n", iota10, iota11, iota12, iota13)

}
