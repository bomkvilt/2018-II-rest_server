package post

import (
	"time"
	"Forum/app/database"
	"Forum/app/generated/restapi/operations/post"
	"Forum/utiles/walhalla"
	"github.com/go-openapi/runtime/middleware"
)


// walhalla:gen
func PostsCreate(params post.PostsCreateParams, ctx *walhalla.Context, model *database.DB) middleware.Responder {
	created := time.Now().Format("2006-01-02T15:04:05.999999999Z07:00")
	for _, post := range params.Posts {
		model.CreateNewPost(params.SlugOrID, created, post)
	}
	return post.NewPostsCreateCreated().WithPayload(params.Posts)
}

// walhalla:gen
func ThreadGetPosts(params post.ThreadGetPostsParams, ctx *walhalla.Context, model *database.DB) middleware.Responder {
	return nil
}