package main

import (
	"fmt"
	"os"

	"github.com/alfianvitoanggoro/go-postgres/config"
	"github.com/alfianvitoanggoro/go-postgres/db"
	"github.com/alfianvitoanggoro/go-postgres/internal/migration"
	"github.com/alfianvitoanggoro/go-postgres/internal/router"

	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadEnv()

	// Jalankan migrasi jika ENV=DEV
	if os.Getenv("ENV") == "DEV" {
		fmt.Println("Running migration...")
		migration.RunMigrations(os.Getenv("DATABASE_URL"))
	}

	// Init DB
	db.InitPgxPool()

	// Init Echo
	e := echo.New()
	router.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":" + "8080"))
}
