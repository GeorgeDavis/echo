package healthprobe

import (
	"net/http"

	"github.com/gobuffalo/logger"
	"github.com/labstack/echo/v4"
)

type HealthProbe interface {
	Alive(c echo.Context) error
}

type NewParams struct {
	Log logger.Logger
}

type healthprobe struct {
	log logger.Logger
}

func New(n NewParams) HealthProbe {
	return &healthprobe{
		log: n.Log,
	}
}

// Alive determines if the service is running
func (h *healthprobe) Alive(c echo.Context) error {
	return c.String(http.StatusOK, "Service is Alive!!")
}
