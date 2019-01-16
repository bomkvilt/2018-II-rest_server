package main

import (
	// "time"
	"AForum/internal/api"
	"AForum/internal/database"
	"AForum/internal/router"

	// "fmt"
	// "github.com/pkg/profile"
	"github.com/valyala/fasthttp"
)

func main() {
	// defer profile.Start(profile.CPUProfile, profile.ProfilePath("./pro")).Stop()
	// go func() { 
	// 	for {
	// 		time.Sleep(20*time.Second)
	// 		fmt.Println(router.Times, "-----------------------------")
	// 	}
	// }()

	var (
		d = database.New()
		h = api.New(d)
		r = router.New(h)
	)
	fasthttp.ListenAndServe(":5000", r)
}
