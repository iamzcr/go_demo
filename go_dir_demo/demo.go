package main

import (
	"fmt"
	"io/ioutil"
)

// 复制文件
func main() {
	slice, err := ioutil.ReadFile("./read.txt")
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile("./write.txt", slice, 0666)
	if err != nil {
		fmt.Println(err)
	}
}
