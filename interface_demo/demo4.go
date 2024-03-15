package main

import "fmt"

// 结构体值接收者，和指针接收者实现接口的区别
// 值接收者
// 如果结构体中的方法是值接收者，那么实例化后的结构体值类型和结构体指针类型都可以赋值给接口变量
// 如果结构体中的方法是指针变量接收者，那么实例化后只能通过结构体指针类型赋值给接口变量
type Db interface {
	connect()
	query()
}
type Msql struct {
	Name string
}

func (m Msql) connect() { //值类型接收者
	fmt.Println(m.Name, "mysql结构体实现了Database接口的方法connect")
}
func (m Msql) query() { //值类型接收者
	fmt.Println(m.Name, "mysql结构体实现了Database接口的方法query")
}

// 结构体里面可以有自己的方法
func (m Msql) myfunc() { //值类型接收者
	fmt.Println(m.Name, "mysql结构体调用myfunc")
}

type Psql struct {
	Name string
}

func (p *Psql) connect() { //指针类型接收者
	fmt.Println(p.Name, "Pgsql结构体实现了Database接口的方法connect")
}
func (p *Psql) query() { //指针类型接收者
	fmt.Println(p.Name, "Pgsql结构体实现了Database接口的方法query")
}

func main() {

	var msql = Msql{
		Name: "mysql",
	}
	//msql作为值类型赋值给接口变量，并且结构体接收者是值接收者类型
	var db Db = msql
	db.connect()

	var msql1 = &Msql{
		Name: "mysql",
	}
	//msql作为指针类型赋值给接口变量，并且结构体接收者是值接收者类型
	var db1 Db = msql1
	db1.connect()

	var psql = &Psql{
		Name: "PGSQL",
	}
	//msql作为指针类型赋值给接口变量，并且结构体接收者是指针接收者类型
	var db2 Db = psql
	db2.connect()

}
