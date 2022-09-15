package data

import (
	"context"
	"github.com/jackc/pgx/v4"
	"os"
)

func GetConnection() (*pgx.Conn, error) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return conn, err
	}
	return conn, nil
}
