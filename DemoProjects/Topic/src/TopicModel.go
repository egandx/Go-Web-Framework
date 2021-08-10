package src

type Topic struct {
	TopicID int `json:"id" binding:"required"`
	TopicTitle string `json:"title" binding:"required"`
}

type TopicQuery struct {
	Username string `form:"username" json:"username" binding:"required"`
	Page int `form:"page" json:"page" binding:"required"`
	Pagesize int `form:"pagesize" json:"pagesize" binding:"required"`
}

func CreateTopic(id int,title string) Topic {
	return Topic{id,title}
}