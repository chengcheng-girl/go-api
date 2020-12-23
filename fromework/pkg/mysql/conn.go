package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
)

var once sync.Once
var instance  *gorm.DB

func Init() {
	once.Do(func() {
		db, err := gorm.Open("mysql", "root:123456@tcp(172.27.139.87:3306)/workflow?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			fmt.Println("open mysql failed,", err)

		}
		instance = db
		db.SingularTable(true)
	})
}

func GetDB() *gorm.DB {
	return instance
}