package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/go-playground/validator.v8"
	. "mygo/src"
)

//全局变量
func main() {

	//db, _ := gorm.Open("mysql", "root:@/gin?charset=utf8&parseTime=True&loc=Local")

	//rows, _ := db.Raw("select topic_id,topic_title from topics").Rows()
	//for rows.Next() {
	//	var t_id int
	//	var t_title string
	//	rows.Scan(&t_id, &t_title)
	//	fmt.Println(t_id, t_title)
	//}

	//db.LogMode(true)
	//topics := Topic{
	//	TopicTitle:      "TopicTitle",
	//	TopicShortTitle: "TopShTi",
	//	UserIP:          "127.0.0.1",
	//	TopicScore:      0,
	//	TopicUrl:        "azimao.com",
	//	TopicDate:       time.Now()}
	//fmt.Println(db.Create(&topics).RowsAffected)
	//fmt.Println(topics.TopicID)
	//defer db.Close()

	//tc := TopicClass{}
	//db.First(&tc, 2)
	//fmt.Println(tc)
	//
	//var tcs []TopicClass
	//db.Where("class_name=?", "t").Find(&tcs)
	//fmt.Println(tcs)
	//defer db.Close()
	//return;
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
