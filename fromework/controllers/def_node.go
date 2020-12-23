package controllers

import (
	"fromework/modules/po"
	"fromework/modules/qo"
	"github.com/kataras/iris/v12/context"
)

func FunPDN(ctx context.Context) {
	//获取前端传过来的id
	if id, err := ctx.Params().GetInt("id"); err == nil {
		ctx.JSON(&qo.Resp{
			Code: 200,
			Msg:  "ok",
			Data: po.SelectDBPDN(id), //带着参数id进行数据库的查询
		})
		return
	} else {
		ctx.JSON(&qo.Resp{
			Code: 404,
			Msg:  " not found",
			Data: nil,
		})
		return
	}
}
