package main

import (
	. "gin/DemoProjects/Topic/src"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("topicurl",TopicUrl)
		v.RegisterValidation("topics",TopicsValidate)
	}

	v1 := router.Group("/v1/topics") //单条帖子处理
	{
		v1.GET("", GetTopicList)
		v1.GET("/:topic_id", GetTopicDetail)

		//v1.Use(gin.BasicAuth(gin.Accounts{
		//	"admin":"123",
		//}))

		v1.Use(MustLogin())
		{
			v1.POST("", AddTopic)
			v1.DELETE("/:topic_id", DelTopic)
		}
	}

	v2 := router.Group("/v2/mtopics") //多条帖子处理
	{
		//v2.Use(gin.BasicAuth(gin.Accounts{
		//	"admin":"123",
		//}))

		v2.Use(MustLogin())
		{
			v2.POST("", AddTopics)
			v2.DELETE("/:topic_id", DelTopic)
		}
	}

	router.Run(":8080")
}
