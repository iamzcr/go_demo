package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//一次性读取，不推荐
	slice, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(slice))

}
