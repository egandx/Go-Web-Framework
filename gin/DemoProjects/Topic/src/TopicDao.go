package src

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// MustLogin 中间件
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

// GetTopicDetail get topic detail by topic_id
func GetTopicDetail(c *gin.Context) {

	tid := c.Param("topic_id")
	topics := Topics{}

	db.Find(&topics, tid) //data from DB

	c.Set("dbResultTopicDetail", topics) //send to Decorator,this is key
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

// GetTopicList get all of topic
func GetTopicList(c *gin.Context) {

	var topicslist []Topics
	rows, _ := db.Find(&Topics{}).Rows()
	for rows.Next() {
		db.Find(&topicslist)
	}

	c.Set("dbResultTopicList", topicslist)

	//tl := TopicArray{
	//	topicslist,
	//	len(topicslist),
	//}
}

// QueryTopics query topic for search by username
func QueryTopics(c *gin.Context) {

	username := c.Query("username")
	query := TopicQuery{}
	query.Pagesize = 10

	err := c.BindQuery(&query)
	if err != nil {
		c.String(400, "参数错误：%s", err.Error())
	} else {
		t := Topics{}
		var topiclist []Topics
		size := 0
		rows, _ := db.Raw("select * from topics where username=?", username).Rows()
		for rows.Next() {
			rows.Scan(&t.TopicID, &t.TopicTitle, &t.TopicShortTitle, &t.UserIP, &t.TopicUrl, &t.TopicScore, &t.TopicDate, &t.UserName)
			topiclist = append(topiclist, t)
			size++
		}
		TotalPage := size/query.Pagesize + 1

		var p map[string]interface{}
		p = make(map[string]interface{})

		p["Topic_list"] = topiclist
		p["Size"] = size
		p["Total_page"] = TotalPage
		p["Page_size"] = query.Pagesize

		//c.JSON(200,p)
		c.Set("dbResultUsername", p)
	}

}

// AddTopic add one topic,need login
func AddTopic(c *gin.Context) { //单条帖子新增

	topic := Topics{}

	err := c.BindJSON(&topic)
	if err != nil {
		msg := fmt.Sprintf("参数错误：%s", err.Error())
		c.JSON(400, msg)

	} else {
		t := Topics{
			TopicTitle:      topic.TopicTitle,
			TopicShortTitle: topic.TopicShortTitle,
			UserIP:          topic.UserIP,
			TopicScore:      topic.TopicScore,
			TopicUrl:        topic.TopicUrl,
			TopicDate:       time.Now().Format("2006-01-02 15:04:05"),
			UserName:        topic.UserName,
		}

		if err := db.Create(&t).Error; err != nil {
			fmt.Println("插入失败:", err)
			return
		}else{
			c.JSON(200, "OK")

		}



	}

}

// AddTopics add more topics,need login
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

// DelTopic del topic, need login
func DelTopic(c *gin.Context) {
	c.String(200, "删除topicID为%s帖子", c.Param("topic_id"))
}
