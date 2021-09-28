package src

import "time"

type Topics struct { //单个topic实体
	TopicID int `json:"id" binding:"required" gorm:"primaryKey"`
	TopicTitle string `json:"title" binding:"min=4,max=20"`
	TopicShortTitle string `json:"stitle" binding:"required,nefield=TopicTitle"`
	UserIP string `json:"ip" binding:"ipv4"`
	TopicScore int `json:"score" binding:"omitempty,gt=5"`
	TopicUrl string `json:"url" binding:"omitempty,topicurl"`
	TopicDate time.Time `json:"date" binding:"required"`
}

type TopicArray struct { //多个实体
	TopicList []Topics `json:"topics" binding:"gt=0,lte=5,topics,dive"`
	TopicListSize int  `json:"size"`
}


type TopicQuery struct {
	Username string `form:"username" json:"username" binding:"required"`
	Page int `form:"page" json:"page" binding:"required"`
	Pagesize int `form:"pagesize" json:"pagesize" binding:"required"`
}

type TopicClass struct {
	ClassId int `gorm:"primaryKey"`
	ClassName string
	ClassRemark string
	ClassType string `gorm:"column:classtype"`
}

//func CreateTopic(id int,title string) Topics {
//	return Topics{id,title}
//}