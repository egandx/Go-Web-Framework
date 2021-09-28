package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

//
//var db = make(map[string]string)
//
//func setupRouter() *gin.Engine {
//	// Disable Console Color
//	// gin.DisableConsoleColor()
//	r := gin.Default()
//
//	// Ping test
//	r.GET("/ping", func(c *gin.Context) {
//		c.String(http.StatusOK, "pong")
//	})
//
//	// Get user value
//	r.GET("/user/:name", func(c *gin.Context) {
//		user := c.Params.ByName("name")
//		value, ok := db[user]
//		if ok {
//			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
//		} else {
//			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
//		}
//	})
//
//	// Authorized group (uses gin.BasicAuth() middleware)
//	// Same than:
//	// authorized := r.Group("/")
//	// authorized.Use(gin.BasicAuth(gin.Credentials{
//	//	  "foo":  "bar",
//	//	  "manu": "123",
//	//}))
//	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
//		"foo":  "bar", // user:foo password:bar
//		"manu": "123", // user:manu password:123
//	}))
//
//	/* example curl for /admin with basicauth header
//	   Zm9vOmJhcg== is base64("foo:bar")
//
//		curl -X POST \
//	  	http://localhost:8080/admin \
//	  	-H 'authorization: Basic Zm9vOmJhcg==' \
//	  	-H 'content-type: application/json' \
//	  	-d '{"value":"bar"}'
//	*/
//	authorized.POST("admin", func(c *gin.Context) {
//		user := c.MustGet(gin.AuthUserKey).(string)
//
//		// Parse JSON
//		var json struct {
//			Value string `json:"value" binding:"required"`
//		}
//
//		if c.Bind(&json) == nil {
//			db[user] = json.Value
//			c.JSON(http.StatusOK, gin.H{"status": "ok"})
//		}
//	})
//
//	return r
//}
//
//func main() {
//	r := setupRouter()
//	// Listen and Server in 0.0.0.0:8080
//	r.Run(":8080")
//}

//func formatAsDate(t time.Time)string {
//	year, month, day := t.Date()
//	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
//}

//var html = template.Must(template.New("https").Parse(`
//<html>
//<head>
//  <title>Https Test</title>
//  <script src="/assets/app.js"></script>
//</head>
//<body>
//  <h1 style="color:red;">Welcome, Ginner!</h1>
//</body>
//</html>
//`))

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	//r := gin.Default()
	//r.LoadHTMLGlob("html/*")
	//
	//r.GET("/json", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//
	//	data := map[string]interface{}{
	//		"lang":"Golang",
	//		"tag":"<br>",
	//		"framework":"Gin框架",
	//	}
	//
	//	c.AsciiJSON(http.StatusOK,data)
	//
	//	c.Writer.WriteString("\n hello world!")
	//})
	//
	//r.GET("/html", func(c *gin.Context) {
	//	c.HTML(http.StatusOK,"main.html",gin.H{
	//		"title":"hello html",
	//		"message":"yyds",
	//	})
	//})
	//
	//g1 := r.Group("g1")
	//g1.GET("/say", func(c *gin.Context) {
	//	c.Writer.WriteString("say hello g1")
	//})
	//
	//g2:=r.Group("g2")
	//g2.GET("/say2", func(c *gin.Context) {
	//	c.Writer.WriteString("say hello g2")
	//})

	//r := gin.Default()
	//
	//r.Delims("{[[{","}]]}")
	//r.SetFuncMap(template.FuncMap{
	//	"formatAsDate":formatAsDate,
	//})
	//
	//r.LoadHTMLGlob("html/*")
	//r.GET("/html", func(c *gin.Context) {
	//	c.HTML(http.StatusOK,"main.html",gin.H{
	//		"title":"Gin",
	//		"now":time.Now(),
	//	})
	//})
	//
	//r.Run(":8080") // listen and serve on 0.0.0.0:8080

	//r := gin.Default()
	//r.Static("/assets", "./assets")
	//r.SetHTMLTemplate(html)
	//
	//r.GET("/", func(c *gin.Context) {
	//	if pusher := c.Writer.Pusher(); pusher != nil {
	//		// 使用 pusher.Push() 做服务器推送
	//		if err := pusher.Push("/assets/app.js", nil); err != nil {
	//			log.Printf("Failed to push: %v", err)
	//		}
	//	}
	//	c.HTML(200, "https", gin.H{
	//		"status": "success",
	//	})
	//})
	//
	//// 监听并在 https://127.0.0.1:8080 上启动服务
	//r.Run(":8080")

	//r := gin.Default()
	//
	//r.GET("/JSONP", func(c *gin.Context) {
	//	data := map[string]interface{}{
	//		"foo": "bar",
	//	}
	//
	//	// /JSONP?callback=x
	//	// 将输出：x({\"foo\":\"bar\"})
	//	c.JSONP(http.StatusOK, data)
	//})

	//r:=gin.Default()
	//
	//r.POST("/login", func(c *gin.Context) {
	//	var form LoginForm
	//	if c.ShouldBind(&form) ==nil {
	//		if form.User == "user" && form.Password =="password" {
	//			c.JSON(http.StatusOK,"login success")
	//		} else {
	//			c.JSON(401,"login failed")
	//		}
	//	}
	//})

	//router := gin.Default()
	//
	//router.POST("/post", func(c *gin.Context) {
	//	mess := c.PostForm("message")
	//	nick := c.DefaultPostForm("nick", "ano")
	//
	//	c.JSON(200, gin.H{
	//		"status":  "post",
	//		"message": mess,
	//		"nick":    nick,
	//	})
	//})

	router:= gin.Default()

	//router.GET("/json", func(c *gin.Context) {
	//	c.JSON(http.StatusOK,gin.H{
	//		"html":"<h1>Hello</h1>",
	//	})
	//})

	router.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(http.StatusOK,gin.H{
			"html":"<h1>Hello</h1>",
		})
	})

	//router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s \n", id, page, name, message)
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	router.Run(":8080")
}
