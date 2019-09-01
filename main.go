package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	. "mygo/src"
)

func main() {
	r := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("topicurl", TopicUrl)
		v.RegisterValidation("topics", TopicsValidate)
	}

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

	v2 := r.Group("/v2/mtopics")
	{

		v2.Use(LoginAuth())
		{
			v2.POST("", NewTopics)
		}
	}
	r.Run()
	// 开启服务，默认端口是8080
}
