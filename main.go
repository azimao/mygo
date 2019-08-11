package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/topic/:topic_id", func(c *gin.Context) {
		c.JSON(200, c.Param("topic_id"))
	})
	r.Run() // 开启服务，默认端口是8080
}
