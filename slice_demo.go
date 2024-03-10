package main

import "fmt"

// 切片是一个引用类型,内部结构包括地址,长度,和容量,实际上切片是数组的一个引用
// 切片只声明,不赋值,默认值是一个nil
func main() {
	//切片只声明,不赋值,默认值是一个nil
	//var slice []int
	//fmt.Printf("sliceMake长度%v，容量%v，初始值%v\n", len(slice), cap(slice), slice)
	//if slice == nil {
	//	fmt.Println("slice is nil")
	//}
	////make声明和创建
	//sliceMake := make([]int, 0)
	//if sliceMake == nil {
	//	fmt.Println("sliceMake is nil")
	//}
	//fmt.Printf("sliceMake长度%v，容量%v，初始值%v\n", len(sliceMake), cap(sliceMake), sliceMake)
	//
	//sliceMake1 := make([]int, 5, 10)
	//fmt.Printf("sliceMake1长度%v，容量%v，初始值%v\n", len(sliceMake1), cap(sliceMake1), sliceMake1)
	//
	////直接声明和初始化
	//var areaSlice = []string{"广州", "北京", "上海", "深圳", "武汉", "重庆", "成都", "洛阳"}
	//fmt.Printf("areaSlice长度%v，容量%v，初始值%v\n", len(areaSlice), cap(areaSlice), areaSlice)
	//
	////数组切分创建和初始化切片
	//var areaArr = [...]string{"广州", "北京", "上海", "深圳", "武汉", "重庆", "成都", "洛阳"}
	////忽略掉slice的起始索引，表示从数组的起始位置进行切分
	//areaSlice1 := areaArr[:3]
	//fmt.Printf("areaSlice1长度%v，容量%v，初始值%v\n", len(areaSlice1), cap(areaSlice1), areaSlice1)
	//
	////忽略掉slice的结束索引，表示从使用数组的长度作为结束索引
	//areaSlice2 := areaArr[4:]
	//fmt.Printf("areaSlice2长度%v，容量%v，初始值%v\n", len(areaSlice2), cap(areaSlice2), areaSlice2)
	////同时忽略开始和结束索引，就是包含数组所有元素的一个slice
	//areaSlice3 := areaArr[:]
	//fmt.Printf("areaSlice3长度%v，容量%v，初始值%v\n", len(areaSlice3), cap(areaSlice3), areaSlice3)

	//range方式
	//var areaSlice = []string{"广州", "北京", "上海", "深圳", "武汉", "重庆", "成都", "洛阳"}
	//for index, item := range areaSlice {
	//	fmt.Printf("areaSlice遍历第%v个元素的值为%v\n", index, item)
	//}
	////for方式
	//for i := 0; i < len(areaSlice); i++ {
	//	fmt.Printf("areaSlice遍历第%v个元素的值为%v\n", i, areaSlice[i])
	//}

	//修改切片内容
	//var areaSlice = []string{"广州", "北京", "上海", "深圳", "武汉", "重庆", "成都", "洛阳"}
	//fmt.Printf("areaSlice修改前值%v\n", areaSlice)
	//areaSlice[0] = "洛阳"
	//fmt.Printf("areaSlice修改后值%v\n", areaSlice)
	//append新增切片元素
	//sliceMake := make([]int, 1)
	//fmt.Printf("sliceMake长度%v，容量%v，初始值%v，首个元素的内存地址%p\n", len(sliceMake), cap(sliceMake), sliceMake, &sliceMake[0])
	//sliceMakeAppend := append(sliceMake, 1, 2)
	//fmt.Printf("sliceMakeAppend长度%v，容量%v，初始值%v，首个元素的内存地址%p\n", len(sliceMakeAppend), cap(sliceMakeAppend), sliceMakeAppend, &sliceMakeAppend[0])
	//sliceMakeAppend1 := append(sliceMakeAppend, 1)
	//fmt.Printf("sliceMakeAppend1长度%v，容量%v，初始值%v，首个元素的内存地址%p\n", len(sliceMakeAppend1), cap(sliceMakeAppend1), sliceMakeAppend1, &sliceMakeAppend1[0])

	//append可以在切片声明之后直接追加，因为append底层实现有make方法，在不符合长度的时候，会初始化该切片。
	var slice []int
	if slice == nil {
		fmt.Printf("slice is nil\n")
	}
	slice = append(slice, 11)
	if slice != nil {
		fmt.Printf("slice is not nil\n")
	}
	fmt.Printf("slice长度%v，容量%v，初始值%v\n", len(slice), cap(slice), slice)

	//slice1 := []string{"广州", "北京"}
	//fmt.Printf("slice1长度%v，容量%v，初始值%v，首个元素的内存地址%p\n", len(slice1), cap(slice1), slice1, &slice1[0])
	//
	//slice2 := []string{"广州", "北京"}
	//fmt.Printf("slice2长度%v，容量%v，初始值%v，首个元素的内存地址%p\n", len(slice2), cap(slice2), slice2, &slice2[0])
	//
	//sliceMerge := append(slice1, slice2...)
	//fmt.Printf("sliceMerge长度%v，容量%v，初始值%v，首个元素的内存地址%p\n", len(sliceMerge), cap(sliceMerge), sliceMerge, &sliceMerge[0])
	//slice2[0] = "深圳"
	//fmt.Printf("sliceMerge长度%v，容量%v，初始值%v，首个元素的内存地址%p\n", len(sliceMerge), cap(sliceMerge), sliceMerge, &sliceMerge[0])
	//
	////切片复制
	//sliceCopy := []string{"广州", "北京"}
	//sliceNew := make([]string, 2)
	//copy(sliceNew, sliceCopy)
	//fmt.Printf("sliceMerge长度%v，容量%v，初始值%v，首个元素的内存地址%p\n", len(sliceCopy), cap(sliceCopy), sliceCopy, &sliceCopy[0])
	//fmt.Printf("sliceMerge长度%v，容量%v，初始值%v，首个元素的内存地址%p\n", len(sliceNew), cap(sliceNew), sliceNew, &sliceNew[0])

	////切片数组可以用于切分字符串
	//var strByte = "this is test"
	//testSlice := strByte[:4]
	//fmt.Printf("testSlice长度%v，初始值%v，首个元素的内存地址%p\n", len(testSlice), testSlice)
	//
	////切片字符串，索引代表的是字节数而非rune的数，例如下面的"go语言,"长度为9个字节，一个汉字占用3个字节
	//var strRune = "go语言,Hello！"
	//testSlice1 := strRune[:7]
	//fmt.Printf("testSlice1长度%v，初始值%v，首个元素的内存地址%p\n", len(testSlice1), testSlice1)
	//

	//var strByte = "this is test"
	//var strRune = "go语言,Hello！"
	//var tempByte = []byte(strByte)
	//fmt.Printf("tempByte长度%v，容量%v，初始值%v，首个元素的内存地址%p\n", len(tempByte), cap(tempByte), tempByte, &tempByte[0])
	//var newByteSlice = make([]byte, 0)
	//for i := 0; i < 4; i++ {
	//	newByteSlice = append(newByteSlice, tempByte[i])
	//}
	//newStr := string(newByteSlice)
	//fmt.Printf("newStr值%v，类型%T\n", newStr, newStr)
	//
	//var tempRune = []rune(strRune)
	//fmt.Printf("tempRune长度%v，容量%v，初始值%v，首个元素的内存地址%p\n", len(tempRune), cap(tempRune), tempRune, &tempRune[0])
	//var newRuneSlice = make([]rune, 0)
	//for i := 0; i < 7; i++ {
	//	newRuneSlice = append(newRuneSlice, tempRune[i])
	//}
	//newRuneStr := string(newRuneSlice)
	//fmt.Printf("newStr值%v，类型%T\n", newRuneStr, newRuneStr)

	//在1.2版本中，切片引入了三个索引的拆分操作

	//var areaSlice = []string{"广州", "北京", "上海", "深圳", "武汉", "重庆", "成都", "洛阳"}
	//areaSlice1 := areaSlice[1:3]
	//fmt.Printf("areaSlice1值%v，长度%v，容量%v，类型%T\n", areaSlice1, len(areaSlice1), cap(areaSlice1), areaSlice1)
	//areaSlice2 := areaSlice[1:3:5]
	//fmt.Printf("areaSlicea2值%v，长度%v，容量%v，类型%T\n", areaSlice2, len(areaSlice2), cap(areaSlice2), areaSlice2)

}
