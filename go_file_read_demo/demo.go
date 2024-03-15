package main

//file.Read读取
import (
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

	//每次读取128字节
	var tempSlice = make([]byte, 128)
	// 存放总的读取记录
	var strSlice []byte
	for {
		num, err := file.Read(tempSlice)
		if err == io.EOF { //io.EOF表示文件读取完毕
			fmt.Println("文件读取完成")
			break
		}
		if err != nil {
			fmt.Println("读取文件失败：", err)
			return
		}
		fmt.Printf("读取了%v个字节\n", num)
		//拼接切片的时候，如果最后一次读取不够128个字节，会出现问题，所以应该取每次读出来的字节数切割
		strSlice = append(strSlice, tempSlice[:num]...)

	}
	fmt.Println(string(strSlice))
}
