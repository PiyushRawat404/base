package main

import (
	"blog-backend/internal/config"
	"blog-backend/internal/database"
	"blog-backend/internal/handler"
	"blog-backend/internal/repository"
	"blog-backend/internal/routes"
	"blog-backend/internal/service"
	"context"
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Blog-backend starting")
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
