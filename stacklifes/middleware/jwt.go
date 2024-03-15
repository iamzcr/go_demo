package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func InitJwt(c *gin.Context) {
	c.Set("title", "InitJwt中间件设置一个值")
	fmt.Println("登录验证中间件.......")
	// 创建上下文的副本
	copiedContext := c.Copy()

	// 当在中间件或者handler启动了新的goroutine时，不能使用原始的上下文c *gin.Context，必须使用其自读副本(c.Copy())
	//	//这是因为 Gin 的上下文 c *gin.Context 是并发安全的，它在每个请求中都会被重用。如果直接在新的 goroutine 中使用原始的上下文 c *gin.Context，
	//	//可能会导致并发访问和修改上下文的状态，从而引发竞态条件和数据不一致的问题。
	go func() {
		// 在新的 goroutine 中使用副本进行处理
		// 可以安全地访问副本的属性和方法
		time.Sleep(3 * time.Second)
		fmt.Println("InitJwt中间件的协程" + copiedContext.Request.URL.Path)

	}()

}
