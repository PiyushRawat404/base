package db

import (
	"context"
	"fmt"
	"os"
	"product/pkg/config"

	"github.com/jackc/pgx/v5"
)

func LoadDB(cfg config.Env) (*pgx.Conn, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	return pgx.Connect(context.Background(), dsn)
}

func RunMigration(ctx context.Context, conn *pgx.Conn, path string) error {
	query, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx, string(query))
	return err
}
