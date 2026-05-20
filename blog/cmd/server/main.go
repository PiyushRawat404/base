package main

import (
	"blog/pkg/config"
	"blog/pkg/db"
	"blog/internal/handler"
	"blog/internal/repository"
	"blog/internal/routes"
	"blog/internal/service"
	"context"
	"net/http"
	"fmt"
)

func main() {
	cfg := config.LoadConfig()
	db, err := database.LoadDB(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close(context.Background())
	blogRepo := repository.NewBlogRepository(db)
	blogService := service.NewBlogService(blogRepo)
	blogHandler := handler.NewBlogHandler(blogService)
	routes.Routes(blogHandler)
	fmt.Println("Server running on :8000")
	http.ListenAndServe(":"+cfg.Port, nil)
}
