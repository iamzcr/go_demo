package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"stacklifes/middleware"
)

func InitMiddlewareRouterGroup(c *gin.Context) {
	fmt.Println("路由分组中间件InitMiddlewareRouterGroup")
}
func InitMiddlewareRouterGroupTemp(c *gin.Context) {
	fmt.Println("路由分组中间件InitMiddlewareRouterGroup")
}
func ApiRoutersInit(r *gin.Engine) {
	ApiRouter := r.Group("/api", InitMiddlewareRouterGroup)                               //可以加载这个后面
	ApiRouter.Use(InitMiddlewareRouterGroupTemp, middleware.InitJwt, middleware.TestInit) //也可以单独注册,注册独立文件的middleware.InitJwt中间件
	{
		ApiRouter.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"success": true,
				"msg":     "api",
			})
		})
		ApiRouter.GET("/articles", func(c *gin.Context) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"success": true,
				"msg":     "articles",
			})
		})
		ApiRouter.GET("/articles/detail", func(c *gin.Context) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"success": true,
				"msg":     "/articles/detail",
			})
		})
	}
}
