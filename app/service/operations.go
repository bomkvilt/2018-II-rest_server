package service

import (
	"github.com/bomkvilt/tech-db-app/app/database"
	// "github.com/bomkvilt/tech-db-app/app/generated/models"
	"github.com/bomkvilt/tech-db-app/app/generated/restapi/operations/service"
	"github.com/bomkvilt/tech-db-app/utiles/walhalla"
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