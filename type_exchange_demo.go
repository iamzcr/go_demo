package main

import (
	"fmt"
	"strconv"
)

// 基本数据类型中转换
// 整型,浮点型之间可以相互显式转换,但是需要注意精度
// 整型,浮点型,字符型,布尔型转可以换成字符串,有两种方式,一种是strconv库操作,还可以通过fmt库的Sprintf函数操作
func main() {
	//整型换成字符串
	var num = 64
	numStr := fmt.Sprintf("%d", num)
	fmt.Printf("值%v,类型%T\n", numStr, numStr)
	//浮点换成字符串
	var float1 = 125.6
	floatStr := fmt.Sprintf("%f", float1)
	fmt.Printf("值%v,类型%T\n", floatStr, floatStr)
	//布尔类型换成字符串
	var bool1 = true
	boolStr := fmt.Sprintf("%t", bool1)
	fmt.Printf("值%v,类型%T\n", boolStr, boolStr)
	//字符换成字符串
	var byte1 = 'a'
	byteStr := fmt.Sprintf("%c", byte1)
	fmt.Printf("值%v,类型%T\n", byteStr, byteStr)

	//strconv整型int64转字符串,第二个参数表示格式化为什么进制
	numStr1 := strconv.FormatInt(int64(num), 10)
	fmt.Printf("strconv转换,值%v,类型%T\n", numStr1, numStr1)
	//strconv整型int转字符串
	numStr2 := strconv.Itoa(num)
	fmt.Printf("strconv转换,值%v,类型%T\n", numStr2, numStr2)
	//strconv浮点转字符串
	floatStr1 := strconv.FormatFloat(float1, 'f', -1, 64)
	fmt.Printf("strconv转换,值%v,类型%T\n", floatStr1, floatStr1)
	//strconv字符转字符串,第二个参数表示格式化为什么进制
	byteStr1 := strconv.FormatUint(uint64(byte1), 10)
	fmt.Printf("strconv转换,值%v,类型%T\n", byteStr1, byteStr1)
	//strconv字符转字符串,第二个参数表示格式化为什么进制
	boolStr2 := strconv.FormatBool(bool1)
	fmt.Printf("strconv转换,值%v,类型%T\n", boolStr2, boolStr2)

	//字符串转浮点类型
	var str = "123.5666"
	float2, _ := strconv.ParseFloat(str, 64)
	fmt.Printf("strconv字符串转换,值%v,类型%T\n", float2, float2)

	//字符串转整型int和转int64
	var str1 = "123"
	int1, _ := strconv.Atoi(str1)
	fmt.Printf("strconv字符串转换,值%v,类型%T\n", int1, int1)
	//第一个参数是进制,第二个是格式化成位数
	int2, _ := strconv.ParseInt(str1, 10, 64)
	fmt.Printf("strconv字符串转换,值%v,类型%T\n", int2, int2)

	//字符串转字符
	var str2 = "a"
	byte2, _ := strconv.ParseUint(str2, 10, 64)
	fmt.Printf("strconv字符串转换,值%v,类型%T\n", byte2, byte2)

	//字符串转字符
	var str3 = "true"
	bool2, _ := strconv.ParseBool(str3)
	fmt.Printf("strconv字符串转换,值%v,类型%T\n", bool2, bool2)
}
