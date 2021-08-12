package main

import (
	"fmt"
	. "gin/DemoProjects/Topic/src"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:12345678@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	tc:=TopicClass{}
	//rows, _ := db.Raw("select topic_id,topic_title from topics").Rows()
	//for rows.Next(){
	//	var id int
	//	var title string
	//	rows.Scan(&id,&title)
	//	fmt.Println(id,title)
	//}
	db.First(&tc)

	fmt.Println(tc)



	//router := gin.Default()
	//
	//if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	//	v.RegisterValidation("topicurl",TopicUrl)
	//	v.RegisterValidation("topics",TopicsValidate)
	//}
	//
	//v1 := router.Group("/v1/topics") //单条帖子处理
	//{
	//	v1.GET("", GetTopicList)
	//	v1.GET("/:topic_id", GetTopicDetail)
	//
	//	//v1.Use(gin.BasicAuth(gin.Accounts{
	//	//	"admin":"123",
	//	//}))
	//
	//	v1.Use(MustLogin())
	//	{
	//		v1.POST("", AddTopic)
	//		v1.DELETE("/:topic_id", DelTopic)
	//	}
	//}
	//
	//v2 := router.Group("/v2/mtopics") //多条帖子处理
	//{
	//	//v2.Use(gin.BasicAuth(gin.Accounts{
	//	//	"admin":"123",
	//	//}))
	//
	//	v2.Use(MustLogin())
	//	{
	//		v2.POST("", AddTopics)
	//		v2.DELETE("/:topic_id", DelTopic)
	//	}
	//}
	//
	//router.Run(":8080")
}
