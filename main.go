package main

import (
	"github.com/gin-gonic/gin"
	. "mygo/src"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1/topics")
	{

		v1.GET("", GetTopicList)

		v1.GET("/:topic_id", GetTopicDetail)
		v1.Use(LoginAuth())
		{
			v1.POST("", NewTopic)
			v1.DELETE("/:topic_id", DelTopic)
		}
	}
	r.Run()
	// 开启服务，默认端口是8080
}
