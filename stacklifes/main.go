package main

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"log"
	"net/http"
	"os"
	"path"
	"stacklifes/db"
	"stacklifes/model"
	"stacklifes/routers"
	"strconv"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

type MovieData struct {
	Title  string  `json:"title" form:"title" xml:"title"`
	Rating float64 `json:"rating" form:"rating"  xml:"rating"`
	Year   int     `json:"year" form:"year"  xml:"year"`
}

func UnixToTime(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format("2006-01-02 15:04:05")
}

// 中间件
func InitMiddleware(c *gin.Context) {
	start := time.Now().UnixNano()
	fmt.Println("抽离中间件")
	//调用该请求的剩余处理程序，统计某个接口的程序执行时间
	c.Next()
	end := time.Now().UnixNano()
	fmt.Println("抽离中间件-1")
	fmt.Println("程序执行时间", end-start)
}
func InitMiddlewareTemp(c *gin.Context) {
	fmt.Println("抽离中间件InitMiddlewareTemp")
	//停止执行后续程序
	c.Abort()
	fmt.Println("抽离中间件InitMiddlewareTemp-1")
}
func main() {
	//New() 函数：
	//New() 函数用于创建一个全新的空白 Gin 引擎实例。
	//使用 New() 创建的引擎实例没有任何默认的中间件或配置，您需要手动添加所需的中间件和配置。
	//这使您有更大的自由度来根据自己的需求自定义和配置 Gin 引擎。
	//Default() 函数：
	//Default() 函数用于创建一个带有默认配置和中间件的 Gin 引擎实例。
	//使用 Default() 创建的引擎实例已经预先配置了一些常用的中间件和默认设置，例如 Logger 中间件、Recovery 中间件等。
	//这使您可以快速启动一个具备常用功能的 Gin 应用程序，而无需手动添加和配置这些中间件。

	//创建一个默认路由引擎,默认加载了Logger(), Recovery()中间间，看源码
	r := gin.Default()

	//自定义模板函数，放在加载模板前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
	})
	//加载模板文件
	//r.LoadHTMLGlob("templates/*")
	r.LoadHTMLGlob("templates/**/*")

	//配置静态资源，第一个参数是路由，第二个是资源目录
	r.Static("/static", "./static")
	//注册全局中间件 ，当在中间件或者handler启动了新的goroutine时，不能使用原始的上下文c *gin.Context，必须使用其自读副本(c.Copy())
	//这是因为 Gin 的上下文 c *gin.Context 是并发安全的，它在每个请求中都会被重用。如果直接在新的 goroutine 中使用原始的上下文 c *gin.Context，
	//可能会导致并发访问和修改上下文的状态，从而引发竞态条件和数据不一致的问题。
	//r.Use(InitMiddleware)
	//设置session中间件
	//创建一个基于cookie的存储引擎
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	//渲染模板
	r.GET("/",
		//在路由加载前可以定义中间件，因为GET,POST这些方法可以接受多个函数
		func(context *gin.Context) {
			fmt.Println("中间件")
		},
		//可以把中间件抽离出来
		InitMiddleware,
		func(c *gin.Context) {
			movieData := &MovieData{
				Title:  "The Shawshank Redemption",
				Rating: 9.3,
				Year:   1994,
			}
			c.HTML(http.StatusOK, "frontend/index", gin.H{
				"message": "index",
				"time":    time.Now().Unix(),
				"data":    movieData,
			})
		})

	//路由传值 get
	r.GET("/query", func(c *gin.Context) {

		name := c.Query("name")
		pwd := c.DefaultQuery("pwd", "123456")
		c.HTML(http.StatusOK, "frontend/query", gin.H{
			"name": name,
			"pwd":  pwd,
		})
	})
	//路由传值 post
	r.GET("/post", func(c *gin.Context) {

		c.HTML(http.StatusOK, "frontend/post", gin.H{})
	})
	r.POST("/postAction", func(c *gin.Context) {

		username := c.PostForm("username")
		pwd := c.DefaultPostForm("pwd", "123456")
		c.JSON(http.StatusOK, gin.H{ //这里就是一个map[string]interface{}
			"name": username,
			"pwd":  pwd,
		})
	})

	//路由传值 post 绑定结构体
	r.GET("/getMovie", func(c *gin.Context) {

		movieD := &MovieData{}
		fmt.Println(movieD)
		if err := c.ShouldBind(movieD); err == nil {
			c.JSON(http.StatusOK, movieD)
		} else {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		}
	})

	r.GET("/movie", func(c *gin.Context) {
		c.HTML(http.StatusOK, "frontend/movie", gin.H{})
	})

	r.POST("/postMovie", func(c *gin.Context) {

		var movieData MovieData

		if err := c.ShouldBind(&movieData); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "movieData created successfully", "data": movieData})
	})
	//获取xml
	r.POST("/postXmlMovie", func(c *gin.Context) {

		var movieData MovieData

		xmlSlice, _ := c.GetRawData() //c.Request.body
		err := xml.Unmarshal(xmlSlice, &movieData)
		if err == nil {
			c.JSON(http.StatusOK, movieData)
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		}

	})
	//渲染模板
	r.GET("/test", func(c *gin.Context) {
		movieData := &MovieData{
			Title:  "The Shawshank Redemption",
			Rating: 9.3,
			Year:   1994,
		}
		c.HTML(http.StatusOK, "test", gin.H{
			"message": "index",
			"data":    movieData,
		})
	})

	//配置路由：返回json
	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{ //这里就是一个map[string]interface{}
			"message": "989898",
		})
	})
	r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"message": "989898",
		})
	})
	//配置路由：返回json
	r.GET("/jsonTest", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"success": true,
			"msg":     "success",
		})
	})

	//http://localhost:8080/jsonp?callback=func
	r.GET("/jsonp", func(c *gin.Context) {
		c.JSONP(http.StatusOK, map[string]interface{}{
			"success": true,
			"msg":     "success",
		})
	})

	r.GET("/string", func(c *gin.Context) {
		c.String(http.StatusOK, "%v", "string")
	})
	r.POST("/add", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "989898",
		})
	})
	r.PUT("/edit", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "989898",
		})
	})

	r.DELETE("/del", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "989898",
		})
	})

	//操作gorm
	r.GET("/category", func(c *gin.Context) {
		var categoryList []model.Category
		db.MysqlDB.Find(&categoryList)
		c.JSON(http.StatusOK, gin.H{
			"message": "989898",
			"result":  categoryList,
		})
	})
	//sharecookie，同一个站点不同域名共享cookie
	r.GET("/sharecookie", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_share_cookie")

		if err != nil {
			cookie = "NotSet"
			//斜杠表示当前所有页面都能访问,.iamzcr.com表示该主域名下面的子域名共享
			c.SetCookie("gin_share_cookie", "test", 3600, "/", ".iamzcr.com", false, false)
		}

		fmt.Printf("Cookie value: %s \n", cookie)

		cookie, err = c.Cookie("gin_share_cookie")

		c.JSON(http.StatusOK, gin.H{
			"message": "989898",
			"Cookie":  cookie,
		})
	})
	//session
	r.GET("/session", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		fmt.Println(v)
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		err := session.Save()
		if err != nil {
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "989898",
			"session": v,
		})
	})
	//cookie
	r.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie")

		if err != nil {
			cookie = "NotSet"
			//斜杠表示当前所有页面都能访问
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, false)
		}

		fmt.Printf("Cookie value: %s \n", cookie)

		cookie, err = c.Cookie("gin_cookie")

		c.JSON(http.StatusOK, gin.H{
			"message": "989898",
			"Cookie":  cookie,
		})
	})
	//动态路由
	//http://localhost:8080/testRoute/123
	r.GET("/testRoute/:cid", func(c *gin.Context) {
		cid := c.Param("cid")
		c.JSON(http.StatusOK, gin.H{ //这里就是一个map[string]interface{}
			"message": "989898",
			"cid":     cid,
		})
	})

	//文件上传
	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "frontend/upload", gin.H{})
	})
	r.POST("/uploadAction", func(c *gin.Context) {
		// 单文件
		userName := c.PostForm("username")
		log.Println(userName)

		file, _ := c.FormFile("myfile")
		log.Println(file.Filename)
		ext := path.Ext(file.Filename)
		log.Println(ext)

		allowExt := map[string]bool{
			".jpg": true,
			".JPG": true,
		}
		if _, ok := allowExt[ext]; !ok {
			return
		}
		dir := "./static/upload/" + "2023"
		err := os.MkdirAll(dir, 0666)
		if err != nil {
			return
		}

		fileName := strconv.FormatInt(time.Now().Unix(), 10) + ext
		dst := path.Join(dir, fileName)

		// 上传文件至指定的完整文件路径
		err = c.SaveUploadedFile(file, dst)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, gin.H{ //这里就是一个map[string]interface{}
			"message": "success",
		})
	})

	r.GET("/mutiUpload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "frontend/muti_upload", gin.H{})
	})

	r.POST("/mutiUploadAction", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["files"]
		for _, file := range files {
			log.Println(file.Filename)
			dst := path.Join("./static/upload", file.Filename)
			// 上传文件至指定目录
			err := c.SaveUploadedFile(file, dst)
			if err != nil {
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{ //这里就是一个map[string]interface{}
			"message": "success",
		})
	})

	//路由分组
	//路由分组抽离成文件，注册路由
	routers.TestRoutersInit(r)
	routers.ApiRoutersInit(r)
	adminRouter := r.Group("/admin")
	{
		adminRouter.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"success": true,
				"msg":     "admin",
			})
		})
		adminRouter.GET("/user", func(c *gin.Context) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"success": true,
				"msg":     "success",
			})
		})
		adminRouter.GET("/user/add", func(c *gin.Context) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"success": true,
				"msg":     "success",
			})
		})
	}

	//启动web服务
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
