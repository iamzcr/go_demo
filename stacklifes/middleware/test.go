package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func TestInit(c *gin.Context) {
	title, _ := c.Get("title") //返回的是空接口
	s, ok := title.(string)
	if ok {
		fmt.Println("在TestInit中间件方法里面获取：", s)

	} else {
		fmt.Println("在TestInit中间件方法里面获取不到")

	}
}
