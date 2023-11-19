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
	echoRouter := echo.New()

	// Middleware
	echoRouter.Use(middleware.Logger())
	echoRouter.Use(middleware.Recover())

	echoRouter.GET(routes.Alive, s.HealthProbe.Alive)

	return echoRouter, nil
}
