package main

import (
	"fmt"
	"time"
)

func main() {
	timeObj := time.Now()
	//输出年份
	fmt.Println(timeObj.Year())
	//12小时制
	fmt.Println(timeObj.Format("2006-01-02 03:04:05"))
	//24小时制
	fmt.Println(timeObj.Format("2006-01-02 15:04:05"))
	//格式化模板
	fmt.Println(timeObj.Format("2006/01/02 15:04:05"))
	//当前时间戳 秒
	fmt.Println(timeObj.Unix())
	//当前时间戳 纳秒
	fmt.Println(timeObj.UnixNano())
	//当前时间戳 微妙
	fmt.Println(timeObj.UnixMilli())
	//当前时间戳 微妙
	fmt.Println(timeObj.UnixMicro())

}
