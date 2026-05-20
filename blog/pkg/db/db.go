package database

import (
	"context"
	"fmt"
	"os"

	"blog/pkg/config"

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

	schema, err := os.ReadFile("migrations/schema.up.sql")
	if err != nil {
		conn.Close(context.Background())
		return nil, err
	}

	_, err = conn.Exec(context.Background(), string(schema))
	if err != nil {
		conn.Close(context.Background())
		return nil, err
	}

	fmt.Println("Connected to postgres ")
	return conn, nil

}
