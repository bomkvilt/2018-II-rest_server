package thread

import (
	"github.com/bomkvilt/tech-db-ap/app/database"
	"github.com/bomkvilt/tech-db-ap/app/generated/models"
	"github.com/bomkvilt/tech-db-ap/app/generated/restapi/operations/thread"
	"github.com/bomkvilt/tech-db-ap/utiles/walhalla"
	"github.com/go-openapi/runtime/middleware"
)

// walhalla:gen
func ThreadCreate(params thread.ThreadCreateParams, ctx *walhalla.Context, model *database.DB) middleware.Responder {
	switch err := model.CreateNewThread(params.Slug, params.Thread); err.(type) {
	case *database.ErrorNotFound:
		return thread.NewThreadCreateNotFound().WithPayload(&models.Error{
			Message: "",
		})
	case *database.ErrorAlreadyExist:
		return thread.NewThreadCreateConflict().WithPayload(params.Thread)
	default:
		return thread.NewThreadCreateCreated().WithPayload(params.Thread)
	}
}

// walhalla:gen
func ThreadUpdate(params thread.ThreadUpdateParams, ctx *walhalla.Context, model *database.DB) middleware.Responder {
	res, err := model.UpdateThread(params.SlugOrID, params.Thread)
	if err != nil {
		return thread.NewThreadUpdateNotFound().WithPayload(&models.Error{
			Message: "",
		})
	}
	return thread.NewThreadUpdateOK().WithPayload(res)
}

// walhalla:gen
func ThreadGetOne(params thread.ThreadGetOneParams, ctx *walhalla.Context, model *database.DB) middleware.Responder {
	res, err := model.GetThreadBySlugOrID(params.SlugOrID)
	if err != nil {
		return thread.NewThreadGetOneNotFound().WithPayload(&models.Error{
			Message: "",
		})
	}
	return thread.NewThreadGetOneOK().WithPayload(res)
}

// walhalla:gen
func ForumGetThreads(params thread.ForumGetThreadsParams, ctx *walhalla.Context, model *database.DB) middleware.Responder {
	res, err := model.GetThreads(params)
	if err != nil {
		return thread.NewForumGetThreadsNotFound().WithPayload(&models.Error{
			Message: "",
		})
	}
	return thread.NewForumGetThreadsOK().WithPayload(res)
}

// walhalla:gen
func ThreadVote(param thread.ThreadVoteParams, ctx *walhalla.Context, model *database.DB) middleware.Responder {
	res, err := model.VoteThread(param.SlugOrID, param.Vote)
	if err != nil {
		return thread.NewThreadVoteNotFound().WithPayload(&models.Error{
			Message: "",
		})
	}
	return thread.NewThreadVoteOK().WithPayload(res)
}
