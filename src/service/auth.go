package service

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/config"
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

func (s *AuthService) CreateUser(user models.User) (string, error) {
	if err := user.HashPassword(); err != nil {
		return "", err
	}
	resUser, err := s.repo.CreateUser(user)
	if err != nil {
		return "", err
	}
	token, err := helper.GenerateJwtToken(*resUser, s.cfg)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *AuthService) LoginUser(name string, password string) (string, error) {
	resUser, err := s.repo.CheckUserNamePassword(name)
	if err != nil {
		return "", err
	}
	if result := resUser.CheckPasswordHash(password); !result {
		return "", err
	}
	token, err := helper.GenerateJwtToken(*resUser, s.cfg)
	if err != nil {
		return "", err
	}
	return token, nil
}
