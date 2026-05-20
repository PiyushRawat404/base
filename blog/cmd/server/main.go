package main

import (
	"blog/internal/config"
	"blog/internal/db"
	"blog/internal/handler"
	"blog/internal/repository"
	"blog/internal/routes"
	"blog/internal/service"
	"context"
	"net/http"
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
	http.ListenAndServe(":"+cfg.Port, nil)
}
