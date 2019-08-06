package lib

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func InitMysql() {
	Db, err := gorm.Open("mysql", "root:123456@/mydb?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	fmt.Println("DB", Db)
}
