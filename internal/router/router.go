package router

import (
	"github.com/alfianvitoanggoro/go-postgres/internal/app/user"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	api := e.Group("/api/v1")

	// Group User
	user.RegisterUserRoutes(api.Group("/users"))
}
