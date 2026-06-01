package routes

import (
	"blog/internal/handler"
	"blog/internal/middleware"
	"blog/internal/model"
	"net/http"
)

type Router struct {
	mux       *http.ServeMux
	jwtSecret string
}

func New(blogHandler *handler.BlogHandler, userHandler *handler.UserHandler, jwtSecret string) *http.ServeMux {
	router := &Router{
		mux:       http.NewServeMux(),
		jwtSecret: jwtSecret,
	}

	router.registerHealthRoutes(blogHandler)
	router.registerAuthRoutes(userHandler)
	router.registerBlogRoutes(blogHandler)
	router.registerUserRoutes(userHandler)

	return router.mux
}

func (r *Router) registerHealthRoutes(blogHandler *handler.BlogHandler) {
	r.public("/", blogHandler.Static)
}

func (r *Router) registerAuthRoutes(userHandler *handler.UserHandler) {
	r.public("/auth/register", userHandler.Register)
	r.public("/auth/login", userHandler.Login)
}

func (r *Router) registerBlogRoutes(blogHandler *handler.BlogHandler) {
	r.public("/blogs", blogHandler.GetBlogs)
	r.admin("/blogs/create", blogHandler.CreateBlog)
	r.admin("/blogs/update", blogHandler.UpdateBlog)
	r.admin("/blogs/delete", blogHandler.DeleteBlog)

	r.public("/blogs/comments", blogHandler.GetComments)
	r.protected("/blogs/comments/create", blogHandler.CreateComment)

	r.public("/blogs/likes", blogHandler.GetLikeSummary)
	r.protected("/blogs/likes/add", blogHandler.AddLike)
	r.protected("/blogs/likes/remove", blogHandler.RemoveLike)
}

func (r *Router) registerUserRoutes(userHandler *handler.UserHandler) {
	r.protected("/users/me", userHandler.Me)
	r.admin("/admin/users", userHandler.ListUsers)
}

func (r *Router) public(pattern string, handler http.HandlerFunc) {
	r.mux.HandleFunc(pattern, handler)
}

func (r *Router) protected(pattern string, handler http.HandlerFunc) {
	r.mux.HandleFunc(pattern, middleware.Authenticate(r.jwtSecret, handler))
}

func (r *Router) admin(pattern string, handler http.HandlerFunc) {
	r.mux.HandleFunc(
		pattern,
		middleware.Authenticate(
			r.jwtSecret,
			middleware.RequireRole(model.AdminRole, handler),
		),
	)
}
