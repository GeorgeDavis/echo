package api

import (
	"net/http"

	"github.com/GeorgeDavis/echo/internal/api/routes"
	"github.com/GeorgeDavis/echo/internal/healthprobe"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type SetupRouterParams struct {
	HealthProbe healthprobe.HealthProbe
	// HandlerFactoryV1 handler.Factory
}

func SetupRouter(s SetupRouterParams) (http.Handler, error) {
	// Echo instance
	r := echo.New()

	// Middleware
	r.Use(middleware.Logger())
	r.Use(middleware.Recover())

	r.GET(routes.Alive, s.HealthProbe.Alive)

	return r, nil
}
