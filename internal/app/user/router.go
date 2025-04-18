package user

import (
	"github.com/labstack/echo/v4"
)

func RegisterUserRoutes(g *echo.Group) {
	g.GET("", GetAllUsers)
	g.POST("", CreateUser)
}
