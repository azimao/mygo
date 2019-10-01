package src

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

var DBS *gorm.DB
var err error

func init() {
	DBS, err = gorm.Open("mysql", "azimao:320199@/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}

	DBS.DB().SetMaxIdleConns(10)
	DBS.DB().SetMaxOpenConns(100)
	DBS.DB().SetConnMaxLifetime(time.Hour)

	DBS.LogMode(true)
}
