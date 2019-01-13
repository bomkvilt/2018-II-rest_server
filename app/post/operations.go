package post

import (
	"time"
	"AForum/app/database"
	"AForum/app/generated/models"
	"AForum/app/generated/restapi/operations/post"
	"AForum/utiles/walhalla"
	"github.com/go-openapi/runtime/middleware"
)


// walhalla:gen
func PostsCreate(params post.PostsCreateParams, ctx *walhalla.Context, model *database.DB) middleware.Responder {
	if _, err := model.GetThreadBySlugOrID(params.SlugOrID); err != nil {
		return post.NewPostsCreateNotFound().WithPayload(&models.Error{
			Message: "Can't find post thread by id: "+params.SlugOrID,
		})	
	}

	created := time.Now().Format("2006-01-02T15:04:05.999999999Z07:00")
	for _, ps := range params.Posts {
		err := model.CreateNewPost(params.SlugOrID, created, ps)
		if err == nil {
			continue
		}

		switch err.(type) {
		case *database.ErrorConflict:
			return post.NewPostsCreateConflict().WithPayload(&models.Error{
				Message: err.Error(),
			})
		// case *database.ErrorNotFound:
		default:
			return post.NewPostsCreateNotFound().WithPayload(&models.Error{
				Message: err.Error(),
			})
		}
	}
	return post.NewPostsCreateCreated().WithPayload(params.Posts)
}

// walhalla:gen
func ThreadGetPosts(params post.ThreadGetPostsParams, ctx *walhalla.Context, model *database.DB) middleware.Responder {
	res, err := model.GetPosts(params)
	if err != nil {
		return post.NewThreadGetPostsNotFound().WithPayload(&models.Error{
			Message: "",
		})
	}
	return post.NewThreadGetPostsOK().WithPayload(res)
}

// walhalla:gen
func PostGetOne(params post.PostGetOneParams, ctx *walhalla.Context, model *database.DB) middleware.Responder {
	res, err := model.GetPost(params.ID, params.Related)
	if err != nil {
		return post.NewPostGetOneNotFound().WithPayload(&models.Error{
			Message: "",
		})
	}
	return post.NewPostGetOneOK().WithPayload(res)
}

// walhalla:gen
func PostUpdate(params post.PostUpdateParams, ctx *walhalla.Context, model *database.DB) middleware.Responder {
	res, err := model.UpdatePost(params.ID, params.Post)
	if err != nil {
		return post.NewPostUpdateNotFound().WithPayload(&models.Error{
			Message: "",
		})
	}
	return post.NewPostUpdateOK().WithPayload(res)
}