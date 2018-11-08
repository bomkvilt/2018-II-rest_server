package user

import (
	"Forum/app/database"
	"Forum/app/generated/models"
	"Forum/app/generated/restapi/operations/user"
	"Forum/utiles/walhalla"
	"github.com/go-openapi/runtime/middleware"
)

// walhalla:gen
func UserCreate(param user.UserCreateParams, ctx *walhalla.Context, model *database.DB) middleware.Responder {
	usr := &models.User{
		Nickname: param.Nickname,
		Fullname: param.Profile.Fullname,
		About:    param.Profile.About,
		Email:    param.Profile.Email,
	}
	if err := model.InsertNewUser(usr); err != nil {
		usrs, _ := model.GetAllCollisions(usr)
		return user.NewUserCreateConflict().WithPayload(usrs)
	}
	return user.NewUserCreateCreated().WithPayload(usr)
}

// walhalla:gen
func UserGetOne(param user.UserGetOneParams, ctx *walhalla.Context, model *database.DB) middleware.Responder {
	usr, _, err := model.GetUserByName(param.Nickname)
	if err != nil {
		return user.NewUserGetOneNotFound().WithPayload(&models.Error{
			Message: "Can't find user by nickname: " + param.Nickname,
		})
	}
	return user.NewUserGetOneOK().WithPayload(usr)
}

// walhalla:gen
func UserUpdate(param user.UserUpdateParams, ctx *walhalla.Context, model *database.DB) middleware.Responder {
	// no author found
	if _, _, err := model.GetUserByName(param.Nickname); err != nil {
		return user.NewUserUpdateNotFound().WithPayload(&models.Error{
			Message: "Can't find user by nickname: " + param.Nickname,
		})
	}
	// sucess update
	if err := model.UpdateUser(param.Nickname, param.Profile); err == nil {
		usr, _, _ := model.GetUserByName(param.Nickname)
		return user.NewUserUpdateOK().WithPayload(usr)
	}
	// the email is already in use
	other, _, _ := model.GetUserByEmail(param.Profile.Email.String())
	return user.NewUserUpdateConflict().WithPayload(&models.Error{
		Message: "This email is already registered by user: " + other.Nickname,
	})
}
