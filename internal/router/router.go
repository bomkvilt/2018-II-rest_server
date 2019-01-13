package router

import (
	"net/http"

	"AForum/internal/api"
	"github.com/gorilla/mux"
)

type Router struct {
	mux.Router
}

func New(h *api.Handler) *Router {
	r := &Router{
		Router: *mux.NewRouter(),
	}
	r.Get(`/api/user/{nick}/profile`, h.UserGetOne)
	r.Post(`/api/user/{nick}/create`, h.UserCreate)
	r.Post(`/api/user/{nick}/profile`, h.UserUpdate)

	r.Post(`/api/forum/create`, h.ForumCreate)
	r.Get(`/api/forum/{slug}/users`, h.ForumGetUsers)
	r.Get(`/api/forum/{slug}/details`, h.ForumGetOne)
	r.Get(`/api/forum/{slug}/threads`, h.ForumGetThreads)
	r.Post(`/api/forum/{slug}/create`, h.ThreadCreate)

	r.Post(`/api/thread/{slug_or_id}/create`, h.PostsCreate)
	r.Post(`/api/thread/{slug_or_id}/vote`, h.ThreadVote)
	r.Get(`/api/thread/{slug_or_id}/details`, h.ThreadGetOne)
	r.Post(`/api/thread/{slug_or_id}/details`, h.ThreadUpdate)

	r.Get(`/api/thread/{slug_or_id}/posts`, h.ThreadGetPosts)

	r.Get(`/api/post/{id}/details`, h.PostGetOne)
	r.Post(`/api/post/{id}/details`, h.PostUpdate)

	r.Get(`/api/service/status`, h.Status)
	r.Post(`/api/service/clear`, h.Clear)

	return r
}

// ----------------||

func (r *Router) Get(uri string, h http.HandlerFunc)  { r.HandleFunc(uri, h).Methods("GET") }
func (r *Router) Post(uri string, h http.HandlerFunc) { r.HandleFunc(uri, h).Methods("POST") }
