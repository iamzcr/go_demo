package main

import (
	"fmt"
	"os"
)

// 创建目录
func main() {
	//如果目录存在，会报错
	//err := os.Mkdir("./test", 0666)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//err := os.MkdirAll("./test/test1/test", 0666)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	////删除文件
	//err = os.Remove("./text.txt")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	////删除目录
	//err = os.RemoveAll("./test/test1/test")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//err = os.Remove("./test")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//重命名
	err := os.Rename("./write.txt", "./test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
}
