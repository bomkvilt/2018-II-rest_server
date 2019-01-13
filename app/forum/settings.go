package forum

import (
	"AForum/app/database"
	"AForum/utiles/walhalla"
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
