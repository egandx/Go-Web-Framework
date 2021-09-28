package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	. "gin/DemoProjects/Topic/src"
)

func main() {

	//dsn := "root:12345678@/gin?charset=utf8mb4&parseTime=True&loc=Local"
	//db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
	//	Logger: logger.Default.LogMode(logger.Info),
	//	NamingStrategy: schema.NamingStrategy{
	//		SingularTable: false,
	//	},
	//})
	//
	//topics := Topics{
	//	TopicTitle:"TopicTitle",
	//	TopicShortTitle:"TopicShortTitle",
	//	UserIP:"127.0.0.1",
	//	TopicScore:0,
	//	TopicUrl:"testurl",
	//	TopicDate:time.Now(),
	//}
	//
	//fmt.Println(db.Create(&topics).RowsAffected)
	//fmt.Println(topics.TopicID)
	//db, _ := gorm.Open("", dsn)

	//db.LogMode(true)

	//tc:=TopicClass{}
	//var tcs []TopicClass
	//rows, _ := db.Raw("select topic_id,topic_title from topics").Rows()
	//for rows.Next(){
	//	var id int
	//	var title string
	//	rows.Scan(&id,&title)
	//	fmt.Println(id,title)
	//}

	//db.Find(&tcs)
	////db.Where(&TopicClass{ClassName: "技术类"}).Find(&tcs)
	//
	//fmt.Println(tcs)


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
		v2.Use(gin.BasicAuth(gin.Accounts{
			"admin":"123",
		}))

		v2.Use(MustLogin())
		{
			v2.POST("", AddTopics)
			v2.DELETE("/:topic_id", DelTopic)
		}
	}

	router.Run(":8080")
}
