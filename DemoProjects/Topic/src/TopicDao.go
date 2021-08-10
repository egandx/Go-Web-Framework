package src

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//中间件MustLogin
func MustLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, ok := c.GetQuery("token"); !ok {
			c.String(http.StatusUnauthorized, "缺少token参数")
			c.Abort()
		} else {
			c.Next()
		}
	}
}

func GetTopicDetail(c *gin.Context) {
	//c.String(200, "获取topicID为%s的帖子内容", c.Param("topic_id"))
	c.JSON(200, CreateTopic(111, "帖子标题"))
}

func GetTopicList(c *gin.Context) {
	//if c.Query("username") == "" {
	//	c.String(200, "获取所有帖子列表")
	//} else {
	//	c.String(200, "获取用户名为%s的帖子列表", c.Query("username"))
	//}

	query := TopicQuery{}

	err := c.BindQuery(&query)
	if err != nil {
		c.String(400, "参数错误：%s", err.Error())
	} else {
		c.JSON(200, query)
	}
}

// 需要登陆
func AddTopic(c *gin.Context) {
	//c.String(200, "新增帖子")

	topic := Topic{}

	err := c.BindJSON(&topic)
	if err != nil {
		c.String(400, "参数错误：%s", err.Error())
	} else {
		c.JSON(200, topic)
	}

}

// 需要登陆
func DelTopic(c *gin.Context) {
	c.String(200, "删除topicID为%s帖子", c.Param("topic_id"))
}
