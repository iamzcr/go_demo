package main

import (
	"fmt"
	"go_init_demo/pkg1"
	"go_init_demo/pkg2"
	"go_init_demo/pkg3"
)

// golang的程序执行时会自动触发包内部的init()函数的调用，该函数没有参数，也没有返回值，不能被主动调用。
// 包初始化执行顺序：全局声明->init()->main()函数。
// 在运行时，不在同一个文件导入的包调用顺序是被最后导入的包会最先初始化并调用其init()函数，如果在同一个文件被导入的包，那么就按导入顺序执行。
// 同一个包可以有多个init()函数,按照它们在源代码中的顺序被调用。
// 如果同一个包被拆分为多个文件，那么每个文件可以包含一个 init 函数。这些 init 函数的执行顺序是无法保证的，因此在设计时应避免互相依赖或者需要确定的初始化顺序。
var MainVar int

func init() {
	fmt.Println("main init")
}

func init() {
	fmt.Println("main init1")
}

func main() {
	fmt.Println("main start")
	fmt.Println(pkg1.Pkg1Val)
	fmt.Println(pkg2.Pkg2Val)
	fmt.Println(pkg3.Pkg3Val)

}
