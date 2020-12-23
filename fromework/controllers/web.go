package controllers

import (
	"fromework/modules/po"
	"github.com/kataras/iris/v12"
)

func Router(app *iris.Application) {
	app.Get("/workflow/api/v1/define", FunPD)
	app.Get("/workflow/api/v1/define/{id}/node", FunPDN)
	app.Put("/workflow/api/v1/define", FunPDN_U)
	app.Post("/workflow/api/v1/instance",po.SelectDBPI)

}
