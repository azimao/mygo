package src

import "github.com/gin-gonic/gin"

func GetTopicDetail(c *gin.Context) {
	c.String(200, "%s", c.Param("topic_id"))
}

func NewTopic(c *gin.Context) {
	c.String(200, "Add topic from")
}

func DelTopic(c *gin.Context) {
	c.String(200, "Delete a topic")
}
