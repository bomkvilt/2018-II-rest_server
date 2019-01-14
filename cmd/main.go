package main

import (
	"AForum/internal/api"
	"AForum/internal/database"
	"AForum/internal/router"

	"github.com/pkg/profile"
	"github.com/valyala/fasthttp"
)

func main() {
	defer profile.Start(profile.CPUProfile, profile.ProfilePath("./pro")).Stop()
	defer func() { println("----------------------------------------------") }()

	var (
		d = database.New()
		h = api.New(d)
		r = router.New(h)
	)
	fasthttp.ListenAndServe(":5000", r)
}
