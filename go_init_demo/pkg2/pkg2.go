package pkg2

import (
	"fmt"
)

var Pkg2Val int

func init() {
	fmt.Println("this pkg2 init")
}

func Pkg2Print() {
	fmt.Println("this pkg2 Pkg2Print")
}
