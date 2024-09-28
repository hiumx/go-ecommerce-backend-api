package services

import (
	"github.com/hiumx/go-ecommerce-backend-api/internal/repo"
	"github.com/hiumx/go-ecommerce-backend-api/pkg/response"
)

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	repo repo.IUserRepository
	//... others repo
}

func (us *userService) Register(email string, purpose string) int {
	if us.repo.GetUserByEmail(email) {
		return response.ErrEmailAlreadyExisted
	}

	return response.ErrCodeSuccess
}

func NewUserService(
	ur repo.IUserRepository,
) IUserService {
	return &userService{
		repo: ur,
	}
}
