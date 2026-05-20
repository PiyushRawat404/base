package routes

import (
	"blog/internal/handler"
	"net/http"
)

func Routes(handler *handler.Handler) {

	http.HandleFunc("/", handler.Static)
	http.HandleFunc("/blog", handler.CreateBlog)
	http.HandleFunc("/blogs", handler.GetBlogs)
	http.HandleFunc("/blogs/update", handler.UpdateBlog)
	http.HandleFunc("/blogs/delete", handler.DeleteBlog)
}
