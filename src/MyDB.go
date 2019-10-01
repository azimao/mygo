package src

import (
	"github.com/jinzhu/gorm"
	"time"
)

var DBS *gorm.DB
var err error

func InitDB() {
	DBS, err = gorm.Open("mysql", "azimao:32019i9@/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		ShutDownServer(err)
	}

	DBS.DB().SetMaxIdleConns(10)
	DBS.DB().SetMaxOpenConns(100)
	DBS.DB().SetConnMaxLifetime(time.Hour)

	DBS.LogMode(true)
}
