package main

// 结构体匿名字段（正常定义不建议这么做）
// 声明中只有类型，没有字段名
// 匿名字段采用默认类型名作为字段名，结构体要求字段名称必须唯一
type Test struct {
	string
	int
}

// 结构体嵌套
type Animal struct {
}

func main() {

}
