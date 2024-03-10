package main

import (
	"fmt"
	"math"
	"unsafe"
)

// golang中浮点类型有float32和float64两种
// float64浮点数类型:不显式声明,含有小数的都是默认float64,占用8个字节内存,默认保留6位小数,如果没有为科学计数法形式的数指定类型，那么会被视为float64类型
// float32浮点数类型:占用4个字节内存,默认保留6位小数,可以节省内存，但是math标准库里面的函数操作都是float64，所以实际开发应该选择放float64
func main() {
	//浮点型，浮点型只声明，不赋值默认是0
	var float float64
	fmt.Printf("float 数据类型为%T，值为%v，占用存储空间为%v字节\n", float, float, unsafe.Sizeof(float))
	//只要含有小数，都是float64类型
	float1 := 4.356
	fmt.Printf("float1 数据类型为%T，值为%v，占用存储空间为%v字节\n", float1, float1, unsafe.Sizeof(float1))

	var float2 = 4.356
	fmt.Printf("float2 数据类型为%T，值为%v，占用存储空间为%v字节\n", float2, float2, unsafe.Sizeof(float2))

	var float3 float64 = 4.356
	fmt.Printf("float3 数据类型为%T，值为%v，占用存储空间为%v字节\n", float3, float3, unsafe.Sizeof(float3))

	//默认保留6位小数
	var float4 float32 = 4.356
	var float5 = 4.356
	fmt.Printf("float4 数据类型为%T，值为%v，默认保留小数为%f，占用存储空间为%v字节\n", float4, float4, float4, unsafe.Sizeof(float4))
	fmt.Printf("float5 数据类型为%T，值为%v，默认保留小数为%f，占用存储空间为%v字节\n", float5, float5, float5, unsafe.Sizeof(float5))
	//保留两位小数
	fmt.Printf("float5 数据类型为%T，值为%v，默认保留小数为%.2f，占用存储空间为%v字节\n", float5, float5, float5, unsafe.Sizeof(float5))

	//如果没有为科学计数法形式的数指定类型，那么会被视为float64类型
	var float6 = 2e2 //2 * 10的二次方
	fmt.Printf("float6 数据类型为%T，值为%v，默认保留小数为%f，占用存储空间为%v字节\n", float6, float6, float6, unsafe.Sizeof(float6))

	//精度问题,浮点型的计算可能会丢失精度的问题
	num1, num2 := 8.2, 2.8
	float9 := num1 - num2
	fmt.Printf("float9 数据类型为%T，值为%v，占用存储空间为%v字节\n", float9, float9, unsafe.Sizeof(float9))

	fmt.Println("float9=====", float9 == 5.4)
	//理论上是true的，由于浮点数的精度问题，返回false，所以使用上要小心，如果确实需要比较，可以通过定义一个容差值,如果计算出来的值和我们预想的值相减的绝对值少于容差值，那么就认为足够接近预想的值
	epsilon := 0.00001
	fmt.Println("float9容差值=====", math.Abs(float9-5.4) < epsilon)

	//布尔类型
	//布尔类型类型只有true和false两个值
	//布尔类型变量的默认值位false
	//不允许将整型强制转换位布尔型
	//布尔型没法参与数值的运算,也无法与其他类型进行转换

	//声明,不赋值,默认flase
	var flag bool
	fmt.Printf("flag 数据类型为%T，值为%v，占用存储空间为%v字节\n", flag, flag, unsafe.Sizeof(flag))

	var flag1 = true
	fmt.Printf("flag1 数据类型为%T，值为%v，占用存储空间为%v字节\n", flag1, flag1, unsafe.Sizeof(flag1))

}
