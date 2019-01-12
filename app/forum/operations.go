package forum

import (
	"github.com/bomkvilt/tech-db-app/app/database"
	"github.com/bomkvilt/tech-db-app/app/generated/models"
	"github.com/bomkvilt/tech-db-app/app/generated/restapi/operations/forum"
	"github.com/bomkvilt/tech-db-app/utiles/walhalla"
	"github.com/go-openapi/runtime/middleware"
)

// walhalla:gen
func ForumCreate(params forum.ForumCreateParams, ctx *walhalla.Context, model *database.DB) middleware.Responder {
	switch model.CreateNewForum(params.Forum).(type) {
	case *database.ErrorNotFound:
		return forum.NewForumCreateNotFound().WithPayload(&models.Error{
			Message: "Can't find user with nickname: " + params.Forum.User,
		})
	case *database.ErrorAlreadyExist:
		return forum.NewForumCreateConflict().WithPayload(params.Forum)
	}
	return forum.NewForumCreateCreated().WithPayload(params.Forum)
}

// walhalla:gen
func ForumGetOne(params forum.ForumGetOneParams, ctx *walhalla.Context, model *database.DB) middleware.Responder {
	if frm, _, err := model.GetForumBySlug(params.Slug); err == nil {
		return forum.NewForumGetOneOK().WithPayload(frm)
	}
	return forum.NewForumGetOneNotFound().WithPayload(&models.Error{
		Message: "",
	})
}

// walhalla:gen
func ForumGetUsers(params forum.ForumGetUsersParams, ctx *walhalla.Context, model *database.DB) middleware.Responder {
	res, err := model.GetForumUsers(params)
	if err != nil {
		return forum.NewForumGetUsersNotFound().WithPayload(&models.Error{
			Message: "",
		})
	}
	return forum.NewForumGetUsersOK().WithPayload(res)
}
