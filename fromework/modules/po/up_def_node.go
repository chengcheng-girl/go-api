package po

import (
	"fmt"
	"fromework/modules/qo"
	"fromework/pkg/mysql"
)

func SelectDBPDN_U(udnbody qo.UDNBody) []ProcessDefNode {
	//mysql.GetDB().LogMode(true)打印数据库的语句sql
	var data []ProcessDefNode
	err := mysql.GetDB().Joins("right join process_def on process_def_node.process_def_id=process_def.id").
		Where("node_name=? and approve_user=? and approve_admin=?", udnbody.Items.NodeName, udnbody.Items.ApproveUser, udnbody.Items.ApproveAdmin).
		Find(&data).Error

	if err != nil {
		fmt.Println("exec failed, ", err)
		return nil
	}
	return data
}
