package controllers

import (
	"fromework/modules/po"
	"fromework/modules/qo"
	"github.com/kataras/iris/v12/context"
)

func FunPD(context context.Context) {
	context.JSON(&qo.Resp{ //返回的数据类型
		Code: 200,
		Msg:  "ok",
		Data: po.SelectDBPD(),
	})
	return
}
