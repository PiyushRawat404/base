package handler

import (
	"blog-backend/internal/model"
	"blog-backend/internal/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type BlogHandler struct {
	Service *service.BlogService
}

func NewBlogHandler(service *service.BlogService) *BlogHandler {
	return &BlogHandler{
		Service: service,
	}
}

func Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server is running")
}

func (h *BlogHandler) CreateBlog(w http.ResponseWriter, r *http.Request) {

	var blog model.BlogPost

	json.NewDecoder(r.Body).Decode(&blog)

	h.Service.CreateBlog(blog)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Blog created successfully",
	})
}

func (h *BlogHandler) GetBlogs(w http.ResponseWriter, r *http.Request) {

	blogs, _ := h.Service.GetBlogs()

	json.NewEncoder(w).Encode(blogs)
}

func (h *BlogHandler) UpdateBlog(w http.ResponseWriter, r *http.Request) {

	idParam := r.URL.Query().Get("id")

	id, _ := strconv.Atoi(idParam)

	var blog model.BlogPost

	json.NewDecoder(r.Body).Decode(&blog)

	blog.ID = id

	h.Service.UpdateBlog(blog)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Blog updated successfully",
	})
}

func (h *BlogHandler) DeleteBlog(w http.ResponseWriter, r *http.Request) {

	idParam := r.URL.Query().Get("id")

	id, _ := strconv.Atoi(idParam)

	h.Service.DeleteBlog(id)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Blog deleted successfully",
	})
}
