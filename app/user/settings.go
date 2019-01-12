package user

import (
	"github.com/bomkvilt/tech-db-ap/app/database"
	"github.com/bomkvilt/tech-db-ap/utiles/walhalla"
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
