package service

import (
	"github.com/GeorgeDavis/echo/internal/models"
	"github.com/gobuffalo/logger"
	"github.com/google/uuid"
)

type User interface {
	GetUserById(userId string) (models.User, error)
}

type UserParams struct {
	Logger logger.Logger
}

type userService struct {
	Log logger.Logger
}

func NewUserService(params *UserParams) *userService {
	return &userService{
		Log: params.Logger,
	}
}

func (s *userService) GetUserById(userId string) (models.User, error) {
	return models.User{
		ID:   uuid.New().String(),
		Name: "Trapper",
		Age:  43,
	}, nil
}
