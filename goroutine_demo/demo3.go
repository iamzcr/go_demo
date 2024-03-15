package main

import (
	"fmt"
	"runtime"
)

// 设置golang并行运行的时候占用的cpu数量
// go运行时的调度器使用GOMAXPROCS参数来确定需要使用多少个OS线程来同时执行go代码，默认值时机器上的cpu核心数，例如一个8核的机器上，调度器会把go代码同时调度到8个OS线程上
// go中可以通过runtime.GOMAXPROCS函数设置当前程序并发时占用的cpu逻辑核心数
// 1.5版本前，默认使用的时单核心执行，1.5版本以后，默认使用全部cpu逻辑核心数
func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum=", cpuNum)
	runtime.GOMAXPROCS(cpuNum - 1)

}
