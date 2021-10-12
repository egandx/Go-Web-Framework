package src

type Topics struct { //单个topic实体
	TopicID         int       `json:"id" gorm:"primaryKey"`
	TopicTitle      string    `json:"title" binding:"required,min=4,max=50"`
	TopicShortTitle string    `json:"stitle" binding:"nefield=TopicTitle"`
	UserIP          string    `json:"ip" binding:"ipv4"`
	TopicScore      int       `json:"score" binding:"omitempty,gt=5"`
	TopicUrl        string    `json:"url" binding:"omitempty"`
	TopicDate       string `json:"date"`
	UserName        string    `json:"username" gorm:"column:username" binding:"required"`
}

type TopicArray struct { //多个实体
	TopicList     []Topics `json:"topics" binding:"gt=0,lte=5,topics,dive"`
	TopicListSize int      `json:"size"`
}

type TopicQuery struct {
	Username string `form:"username" json:"username" binding:"required"`
	Page     int    `form:"page" json:"page"`
	Pagesize int    `form:"pagesize" json:"pagesize"`
}

type TopicClass struct {
	ClassId     int `gorm:"primaryKey"`
	ClassName   string
	ClassRemark string
	ClassType   string `gorm:"column:classtype"`
}

//func CreateTopic(id int,title string) Topics {
//	return Topics{id,title}
//}
