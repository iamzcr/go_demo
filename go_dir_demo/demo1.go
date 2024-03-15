package main

import (
	"fmt"
	"os"
)

// 复制文件
func main() {
	file, err := os.Open("./read.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		//打开记得关闭
		err := file.Close()
		if err != nil {
			fmt.Println("打开文件失败：", err)
		}
	}()

	fileWrite, err := os.OpenFile("./text.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		//打开记得关闭
		err := fileWrite.Close()
		if err != nil {
			fmt.Println("打开文件失败：", err)
		}
	}()
	//每次读取128字节
	var tempSlice = make([]byte, 128)
	for {
		num, err := file.Read(tempSlice)
		if err != nil {
			fmt.Println("读取文件失败：", err)
			return
		}
		_, err = fileWrite.Write(tempSlice[:num])
		if err != nil {
			return
		}
	}
}
