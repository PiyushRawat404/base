package routes

import (
	"net/http"
	"product/internal/handler"
)

func Register(mux *http.ServeMux, h *handler.Handler, authMiddleware func(http.Handler) http.Handler) {
	mux.HandleFunc("POST /signup", h.Signup)
	mux.HandleFunc("POST /login", h.Login)
	mux.HandleFunc("GET /products", h.Products)
	mux.Handle("POST /products", authMiddleware(http.HandlerFunc(h.Products)))
	mux.HandleFunc("GET /products/{id}", h.ProductByID)
	mux.Handle("PUT /products/{id}", authMiddleware(http.HandlerFunc(h.ProductByID)))
	mux.Handle("DELETE /products/{id}", authMiddleware(http.HandlerFunc(h.ProductByID)))
}
