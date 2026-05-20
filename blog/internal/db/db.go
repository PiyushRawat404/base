package database

import (
	"context"
	"fmt"

	"blog/internal/config"

	"github.com/jackc/pgx/v5"
)

func LoadDB(cfg config.Config) (*pgx.Conn, error) {

	dsn := "postgres://" + cfg.DBUser + ":" + cfg.DBPassword + "@" + cfg.DBHost + ":" + cfg.DBPort + "/" + cfg.DBName
	conn, err := pgx.Connect(
		context.Background(),
		dsn,
	)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to postgres ")
	return conn, nil

}
