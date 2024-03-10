// main.go

package main

//一个包里面的多个文件的方法不能重复
import (
	"go_mod_demo/pkg1"
	P2 "go_mod_demo/pkg2" //引入后起别名
	_ "go_mod_demo/pkg3"  //引入不使用，匿名引入
)

func main() {
	pkg1.Func1()
	P2.Func2()

}
