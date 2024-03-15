package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	file, err := os.OpenFile("./text.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
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
	var str = "test is test"
	_, err = file.Write([]byte(str))
	if err != nil {
		return
	}
	_, err = file.WriteString(str + "\r\n")
	if err != nil {
		return
	}
	for i := 0; i < 10; i++ {
		_, err = file.WriteString(str + strconv.Itoa(i) + "\r\n")
		if err != nil {
			return
		}
	}
	var bufioStr = "test is test bufio"
	//bufio写入
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(bufioStr) //先写入缓冲区
	if err != nil {
		return
	}
	writer.Flush() //writer.Flush()执行该方法，真正写入文件

	err = ioutil.WriteFile("./text.txt", []byte(bufioStr), 0666)
	if err != nil {
		fmt.Println(err)
	}

}
