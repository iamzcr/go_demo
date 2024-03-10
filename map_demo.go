package main

import (
	"fmt"
	"sort"
)

// map是一种无需的基于key-value的数据结构，是一种引用类型
func main() {
	//make创建声明map，第二个参数可以指定长度
	var user = make(map[string]string)
	fmt.Printf("值=%v,类型=%T\n", user, user)

	user["name"] = "小明"
	user["age"] = "10"
	fmt.Printf("值=%v,类型=%T\n", user, user)
	//直接初始化
	var user1 = map[string]string{
		"name": "小明1",
	}
	fmt.Printf("值=%v,类型=%T\n", user1, user1)

	user2 := map[string]string{
		"name": "小明1",
	}
	fmt.Printf("值=%v,类型=%T,name=%v\n", user2, user2, user2["name"])

	//遍历map
	for k, v := range user {
		fmt.Printf("k=%v,k=%T,v=%v,v=%T\n", k, k, v, v)
	}

	//查找,如果没有这个key，ok是false，v为对应数据类型的初始值
	v, ok := user["name"]
	fmt.Printf("v=%v,ok=%v\n", v, ok)

	//修改map，修改前先查找
	if _, ok := user["age"]; ok {
		user["name"] = "晓明"
		fmt.Printf("修改后值=%v,类型=%T,name=%v\n", user, user, user["name"])
	}

	//删除,删除前最好先查找一下
	if _, ok := user["age"]; ok {
		delete(user, "age")
		fmt.Println("删除user的age")
	}
	fmt.Printf("修改后值=%v,类型=%T,name=%v\n", user, user, user["name"])

	//map和切片
	//切片的元素的是map
	sliceMap := make([]map[string]string, 1)
	fmt.Printf("添加前值=%v，首个元素内存地址%p\n", sliceMap, &sliceMap[0])
	userMap := map[string]string{
		"name": "小张",
	}
	mapSlice1 := append(sliceMap, userMap)
	fmt.Printf("添加后值=%v，首个元素内存地址%p\n", mapSlice1, &mapSlice1[0])

	mapSlice2 := append(sliceMap, map[string]string{
		"name": "小李",
	})
	fmt.Printf("添加后值=%v，首个元素内存地址%p\n", mapSlice2, &mapSlice2[0])

	//map的value是切片
	mapSlice := make(map[string][]string)

	mapSlice["name"] = []string{
		"小张",
		"小李",
	}
	mapSlice["work"] = []string{
		"php",
		"golang",
	}
	fmt.Printf("添加后值=%v\n", mapSlice)

	//map是引用数据类型
	var mapTemp = make(map[string]string)
	mapTemp["name"] = "小张"
	fmt.Printf("修改mapCopy前mapTemp值=%v\n", mapTemp)
	mapCopy := mapTemp
	mapCopy["name"] = "小李"
	fmt.Printf("修改mapCopy后mapTemp值=%v\n", mapTemp)

	//map是无序的，要排序，可以借助切片来排序

	numMap := make(map[int]int)
	numMap[6] = 88
	numMap[2] = 55
	numMap[8] = 5
	var keySlice []int
	for k, _ := range numMap {
		keySlice = append(keySlice, k)
	}
	fmt.Println(keySlice)
	sort.Ints(keySlice)
	fmt.Println(keySlice)
	for _, v := range keySlice {
		fmt.Printf("numMap排序后key值=%v，value值=%v\n", v, numMap[v])
	}

}
