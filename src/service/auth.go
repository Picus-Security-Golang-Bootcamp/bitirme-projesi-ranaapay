package service

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/config"
	"PicusFinalCase/src/pkg/errorHandler"
	"PicusFinalCase/src/pkg/helper"
	"PicusFinalCase/src/repository"
)

type AuthService struct {
	repo *repository.AuthRepository
	cfg  config.JWTConfig
}

func NewAuthService(cfg config.JWTConfig, repo *repository.AuthRepository) *AuthService {
	return &AuthService{
		repo: repo,
		cfg:  cfg,
	}
}

func (s *AuthService) CreateUser(user models.User) string {
	user.HashPassword()
	resUser := s.repo.CreateUser(user)
	token := helper.GenerateJwtToken(*resUser, s.cfg)
	return token
}

func (s *AuthService) LoginUser(name string, password string) string {
	resUser := s.repo.FindUser(name)
	if result := resUser.CheckPasswordHash(password); !result {
		errorHandler.Panic(errorHandler.PasswordNotTrueError)
	}
	token := helper.GenerateJwtToken(*resUser, s.cfg)
	return token
}
