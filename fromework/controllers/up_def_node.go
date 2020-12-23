package controllers

import (
	"fromework/modules/po"
	"fromework/modules/qo"
	"github.com/kataras/iris/v12/context"
)

func FunPDN_U(ctx context.Context) {

	var udnbody qo.UDNBody
	error := ctx.ReadJSON(&udnbody) //解析body里面的参数
	//fmt.Println(body)
	if error == nil {
		//fmt.Println(error)
		ctx.JSON(&qo.Resp{
			Code: 200,
			Msg:  "OK",
			Data: po.SelectDBPDN_U(udnbody),
		})
		return
	} else {
		ctx.JSON(&qo.Resp{
			Code: 404,
			Msg:  "not found",
			Data: nil,
		})
		return

	}
}
