package po

import (
	"fmt"
	"fromework/pkg/mysql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

//数据库中表的定义
type ProcessDef struct {
	ID         int    `gorm:"AUTO_INCREMENT"`
	Name       string `gorm:"size:255"` // string默认长度为255, 使用这种tag重设。
	Type       int    `gorm:"not null;unique"`
	Step       int    `gorm:"not null;unique"`
	CreateUser string `gorm:"size:255"`
	CreateTime time.Time
}

//数据库的查询
func SelectDBPD() []ProcessDef {
	//查出来的数据存放在data数组中
	var data []ProcessDef
	res := mysql.GetDB().Select([]string{"id", "name", "type"}).Find(&data)
	//判断查询的数据是否为空
	if res == nil {
		fmt.Println("exec failed, ", res)
		return nil
	}
	return data
}
