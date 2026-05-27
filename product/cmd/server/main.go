package main

import (
	"context"
	"log"
	"net/http"
	"product/internal/handler"
	"product/internal/repository"
	"product/internal/routes"
	"product/internal/service"
	"product/pkg/config"
	"product/pkg/db"
	"product/pkg/middleware"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	conn, err := db.LoadDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	if err := db.RunMigration(context.Background(), conn, "migration/product.up.sql"); err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepository(conn)
	svc := service.NewService(repo, cfg.JWTSecret)
	h := handler.NewHandler(svc)

	mux := http.NewServeMux()
	routes.Register(mux, h, middleware.JWTAuth(cfg.JWTSecret))

	addr := ":" + cfg.Port
	log.Printf("server listening on %s", addr)
	if err := http.ListenAndServe(addr, middleware.Logging(mux)); err != nil {
		log.Fatal(err)
	}
}
