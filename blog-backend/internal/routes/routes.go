package routes

import (
	"blog-backend/internal/handler"
	"net/http"
)

func Routes(blogHandler *handler.BlogHandler) {

	http.HandleFunc("/health", handler.Health)
	http.HandleFunc("/blogs/create", blogHandler.CreateBlog)
	http.HandleFunc("/blogs", blogHandler.GetBlogs)
	http.HandleFunc("/blogs/update", blogHandler.UpdateBlog)
	http.HandleFunc("/blogs/delete", blogHandler.DeleteBlog)
}
