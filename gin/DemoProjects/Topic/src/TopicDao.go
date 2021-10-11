package src

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 中间件MustLogin
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
	//c.JSON(200, CreateTopic(111, "帖子标题"))

	tid := c.Param("topic_id")
	topics := Topics{}

	db.Find(&topics, tid)		//data from DB

	c.Set("dbResult",topics) //send to Decorator,this is key



	//conn := RedisDefaultPool.Get()
	//defer conn.Close()
	//redisKey := "topic_" + tid
	//
	//ret, err := redis.Bytes(conn.Do("get", redisKey))
	//
	//if err != nil { //缓存里没有，去DB
	//	db.Find(&topics, tid)
	//	retData, _ := json.Marshal(topics)
	//
	//	if topics.TopicID == 0{  //DB 没有匹配到
	//		conn.Do("setex",redisKey,20,retData)
	//	}else{ //DB正常数据，缓存50s
	//		conn.Do("setex",redisKey,50,retData)
	//	}
	//
	//	c.JSON(http.StatusOK,topics)
	//	log.Println("从数据库读取")
	//
	//} else { //缓存里有值
	//	json.Unmarshal(ret,&topics)
	//	c.JSON(http.StatusOK,topics)
	//	log.Println("从Redis读取")
	//}
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
func AddTopic(c *gin.Context) { //单条帖子新增
	//c.String(200, "新增帖子")

	topic := Topics{}

	err := c.BindJSON(&topic)
	if err != nil {
		c.String(400, "参数错误：%s", err.Error())
	} else {
		c.JSON(200, topic)
	}
}

func AddTopics(c *gin.Context) { //批量新增多条帖子
	//c.String(200, "新增帖子")

	topics := Topics{}

	err := c.BindJSON(&topics)
	if err != nil {
		c.String(400, "参数错误：%s", err.Error())
	} else {
		c.JSON(200, topics)
	}
}

func DelTopic(c *gin.Context) {
	c.String(200, "删除topicID为%s帖子", c.Param("topic_id"))
}
