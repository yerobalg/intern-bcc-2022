package services

import (
	"clean-arch-2/app/repositories"
	"clean-arch-2/app/models"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return UserService{repo: userRepo}
}

func (s UserService) Register(model *models.Users)  (error) {
	return s.repo.Register(model)
}