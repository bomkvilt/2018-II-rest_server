package service

import (
	"github.com/bomkvilt/tech-db-app/app/database"
	"github.com/bomkvilt/tech-db-app/utiles/walhalla"
)

// walhalla:pack {model:NewModel}

func NewModel(ctx *walhalla.Context) *database.DB {
	return database.NewModel(ctx)
}

func check(errs ...error) {
	for _, err := range errs {
		if err != nil {
			panic(err)
		}
	}
}
