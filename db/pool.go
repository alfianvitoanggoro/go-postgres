package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func InitPgxPool() {
	dsn := os.Getenv("DATABASE_URL")

	var err error
	Pool, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to create pgx pool: %v", err)
	}

	fmt.Println("✅ Connected to PostgreSQL")
}
