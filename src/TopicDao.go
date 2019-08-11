package src

import (
	"github.com/gin-gonic/gin"
)

func GetTopicDetail(c *gin.Context) {
	//c.String(200, "%s", c.Param("topic_id"))
	c.JSON(200, CreateTopic(1, "TCazimao"))
}

func NewTopic(c *gin.Context) {
	c.String(200, "Add topic from")
}

func DelTopic(c *gin.Context) {
	c.String(200, "Delete a topic")
}

func GetTopicList(c *gin.Context) {
	query := TopicQuery{}
	err := c.BindQuery(&query)

	if err != nil {
		c.String(400, "Param error:%s", err.Error())
	} else {
		c.JSON(200, query)
	}
}
