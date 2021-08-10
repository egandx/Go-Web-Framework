package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//r := gin.Default()
	//
	//r.Use(gin.BasicAuth(gin.Accounts{
	//	"admin": "12345",
	//}))
	//
	//r.GET("/", func(c *gin.Context) {
	//	c.JSON(200, "首页")
	//})
	//
	//r.Run(":8080")

	r:= gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"title":"normal view",
		})
	})

	adminGroup:=r.Group("/admin")

	adminGroup.Use(gin.BasicAuth(gin.Accounts{
		"admin":"123",
	}))

	adminGroup.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK,"admin view")
	})

	r.Run(":8080")
}
