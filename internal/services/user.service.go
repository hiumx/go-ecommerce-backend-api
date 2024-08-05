package services

import "github.com/hiumx/go-ecommerce-backend-api/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetUserByIdService(id int) string {
	return us.userRepo.GetUserByIdRepo(id)
}
