package po

import (
	"fmt"
	"fromework/modules/qo"
	"fromework/pkg/mysql"
	"github.com/golang-sql/civil"
	"github.com/kataras/iris/v12/context"
)

type ProcessInstance struct {
	ID             int            `gorm:"AUTO_INCREMENT"`
	ProcessDefId   int            `json:"process_def_id"`
	ProcessDefName string         `json:"process_def_name"`
	StartTime      civil.DateTime `json:"start_time"`
	EndTime        civil.DateTime ` json:"end_time"`
	CreateUser     string         `json:"create_user"`
	CreateTime     civil.DateTime `json:"create_time"`
	ApplyTime      civil.DateTime `json:"apply_time"`
	NsId           int            `json:"ns_id"`
	NsName         string         `json:"ns_name"`
	AppId          int            `json:"app_id"`
	AppName        string         `json:"app_name"`
	ClusterName    string         `json:"cluster_name"`
	Status         int            `json:"status"`
	Data           string         `json:"data"`
	Remark         string         `json:"remark"`
}
func   SelectDBPI(ctx context.Context) {
	var stinbody qo.StInBody
	err := ctx.ReadJSON(&stinbody)
	if err == nil {
		var data ProcessInstance
		mysql.GetDB().LogMode(true)
		err:=mysql.GetDB().Table("process_instance").Where("app_id=? and app_name=? and ns_id=? and ns_name=? and cluster_name=? and data=? and process_def_id=? and remark=?",
			stinbody.AppId,stinbody.AppName,stinbody.NsId,stinbody.NsName,stinbody.Data,stinbody.Cluster,stinbody.DefineId,stinbody.Remark).Scan(&data).Error
	if err != nil {
			fmt.Println("exec failed, ", err)
	}
	fmt.Println(data)
	}
	//fmt.Println(data)
}
