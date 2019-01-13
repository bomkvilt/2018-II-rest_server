package main

import (
	"net/http"

	"AForum/internal/api"
	"AForum/internal/database"
	"AForum/internal/router"

	"github.com/gorilla/handlers"
)

func main() {
	var (
		d = database.New()
		h = api.New(d)
		r = router.New(h)
	)
	// http.ListenAndServe(":5000", r)
	http.ListenAndServe(":5000", handlers.RecoveryHandler()(r))
}
