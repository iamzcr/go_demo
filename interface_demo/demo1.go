package main

import "fmt"

// Database 接口是一种类型，一种抽象的类型呢，是一组函数的集合，接口不能包含任何变量
// 接口中所有方法都没有方法体，定义一个对象的行为规范，不实现规范，体现了程序实现的多态和高内聚，低耦合的思想
// 接口也是一种数据类型，不需要显式实现，只需要一个变量，还有接口类型的所有方法，那么这个变量就实现了这个接口
/*
	type 接口名 interface {
		方法名(参数)返回值
	}
*/
//接口的方法，要通过结构体或者自定义类型实现这个接口
type Database interface {
	connect()
	query()
}
type Mysql struct {
	Name string
}

func (m Mysql) connect() {
	fmt.Println(m.Name, "mysql结构体实现了Database接口的方法connect")
}
func (m Mysql) query() {
	fmt.Println(m.Name, "mysql结构体实现了Database接口的方法query")
}

// 结构体里面可以有自己的方法
func (m Mysql) myfunc() {
	fmt.Println(m.Name, "mysql结构体调用myfunc")
}

type Pgsql struct {
	Name string
}

func (p Pgsql) connect() {
	fmt.Println(p.Name, "Pgsql结构体实现了Database接口的方法connect")
}
func (p Pgsql) query() {
	fmt.Println(p.Name, "Pgsql结构体实现了Database接口的方法query")
}
func main() {

	var mysql = Mysql{
		Name: "mysql",
	}
	//没实现接口之前，表示结构体方法
	mysql.connect()

	//实现接口调用方法
	var db Database
	db = mysql
	db.connect()

	//实现接口之前，但是也可以独立调用自己的方法
	mysql.myfunc()

	var pgDb Database = Pgsql{}
	pgDb.connect()

}
