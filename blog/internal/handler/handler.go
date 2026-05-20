package handler

import (
	"blog/internal/model"
	"blog/internal/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Handler struct {
	Service *service.BlogService
}

func NewBlogHandler(service *service.BlogService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) Static(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server is running")
}

func (h *Handler) CreateBlog(w http.ResponseWriter, r *http.Request) {

	var blog model.BlogPost

	json.NewDecoder(r.Body).Decode(&blog)

	h.Service.CreateBlog(blog)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Blog created successfully",
	})
}

func (h *Handler) GetBlogs(w http.ResponseWriter, r *http.Request) {

	blogs, _ := h.Service.GetBlogs()

	json.NewEncoder(w).Encode(blogs)
}

func (h *Handler) UpdateBlog(w http.ResponseWriter, r *http.Request) {

	idParam := r.URL.Query().Get("id")

	id, _ := strconv.Atoi(idParam)

	var blog model.BlogPost

	json.NewDecoder(r.Body).Decode(&blog)

	blog.Id = id

	h.Service.UpdateBlog(blog)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Blog updated successfully",
	})
}

func (h *Handler) DeleteBlog(w http.ResponseWriter, r *http.Request) {

	idParam := r.URL.Query().Get("id")

	id, _ := strconv.Atoi(idParam)

	h.Service.DeleteBlog(id)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Blog deleted successfully",
	})
}
