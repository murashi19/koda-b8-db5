package lib

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func Conn() (*pgx.Conn, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Cannot read .env file")
	}
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	return conn, nil

}
