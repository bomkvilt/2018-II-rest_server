// Code generatederated by walhalla; Don't edit!
package restapi

import (
	code "Forum/app"
	codeforum "Forum/app/forum"
	"Forum/app/generated/restapi/operations"
	"Forum/app/generated/restapi/operations/forum"
	"Forum/app/generated/restapi/operations/post"
	"Forum/app/generated/restapi/operations/service"
	"Forum/app/generated/restapi/operations/thread"
	"Forum/app/generated/restapi/operations/user"
	codepost "Forum/app/post"
	codeservice "Forum/app/service"
	codethread "Forum/app/thread"
	codeuser "Forum/app/user"
	"Forum/utiles/walhalla"

	"crypto/tls"
	"net/http"

	"fmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

func configureServer(s *http.Server, scheme, addr string) {}
func configureFlags(api *operations.ForumAPI)             {}
func configureTLS(tlsConfig *tls.Config)                  {}

func configureAPI(api *operations.ForumAPI) http.Handler {
	api.ServeError = errors.ServeError
	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	ctx := walhalla.Context{}
	code.SetupContext(&ctx)

	api.ThreadThreadGetOneHandler = thread.ThreadGetOneHandlerFunc(handlerThreadGetOne(ctx))
	api.ThreadThreadUpdateHandler = thread.ThreadUpdateHandlerFunc(handlerThreadUpdate(ctx))
	api.ForumForumCreateHandler = forum.ForumCreateHandlerFunc(handlerForumCreate(ctx))
	api.ThreadForumGetThreadsHandler = thread.ForumGetThreadsHandlerFunc(handlerForumGetThreads(ctx))
	api.ThreadThreadCreateHandler = thread.ThreadCreateHandlerFunc(handlerThreadCreate(ctx))
	api.UserUserCreateHandler = user.UserCreateHandlerFunc(handlerUserCreate(ctx))
	api.PostPostsCreateHandler = post.PostsCreateHandlerFunc(handlerPostsCreate(ctx))
	api.ForumForumGetOneHandler = forum.ForumGetOneHandlerFunc(handlerForumGetOne(ctx))
	api.ServiceClearHandler = service.ClearHandlerFunc(handlerClear(ctx))
	api.ServiceStatusHandler = service.StatusHandlerFunc(handlerStatus(ctx))
	api.PostThreadGetPostsHandler = post.ThreadGetPostsHandlerFunc(handlerThreadGetPosts(ctx))
	api.ThreadThreadVoteHandler = thread.ThreadVoteHandlerFunc(handlerThreadVote(ctx))
	api.UserUserGetOneHandler = user.UserGetOneHandlerFunc(handlerUserGetOne(ctx))
	api.UserUserUpdateHandler = user.UserUpdateHandlerFunc(handlerUserUpdate(ctx))
	api.ForumForumGetUsersHandler = forum.ForumGetUsersHandlerFunc(handlerForumGetUsers(ctx))
	api.PostPostGetOneHandler = post.PostGetOneHandlerFunc(handlerPostGetOne(ctx))
	api.PostPostUpdateHandler = post.PostUpdateHandlerFunc(handlerPostUpdate(ctx))

	return setupGlobalMiddleware(api.Serve(setupMiddlewares), ctx)
}

type globalMiddlewareHandler struct {
	function func(rw http.ResponseWriter, r *http.Request)
}

func (g *globalMiddlewareHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.function(rw, r)
}

func setupGlobalMiddleware(handler http.Handler, ctx walhalla.Context) http.Handler {

	next := handler.ServeHTTP

	{
		gen, ok := code.MiddlewareGeneratorsGlobal["log"]
		if !ok {
			panic(fmt.Errorf("Unexpected middleware: log"))
		}

		next = gen(next, &ctx)
	}

	return &globalMiddlewareHandler{
		function: next,
	}

}

func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// ----------------| Handlers

func handlerThreadCreate(ctx walhalla.Context) thread.ThreadCreateHandlerFunc {

	paramPtr := new(thread.ThreadCreateParams)
	next := func(*http.Request) middleware.Responder {
		return codethread.ThreadCreate(*paramPtr, &ctx, codethread.NewModel(&ctx))

	}

	return func(param thread.ThreadCreateParams) middleware.Responder {
		*paramPtr = param
		return next(param.HTTPRequest)
	}
}

func handlerThreadUpdate(ctx walhalla.Context) thread.ThreadUpdateHandlerFunc {

	paramPtr := new(thread.ThreadUpdateParams)
	next := func(*http.Request) middleware.Responder {
		return codethread.ThreadUpdate(*paramPtr, &ctx, codethread.NewModel(&ctx))

	}

	return func(param thread.ThreadUpdateParams) middleware.Responder {
		*paramPtr = param
		return next(param.HTTPRequest)
	}
}

func handlerThreadGetOne(ctx walhalla.Context) thread.ThreadGetOneHandlerFunc {

	paramPtr := new(thread.ThreadGetOneParams)
	next := func(*http.Request) middleware.Responder {
		return codethread.ThreadGetOne(*paramPtr, &ctx, codethread.NewModel(&ctx))

	}

	return func(param thread.ThreadGetOneParams) middleware.Responder {
		*paramPtr = param
		return next(param.HTTPRequest)
	}
}

func handlerForumGetThreads(ctx walhalla.Context) thread.ForumGetThreadsHandlerFunc {

	paramPtr := new(thread.ForumGetThreadsParams)
	next := func(*http.Request) middleware.Responder {
		return codethread.ForumGetThreads(*paramPtr, &ctx, codethread.NewModel(&ctx))

	}

	return func(param thread.ForumGetThreadsParams) middleware.Responder {
		*paramPtr = param
		return next(param.HTTPRequest)
	}
}

func handlerThreadVote(ctx walhalla.Context) thread.ThreadVoteHandlerFunc {

	paramPtr := new(thread.ThreadVoteParams)
	next := func(*http.Request) middleware.Responder {
		return codethread.ThreadVote(*paramPtr, &ctx, codethread.NewModel(&ctx))

	}

	return func(param thread.ThreadVoteParams) middleware.Responder {
		*paramPtr = param
		return next(param.HTTPRequest)
	}
}

func handlerForumCreate(ctx walhalla.Context) forum.ForumCreateHandlerFunc {

	paramPtr := new(forum.ForumCreateParams)
	next := func(*http.Request) middleware.Responder {
		return codeforum.ForumCreate(*paramPtr, &ctx, codeforum.NewModel(&ctx))

	}

	return func(param forum.ForumCreateParams) middleware.Responder {
		*paramPtr = param
		return next(param.HTTPRequest)
	}
}

func handlerForumGetOne(ctx walhalla.Context) forum.ForumGetOneHandlerFunc {

	paramPtr := new(forum.ForumGetOneParams)
	next := func(*http.Request) middleware.Responder {
		return codeforum.ForumGetOne(*paramPtr, &ctx, codeforum.NewModel(&ctx))

	}

	return func(param forum.ForumGetOneParams) middleware.Responder {
		*paramPtr = param
		return next(param.HTTPRequest)
	}
}

func handlerForumGetUsers(ctx walhalla.Context) forum.ForumGetUsersHandlerFunc {

	paramPtr := new(forum.ForumGetUsersParams)
	next := func(*http.Request) middleware.Responder {
		return codeforum.ForumGetUsers(*paramPtr, &ctx, codeforum.NewModel(&ctx))

	}

	return func(param forum.ForumGetUsersParams) middleware.Responder {
		*paramPtr = param
		return next(param.HTTPRequest)
	}
}

func handlerUserCreate(ctx walhalla.Context) user.UserCreateHandlerFunc {

	paramPtr := new(user.UserCreateParams)
	next := func(*http.Request) middleware.Responder {
		return codeuser.UserCreate(*paramPtr, &ctx, codeuser.NewModel(&ctx))

	}

	return func(param user.UserCreateParams) middleware.Responder {
		*paramPtr = param
		return next(param.HTTPRequest)
	}
}

func handlerUserGetOne(ctx walhalla.Context) user.UserGetOneHandlerFunc {

	paramPtr := new(user.UserGetOneParams)
	next := func(*http.Request) middleware.Responder {
		return codeuser.UserGetOne(*paramPtr, &ctx, codeuser.NewModel(&ctx))

	}

	return func(param user.UserGetOneParams) middleware.Responder {
		*paramPtr = param
		return next(param.HTTPRequest)
	}
}

func handlerUserUpdate(ctx walhalla.Context) user.UserUpdateHandlerFunc {

	paramPtr := new(user.UserUpdateParams)
	next := func(*http.Request) middleware.Responder {
		return codeuser.UserUpdate(*paramPtr, &ctx, codeuser.NewModel(&ctx))

	}

	return func(param user.UserUpdateParams) middleware.Responder {
		*paramPtr = param
		return next(param.HTTPRequest)
	}
}

func handlerPostsCreate(ctx walhalla.Context) post.PostsCreateHandlerFunc {

	paramPtr := new(post.PostsCreateParams)
	next := func(*http.Request) middleware.Responder {
		return codepost.PostsCreate(*paramPtr, &ctx, codepost.NewModel(&ctx))

	}

	return func(param post.PostsCreateParams) middleware.Responder {
		*paramPtr = param
		return next(param.HTTPRequest)
	}
}

func handlerThreadGetPosts(ctx walhalla.Context) post.ThreadGetPostsHandlerFunc {

	paramPtr := new(post.ThreadGetPostsParams)
	next := func(*http.Request) middleware.Responder {
		return codepost.ThreadGetPosts(*paramPtr, &ctx, codepost.NewModel(&ctx))

	}

	return func(param post.ThreadGetPostsParams) middleware.Responder {
		*paramPtr = param
		return next(param.HTTPRequest)
	}
}

func handlerPostGetOne(ctx walhalla.Context) post.PostGetOneHandlerFunc {

	paramPtr := new(post.PostGetOneParams)
	next := func(*http.Request) middleware.Responder {
		return codepost.PostGetOne(*paramPtr, &ctx, codepost.NewModel(&ctx))

	}

	return func(param post.PostGetOneParams) middleware.Responder {
		*paramPtr = param
		return next(param.HTTPRequest)
	}
}

func handlerPostUpdate(ctx walhalla.Context) post.PostUpdateHandlerFunc {

	paramPtr := new(post.PostUpdateParams)
	next := func(*http.Request) middleware.Responder {
		return codepost.PostUpdate(*paramPtr, &ctx, codepost.NewModel(&ctx))

	}

	return func(param post.PostUpdateParams) middleware.Responder {
		*paramPtr = param
		return next(param.HTTPRequest)
	}
}

func handlerClear(ctx walhalla.Context) service.ClearHandlerFunc {

	paramPtr := new(service.ClearParams)
	next := func(*http.Request) middleware.Responder {
		return codeservice.Clear(*paramPtr, &ctx, codeservice.NewModel(&ctx))

	}

	return func(param service.ClearParams) middleware.Responder {
		*paramPtr = param
		return next(param.HTTPRequest)
	}
}

func handlerStatus(ctx walhalla.Context) service.StatusHandlerFunc {

	paramPtr := new(service.StatusParams)
	next := func(*http.Request) middleware.Responder {
		return codeservice.Status(*paramPtr, &ctx, codeservice.NewModel(&ctx))

	}

	return func(param service.StatusParams) middleware.Responder {
		*paramPtr = param
		return next(param.HTTPRequest)
	}
}
