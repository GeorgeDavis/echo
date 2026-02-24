package main

import (
	"net/http"
	"time"

	"github.com/GeorgeDavis/echo/internal/api"
	"github.com/GeorgeDavis/echo/internal/healthprobe"
	"github.com/GeorgeDavis/echo/internal/service"
	"github.com/gobuffalo/logger"
)

func setupServer(log logger.Logger) *http.Server {
	healthProbeService := healthprobe.New(healthprobe.NewParams{
		Log: log,
	})

	userService := service.NewUserService(&service.UserParams{
		Logger: log,
	})

	router, err := api.SetupRouter(api.SetupRouterParams{
		Log:         log,
		HealthProbe: healthProbeService,
		UserService: userService,
	})
	if err != nil {
		panic(err)
	}

	return &http.Server{
		Addr:         ":3000",
		WriteTimeout: time.Second * 60,
		ReadTimeout:  time.Second * 60,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

}
