package main

import (
	. "gin/DemoProjects/Topic/src"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1:= router.Group("/v1/topics")
	{
		v1.GET("", GetTopicList)
		v1.GET("/:topic_id",GetTopicDetail)

		//v1.Use(gin.BasicAuth(gin.Accounts{
		//	"admin":"123",
		//}))

		v1.Use(MustLogin())
		{
			v1.POST("",AddTopic)
			v1.DELETE("/:topic_id",DelTopic)
		}



	}


	router.Run(":8080")
}
