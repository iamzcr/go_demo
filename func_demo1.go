package main

import (
	"fmt"
	"sort"
)

// 函数的作用域
// 全局变量，全局变量是定义在函数外部的变量，在程序整个运行周期内都有效
// 局部变量是函数内部定义的变量，函数内定义的变量无法在函数外使用
// 函数中切片传参
func sortSlice(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}

	return slice
}

// 函数中map传参
func sortMap(numMap map[int]int) (str string) {
	var keySlice []int
	for k, _ := range numMap {
		keySlice = append(keySlice, k)
	}
	sort.Ints(keySlice)
	for _, v := range keySlice {
		str += fmt.Sprintf("numMap排序后key值=%v，value值=%v", v, numMap[v])
	}
	return
}
func main() {
	var slice = []int{1, 5, 2, 5, 8, 8, 9}
	fmt.Printf("slice长度%v，容量%v，初始值%v，首个元素的内存地址%p\n", len(slice), cap(slice), slice, &slice[0])
	newSlice := sortSlice(slice)
	fmt.Printf("slice长度%v，容量%v，初始值%v，首个元素的内存地址%p\n", len(slice), cap(slice), slice, &slice[0])
	fmt.Printf("slice长度%v，容量%v，初始值%v，首个元素的内存地址%p\n", len(newSlice), cap(newSlice), newSlice, &newSlice[0])

	numMap := make(map[int]int)
	numMap[6] = 88
	numMap[2] = 55
	numMap[8] = 5
	fmt.Printf("numMap值==%v\n", numMap)
	newMap := sortMap(numMap)
	fmt.Printf("newMap值==%v\n", newMap)
	fmt.Printf("numMap值==%v\n", numMap)

}
