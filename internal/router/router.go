package router

import (
	// "time"
	"AForum/internal/api"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

type Router struct {
	fasthttprouter.Router
}

func New(h *api.Handler) fasthttp.RequestHandler {
	r := &Router{
		Router: *fasthttprouter.New(),
	}
	r.GET(`/api/user/:nick/profile`, h.UserGetOne)
	r.POST(`/api/user/:nick/create`, h.UserCreate)
	r.POST(`/api/user/:nick/profile`, h.UserUpdate)

	// r.POST(`/api/forum/create`, h.ForumCreate)
	r.GET(`/api/forum/:slug/users`, h.ForumGetUsers)
	r.GET(`/api/forum/:slug/details`, h.ForumGetOne)
	r.GET(`/api/forum/:slug/threads`, h.ForumGetThreads)
	r.POST(`/api/forum/:slug/create`, h.ThreadCreate)

	r.POST(`/api/thread/:slug_or_id/create`, h.PostsCreate)
	r.POST(`/api/thread/:slug_or_id/vote`, h.ThreadVote)
	r.GET(`/api/thread/:slug_or_id/details`, h.ThreadGetOne)
	r.POST(`/api/thread/:slug_or_id/details`, h.ThreadUpdate)

	r.GET(`/api/thread/:slug_or_id/posts`, h.ThreadGetPosts)

	r.GET(`/api/post/:id/details`, h.PostGetOne)
	r.POST(`/api/post/:id/details`, h.PostUpdate)

	r.GET(`/api/service/status`, h.Status)
	r.POST(`/api/service/clear`, h.Clear)

	return func(ctx *fasthttp.RequestCtx) {
		// defer func (t time.Time) {
		// 	ms := time.Since(t) / time.Microsecond
		// 	if ms > 60*1000 {
		// 		println(ms, string(ctx.Path())+string(ctx.URI().QueryString()))
		// 	}
		// }(time.Now())

		path := string(ctx.Path())
		if path == "/api/forum/create" {
			h.ForumCreate(ctx)
			return
		}
		r.Handler(ctx)
	}
}
