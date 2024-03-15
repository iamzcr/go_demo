package main

import "fmt"

// 结构体
// 自定义类型
type myInt int

// 自定义方法类型
type myFunc func(int, int) int

// 1.9加入类型别名
type myFloat = float64

// 定义结构体，首字母大写表示公有，小写表示私有
type Person struct {
	name string
	age  int
	sex  int
}

func main() {
	var num myInt = 10
	fmt.Printf("num值是%v,类型是%T,地址是%p\n", num, num, &num)

	var testFloat myFloat = 10.0
	fmt.Printf("testFloat值是%v,类型是%T,地址是%p\n", testFloat, testFloat, &testFloat)

	var person Person //实例化结构体
	person.name = "test"
	person.age = 12
	person.sex = 1
	fmt.Printf("person值是%v,全部信息是%#v,类型是%T,地址是%p\n", person, person, person, &person)

	//键值对实例化
	var person3 = Person{
		name: "test3",
		age:  88,
		sex:  1,
	}
	fmt.Printf("person3值是%v,全部信息是%#v,类型是%T,地址是%p\n", person3, person3, person3, &person3)

	//new实例化，但是返回的时指针类型，实际上该方式出来的实例，理论上需要通过结构体指针访问成员变量的，(*person1).name，
	// 而实际上person1.name这样访问，底层也是(*person1).name,而不需要显式声明
	var person1 = new(Person)
	person1.name = "test1"
	person1.age = 15
	(*person1).sex = 2

	fmt.Printf("person1值是%v,全部信息是%#v,类型是%T,地址是%p\n", person1, person1, person1, &person1)

	//也可以使用引用方式实例化
	var person2 = &Person{}
	person2.name = "test2"
	person2.age = 15
	(*person2).sex = 2
	fmt.Printf("person2值是%v,全部信息是%#v,类型是%T,地址是%p\n", person2, person2, person2, &person2)

	//引用方式实例化
	var person4 = &Person{
		name: "test3",
		age:  88,
		sex:  1,
	}
	fmt.Printf("person4值是%v,全部信息是%#v,类型是%T,地址是%p\n", person4, person4, person4, &person4)

	//实例化的时候给部分属性赋值
	var person5 = &Person{
		name: "test3",
	}
	fmt.Printf("person5值是%v,全部信息是%#v,类型是%T,地址是%p\n", person5, person5, person5, &person5)

	//不指定实例化，但是顺序要和定义的类型里面的属性顺序一样
	var person6 = &Person{
		"test3",
		11,
		1,
	}
	fmt.Printf("person6值是%v,全部信息是%#v,类型是%T,地址是%p\n", person6, person6, person6, &person6)

}
