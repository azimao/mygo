package src

import "time"

type Topic struct { //单个Topic实体
	TopicID         int       `json:"id" gorm:"PRIMARY_KEY"`
	TopicTitle      string    `json:"title" binding:"min=4,max=20"`
	TopicShortTitle string    `json:"stitle" binding:"required,nefield=TopicTitle"`
	UserIP          string    `json:"ip" binding:"ipv4"`
	TopicScore      int       `json:"score" binding:"omitempty,gt=5"`
	TopicUrl        string    `json:"url" binding:"omitempty,topicurl"`
	TopicDate       time.Time `json:"date" binding:"required"`
}
type Topics struct {
	TopicList     []Topic `json:"topics" binding:"gt=0,lt=3,topics,dive"`
	TopicListSize int     `json:"size"`
}

type TopicClass struct {
	ClassId     int `gorm:"PRIMARY_KEY"`
	ClassName   string
	ClassRemark string
	ClassType   int `gorm:"Column:classtype"`
}

func CreateTopic(id int, title string) Topic {
	return Topic{TopicID: id, TopicTitle: title}
}

type TopicQuery struct {
	UserName string `json:"username" form:"username"`
	Page     int    `json:"page" form:"page" binding:"required"`
	PageSize int    `json:"pagesize" form:"pagesize"`
}
