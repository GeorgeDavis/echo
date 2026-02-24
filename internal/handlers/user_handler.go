package handlers

import (
	"errors"
	"net/http"

	"github.com/GeorgeDavis/echo/internal/service"
	"github.com/gobuffalo/logger"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	Log         logger.Logger
	UserService service.User
}

func NewUserHandler(log logger.Logger, userService service.User) *userHandler {
	return &userHandler{
		Log:         log,
		UserService: userService,
	}
}

func (u *userHandler) GetUserById(c echo.Context) error {
	userId := c.Param("id")
	if userId == "" {
		return c.JSON(http.StatusBadRequest, errors.New("User Id not in request"))
	}

	user, err := u.UserService.GetUserById(userId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
