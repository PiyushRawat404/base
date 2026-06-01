package main

import (
	"blog/internal/handler"
	"blog/internal/repository"
	"blog/internal/routes"
	"blog/internal/service"
	"blog/pkg/cache"
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

	redisClient, err := cache.LoadRedis(cfg)
	if err != nil {
		log.Fatalf("connect redis: %v", err)
	}
	defer redisClient.Close()

	blogRepo := repository.NewBlogRepository(db)
	blogService := service.NewBlogService(blogRepo)
	blogHandler := handler.NewBlogHandler(blogService)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, cfg.JWTSecret)
	userHandler := handler.NewUserHandler(userService)
	router := routes.New(blogHandler, userHandler, cfg.JWTSecret)

	addr := ":" + cfg.Port
	fmt.Printf("Server running on %s\n", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("start server: %v", err)
	}
}
