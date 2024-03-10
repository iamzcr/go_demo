package main

import "fmt"

// 流程控制
func main() {
	age := 30 //age是当前区域内的全局变量
	if age > 20 {
		fmt.Println("true")
	}
	if age >= 30 {
		fmt.Println("30....true")
	} else if age > 20 {
		fmt.Println("20....true")
	} else {
		fmt.Println("true")

	}
	//age1是if语句内的局部变量
	if age1 := 32; age1 > 20 {
		fmt.Println("true")
	}

	//i当前区域内的全局变量
	i, j, n := 0, 0, 0
	for ; i < 3; i++ {
		fmt.Println("i=", i)
	}
	for j < 3 {
		fmt.Println("j=", j)
		j++
	}
	//i是for语句内的局部变量
	for i := 0; i < 3; i++ {
		fmt.Println("i=", i)
	}
	//无限循环，break中断
	for {
		if n < 3 {
			fmt.Println("n=", n)

		} else {
			break //跳出当前循环
		}
		n++
	}

	str := "这是一个测试，test"
	for index, item := range str {
		fmt.Printf("索引%v,原样输出%c，值%v,内存地址%p\n", index, item, item, &item)
	}

	var ext = ".html"
	switch ext {
	case ".html":
		fmt.Println("html")
		break //golang中break可以不写
	case ".php", ".go": //多个case值中间使用英文逗号分割
		fmt.Println("php")
	default:
		fmt.Println("找不到")

	}
	//ext是switch语句内的局部变量
	switch ext := ".html"; ext {
	case ".html":
		fmt.Println("html")
	case ".php":
		fmt.Println("php")
	default:
		fmt.Println("找不到")

	}

	var age1 = 32
	switch {
	case age1 < 24: //如果条件里面有表达式，那么switch后面不写判断变量
		fmt.Println("<24")
		break //golang中break可以不写
	default:
		fmt.Println("找不到")
	}

	var score = "A"
	switch score {
	case "A":
		fmt.Println("优秀")
		fallthrough //fallthrough可以执行满足条件的case的下个case，一次只能穿透一层，离最近的case
	case "B":
		fmt.Println("良好")
	default:
		fmt.Println("找不到")
	}

	//label跳出多重循环
label1:
	for i := 0; i < 2; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break label1
			}
			fmt.Printf("i=%v,j=%v\n", i, j)
		}
	}

	//continue结束某一次循环，开始下一次循环迭代
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
	//continue配合label，跳出里层循环
label2:
	for i := 0; i < 2; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				continue label2
			}
			fmt.Printf("i=%v,j=%v\n", i, j)
		}
	}

	//goto语句通过标签进行代码间无条件跳转，可以在快速跳出循环，避免重复推出上有一定的帮助，能简化一些代码的实现过程

	var num = 30
	if num > 18 {
		fmt.Printf("成年")

		goto label3
	}

	fmt.Println("未执行goto，label,跳过该语句")
label3:
	fmt.Println("执行goto，label")

}
