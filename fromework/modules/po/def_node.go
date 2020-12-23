package po

import (
	"fmt"
	"fromework/pkg/mysql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ProcessDefNode struct {
	ID int `gorm:"AUTO_INCREMENT"`
	//ProcessDef   ProcessDef
	ProcessDefName string
	ProcessDefId   int    `json:"process_def_id"`
	NodeName       string `gorm:"size:255"`
	NodeStep       int
	NodeNext       int
	ApproveUser    string
	ApproveGroup   int
	ApproveAdmin   string
}

func SelectDBPDN(id int) []ProcessDefNode {
	//mysql.GetDB().LogMode(true)  打印数据库的语句
	var data []ProcessDefNode
	err := mysql.GetDB().Select([]string{"node_step", "approve_user", "approve_admin"}).Where("process_def_id=?", id).Find(&data).Error
	if err != nil {
		fmt.Println("exec failed, ", err)
		return nil
	}
	return data
}
