package app

import (
	"Forum/app/middleware"
	"Forum/utiles/logger"
	"Forum/utiles/walhalla"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // psql driver
)

//go:generate go run ../utiles/walhalla/main .

// walhalla:app {
// 	globalMiddlewares  : [ log ],
// 	operationMiddlewars: [     ]
// }

var MiddlewareGeneratorsGlobal = walhalla.GlobalMiddlewareGenerationFunctionMap{
	"log": middleware.Logger,
}

type conInfo struct {
	host     string
	port     string
	user     string
	password string
	dbName   string
}

func (ci conInfo) Marshal() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		ci.host, ci.port, ci.user, ci.password, ci.dbName)
}

func SetupContext(ctx *walhalla.Context) {
	var err error
	{ // init db
		ctx.DB, err = sqlx.Open("postgres", conInfo{
			host:     "127.0.0.1",
			port:     "5432",
			user:     "postgres",
			password: "ps",
			dbName:   "postgres",
		}.Marshal())
		if err != nil {
			panic(err)
		}
	}
	{ // setup the logger
		ctx.Log, err = logger.New(logger.Config{
			File:    "log.log",
			BStdOut: true,
		})
		if err != nil {
			panic(err)
		}
	}
}
