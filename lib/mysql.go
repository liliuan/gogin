package lib

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

type OptionDBer interface {
	Find(out interface{}, where interface{})
}

func initMysql() {
	db, err := gorm.Open("mysql", "root:123456@/mydb?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	DB = db
}

func init() {
	initMysql()
}
