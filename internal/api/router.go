package api

import (
	"net/http"

	"github.com/GeorgeDavis/echo/internal/api/routes"
	"github.com/GeorgeDavis/echo/internal/handlers"
	"github.com/GeorgeDavis/echo/internal/healthprobe"
	"github.com/GeorgeDavis/echo/internal/service"
	"github.com/gobuffalo/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type SetupRouterParams struct {
	Log         logger.Logger
	HealthProbe healthprobe.HealthProbe
	UserService service.User
}

func SetupRouter(s SetupRouterParams) (http.Handler, error) {
	// Echo instance
	r := echo.New()

	// Create a new User Handler and inject the userService
	userHandler := handlers.NewUserHandler(s.Log, s.UserService)

	// Middleware
	r.Use(middleware.Logger())
	r.Use(middleware.Recover())

	r.GET(routes.Alive, s.HealthProbe.Alive)
	r.GET(routes.GetUserById, userHandler.GetUserById)

	return r, nil
}
