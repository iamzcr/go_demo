package pkg1

import (
	"fmt"
	"go_init_demo/pkg2"
)

var Pkg1Val int

func init() {
	fmt.Println("this pkg1 init")
}

func Pkg1Print() {
	pkg2.Pkg2Print()
}
