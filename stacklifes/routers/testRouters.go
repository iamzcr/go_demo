package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestRoutersInit(r *gin.Engine) {
	TestRouter := r.Group("/test")
	{
		TestRouter.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"success": true,
				"msg":     "frontend",
			})
		})
		TestRouter.GET("/articles", func(c *gin.Context) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"success": true,
				"msg":     "articles",
			})
		})
		TestRouter.GET("/articles/detail", func(c *gin.Context) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"success": true,
				"msg":     "/articles/detail",
			})
		})
	}
}
