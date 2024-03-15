package main

// bufio.NewReader读取
import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//只读方式打开文件
	file, err := os.Open("./test.txt")
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
	var fileStr string

	reader := bufio.NewReader(file) //资源句柄
	for {
		str, err := reader.ReadString('\n') //一次读取一行
		if err == io.EOF {                  //io.EOF表示文件读取完毕
			fileStr += str //如果最后一行没有换行，可能会存在最后一行读取不到的情况，所以在程序认为读取结束之后再赋值一下
			fmt.Println("文件读取完成")
			break
		}
		if err != nil {
			fmt.Println("读取文件失败：", err)
			return
		}
		fileStr += str
	}
	fmt.Println(fileStr)
}
