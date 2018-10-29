package app

import (
	"Forum/utiles/walhalla"
)

//go:generate go run ../utiles/walhalla/main .

// walhalla:app {
// 	globalMiddlewares  : [ ],
// 	operationMiddlewars: [ ]
// }

func SetupContext(ctx *walhalla.Context) {

}
