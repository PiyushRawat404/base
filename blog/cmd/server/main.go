package main

import (
	"blog/internal/handler"
	"blog/internal/repository"
	"blog/internal/routes"
	"blog/internal/service"
	"blog/pkg/config"
	"blog/pkg/db"
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	db, err := database.LoadDB(cfg)
	if err != nil {
		log.Fatalf("connect database: %v", err)
	}
	defer db.Close(context.Background())

	blogRepo := repository.NewBlogRepository(db)
	userRepo := repository.NewUserRepository(db)
	blogService := service.NewBlogService(blogRepo)
	userService := service.NewUserService(userRepo, cfg.JWTSecret)
	blogHandler := handler.NewBlogHandler(blogService)
	userHandler := handler.NewUserHandler(userService)
	routes.Routes(blogHandler, userHandler, cfg.JWTSecret)

	addr := ":" + cfg.Port
	fmt.Printf("Server running on %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("start server: %v", err)
	}
}
