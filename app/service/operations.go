package service

import (
	"Forum/app/database"
	// "Forum/app/generated/models"
	"Forum/app/generated/restapi/operations/service"
	"Forum/utiles/walhalla"
	"github.com/go-openapi/runtime/middleware"
)

// walhalla:gen
func Clear(params service.ClearParams, ctx *walhalla.Context, model *database.DB) middleware.Responder {
	model.Clear()
	return service.NewClearOK()
}

// walhalla:gen
func Status(params service.StatusParams, ctx *walhalla.Context, model *database.DB) middleware.Responder {
	return service.NewStatusOK().WithPayload(model.GetStatus())
}