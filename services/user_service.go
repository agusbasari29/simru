package services

import "github.com/sip/simru/repository"

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *userService {
	return &userService{userRepo}
}
