package model

type Category struct { //默认操作categorys表
	Id   int
	Name string
}

// 修改操作数据库表的名字
func (Category) TableName() string {
	return "category"
}
