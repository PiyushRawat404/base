package handler

import (
	"blog/internal/middleware"
	"blog/internal/model"
	"blog/internal/service"
	"encoding/json"
	"errors"
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

func (h *BlogHandler) Static(w http.ResponseWriter, _ *http.Request) {
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
	return parseQueryInt(r, "id")
}

func parseQueryInt(r *http.Request, key string) (int, error) {
	value := r.URL.Query().Get(key)
	if value == "" {
		return 0, fmt.Errorf("missing %s query parameter", key)
	}

	id, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("invalid %s query parameter", key)
	}

	return id, nil
}

func parseBlogID(r *http.Request) (int, error) {
	return parseQueryInt(r, "blog_id")
}

func currentUserID(r *http.Request) (int, bool) {
	claims, ok := middleware.ClaimsFromContext(r.Context())
	if !ok {
		return 0, false
	}
	return claims.UserID, true
}

func (h *BlogHandler) CreateBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	defer r.Body.Close()

	claims, ok := middleware.ClaimsFromContext(r.Context())
	if !ok {
		writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "missing authentication context"})
		return
	}

	var request model.CreateBlogRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	if err := model.Validate(request); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	blog := model.BlogPost{
		Title:        request.Title,
		Content:      request.Content,
		Author:       claims.Email,
		Is_Published: request.IsPublished,
	}

	if err := h.Service.CreateBlog(blog); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to create blog"})
		return
	}

	writeJSON(w, http.StatusCreated, map[string]string{
		"message": "blog created successfully",
	})
}

func (h *BlogHandler) GetBlogs(w http.ResponseWriter, r *http.Request) {
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

func (h *BlogHandler) GetComments(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	blogID, err := parseBlogID(r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	comments, err := h.Service.GetComments(blogID)
	if err != nil {
		if errors.Is(err, service.ErrBlogNotFound) {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "blog not found"})
			return
		}
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to fetch comments"})
		return
	}

	writeJSON(w, http.StatusOK, comments)
}

func (h *BlogHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	defer r.Body.Close()

	userID, ok := currentUserID(r)
	if !ok {
		writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "missing authentication context"})
		return
	}

	var request model.CreateCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	if err := model.Validate(request); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	comment, err := h.Service.CreateComment(model.Comment{
		BlogID:  request.BlogID,
		UserID:  userID,
		Content: request.Content,
	})
	if err != nil {
		if errors.Is(err, service.ErrBlogNotFound) {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "blog not found"})
			return
		}
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to create comment"})
		return
	}

	writeJSON(w, http.StatusCreated, comment)
}

func (h *BlogHandler) AddLike(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	defer r.Body.Close()

	userID, ok := currentUserID(r)
	if !ok {
		writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "missing authentication context"})
		return
	}

	var request model.CreateLikeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	if err := model.Validate(request); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	if err := h.Service.AddLike(request.BlogID, userID); err != nil {
		switch {
		case errors.Is(err, service.ErrBlogNotFound):
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "blog not found"})
		case errors.Is(err, service.ErrLikeExists):
			writeJSON(w, http.StatusConflict, map[string]string{"error": err.Error()})
		default:
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to add like"})
		}
		return
	}

	writeJSON(w, http.StatusCreated, map[string]string{"message": "blog liked successfully"})
}

func (h *BlogHandler) RemoveLike(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	userID, ok := currentUserID(r)
	if !ok {
		writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "missing authentication context"})
		return
	}

	blogID, err := parseBlogID(r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	if err := h.Service.RemoveLike(blogID, userID); err != nil {
		switch {
		case errors.Is(err, service.ErrBlogNotFound):
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "blog not found"})
		case errors.Is(err, service.ErrLikeNotFound):
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "like not found"})
		default:
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to remove like"})
		}
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "like removed successfully"})
}

func (h *BlogHandler) GetLikeSummary(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	blogID, err := parseBlogID(r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	summary, err := h.Service.GetLikeSummary(blogID)
	if err != nil {
		if errors.Is(err, service.ErrBlogNotFound) {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "blog not found"})
			return
		}
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to fetch likes"})
		return
	}

	writeJSON(w, http.StatusOK, summary)
}

func (h *BlogHandler) UpdateBlog(w http.ResponseWriter, r *http.Request) {
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

	claims, ok := middleware.ClaimsFromContext(r.Context())
	if !ok {
		writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "missing authentication context"})
		return
	}

	var request model.UpdateBlogRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	if err := model.Validate(request); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	blog := model.BlogPost{
		Id:           id,
		Title:        request.Title,
		Content:      request.Content,
		Author:       claims.Email,
		Is_Published: request.IsPublished,
	}

	if err := h.Service.UpdateBlog(blog); err != nil {
		if errors.Is(err, service.ErrBlogNotFound) {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "blog not found"})
			return
		}
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to update blog"})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"message": "blog updated successfully",
	})
}

func (h *BlogHandler) DeleteBlog(w http.ResponseWriter, r *http.Request) {
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
		if errors.Is(err, service.ErrBlogNotFound) {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "blog not found"})
			return
		}
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to delete blog"})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"message": "blog deleted successfully",
	})
}
