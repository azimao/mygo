package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	v1 := r.Group("/v1/topics")

	v1.GET("", func(c *gin.Context) {
		c.DefaultQuery("username", "azimao")
		if c.Query("username") == "" {
			c.String(200, "get something")
		} else {
			c.String(200, "get param:%s", c.Query("username"))
		}
		//c.JSON(200, c.Param("topic_id"))
	})
	v1.GET("/:topic_id", func(c *gin.Context) {
		c.String(200, "%s", c.Param("topic_id"))
	})

	r.Run()
	// 开启服务，默认端口是8080
}
