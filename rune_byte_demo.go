package main

import (
	"fmt"
	"unsafe"
)

//字符,组成每个字符串的元素叫做字符,可以通过遍历字符串获得字符,
//golang的字符类型有byte和rune类型
//byte代表一个ascii码值的一个字符,实际上是一个无符号的int8类型uint8,也就是和uint8声明的变量可以相互赋值,默认是0,对应的符号是空,占用1个字节存储空间
//rune类型,代表一个unicode编码字符,实际类型是int32,也就是和int32声明的变量可以相互赋值,默认是0,对应的符号是空,占用4个字节的存储空间

func main() {
	var defaultRune rune
	fmt.Printf("defaultRune 数据类型为%T，ascii码值为%v，输出字符值为%c，占用存储空间为%v字节\n", defaultRune, defaultRune, defaultRune, unsafe.Sizeof(defaultRune))

	var defaultByte byte
	fmt.Printf("defaultByte 数据类型为%T，ascii码值为%v，输出字符值为%c，占用存储空间为%v字节\n", defaultByte, defaultByte, defaultByte, unsafe.Sizeof(defaultByte))

	//声明初始化byte类型字符
	//具体参考https://www.w3school.com.cn/charsets/ref_html_ascii.asp
	var byte1 byte = 'a'
	fmt.Printf("byte1 数据类型为%T，ascii码值为%v，输出字符值为%c，占用存储空间为%v字节\n", byte1, byte1, byte1, unsafe.Sizeof(byte1))

	//隐式声明一个字符,默认是rune类型的字符,也就是int32,占用4个字节存储空间.
	rune1 := 'a'
	fmt.Printf("byte2 数据类型为%T，unicode编码值为%v，输出字符值为%c，占用存储空间为%v字节\n", rune1, rune1, rune1, unsafe.Sizeof(rune1))

	rune2 := '我'
	fmt.Printf("byte2 数据类型为%T，unicode编码值为%v，输出字符值为%c，占用存储空间为%v字节\n", rune2, rune2, rune2, unsafe.Sizeof(rune2))

	//字符串的长度是指字节的长度
	var str = "我好"
	fmt.Println("str字节长度:", len(str))
	//byte类型
	byteSlice := []byte(str)
	for i := 0; i < len(byteSlice); i++ {
		fmt.Printf("数据类型为%T，ascii码值为%v，输出字符值为%c，占用存储空间为%v字节\n", byteSlice[i], byteSlice[i], byteSlice[i], unsafe.Sizeof(byteSlice[i]))
	}
	fmt.Println("===========")
	for i := 0; i < len(str); i++ {
		fmt.Printf("数据类型为%T，ascii码值为%v，输出字符值为%c，占用存储空间为%v字节\n", str[i], str[i], str[i], unsafe.Sizeof(str[i]))
	}
	fmt.Println("===========")

	//rune类型
	runeSlice := []rune(str)
	for i := 0; i < len(runeSlice); i++ {
		fmt.Printf("数据类型为%T，unicode编码值为%v，输出字符值为%c，占用存储空间为%v字节\n", runeSlice[i], runeSlice[i], runeSlice[i], unsafe.Sizeof(runeSlice[i]))
	}
	fmt.Println("===========")
	for _, r := range str {
		fmt.Printf("数据类型为%T，unicode编码值为%v，输出字符值为%c，占用存储空间为%v字节\n", r, r, r, unsafe.Sizeof(r))
	}

	//rune类型或者byte类型切片转换字符串,转换后分配新的内存地址
	byteStr := string(byteSlice)
	runeStr := string(runeSlice)
	fmt.Printf("源字符串内存地址:%p,新字符串数据类型为:%T，值为:%v，内存地址值为:%p\n", &str, byteStr, byteStr, &byteStr)
	fmt.Printf("源字符串内存地址:%p,新字符串数据类型为:%T，值为:%v，内存地址值为:%p\n", &str, runeStr, runeStr, &runeStr)

}
