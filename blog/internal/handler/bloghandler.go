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

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, `{"error":"failed to encode response"}`, http.StatusInternalServerError)
	}
}

func parseID(r *http.Request) (int, error) {
	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		return 0, fmt.Errorf("missing id query parameter")
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return 0, fmt.Errorf("invalid id query parameter")
	}

	return id, nil
}

func (h *Handler) CreateBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	defer r.Body.Close()

	var blog model.BlogPost

	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	if err := h.Service.CreateBlog(blog); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to create blog"})
		return
	}

	writeJSON(w, http.StatusCreated, map[string]string{
		"message": "Blog created successfully",
	})
}

func (h *Handler) GetBlogs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	blogs, err := h.Service.GetBlogs()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to fetch blogs"})
		return
	}

	writeJSON(w, http.StatusOK, blogs)
}

func (h *Handler) UpdateBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	defer r.Body.Close()

	id, err := parseID(r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	var blog model.BlogPost

	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	blog.Id = id

	if err := h.Service.UpdateBlog(blog); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to update blog"})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"message": "Blog updated successfully",
	})
}

func (h *Handler) DeleteBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	id, err := parseID(r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	if err := h.Service.DeleteBlog(id); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to delete blog"})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"message": "Blog deleted successfully",
	})
}
