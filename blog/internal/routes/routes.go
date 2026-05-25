package routes

import (
	"blog/internal/handler"
	"blog/pkg/middleware"
	"net/http"
)

func Routes(blogHandler *handler.Handler, userHandler *handler.UserHandler, jwtSecret string) {
	authMiddleware := middleware.AuthMiddleware(jwtSecret)

	http.HandleFunc("/", blogHandler.Static)
	http.Handle("/blog", authMiddleware(http.HandlerFunc(blogHandler.CreateBlog)))
	http.HandleFunc("/blogs", blogHandler.GetBlogs)
	http.Handle("/blogs/update", authMiddleware(http.HandlerFunc(blogHandler.UpdateBlog)))
	http.Handle("/blogs/delete", authMiddleware(http.HandlerFunc(blogHandler.DeleteBlog)))
	http.HandleFunc("/users/register", userHandler.RegisterUser)
	http.HandleFunc("/users/login", userHandler.LoginUser)
}
