package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		err := file.Close()
		if err != nil {

		}
	}()
	var fileStr string

	scanner := bufio.NewScanner(file) //创建了一个 scanner 对象

	for scanner.Scan() { //scanner.Scan() 逐行扫描文件内容
		line := scanner.Text() //scanner.Text() 返回当前行的文本内容,这个方法会把文件本身的换行去掉
		fileStr += (line + "\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("读取文件失败：", err)
		return
	}
	fmt.Println(fileStr)

}
