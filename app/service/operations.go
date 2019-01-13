package service

import (
	"AForum/app/database"
	// "AForum/app/generated/models"
	"AForum/app/generated/restapi/operations/service"
	"AForum/utiles/walhalla"
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