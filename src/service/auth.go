package service

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/config"
	"PicusFinalCase/src/pkg/errorHandler"
	"PicusFinalCase/src/pkg/helper"
	"PicusFinalCase/src/repository"
	log "github.com/sirupsen/logrus"
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

// CreateUser Hash user password. Creates a user in the database. Returns encrypted token.
func (s *AuthService) CreateUser(user models.User) string {

	//Encrypts the user's password. Assigns it to the user password.
	user.HashPassword()

	resUser := s.repo.CreateUser(user)
	if resUser == nil {
		log.Error(errorHandler.DBCreateError)
		errorHandler.Panic(errorHandler.DBCreateError)
	}

	token := helper.GenerateJwtToken(*resUser, s.cfg)
	return token
}

// LoginUser According to the replies from the repo, if the username and password are correct,
// token generated and returned.  If one of them is wrong, it throws panic.
func (s *AuthService) LoginUser(name string, password string) string {

	resUser := s.repo.FindUser(name)
	if resUser == nil {
		log.Error("Couldn't find any user with first name is name.")
		errorHandler.Panic(errorHandler.FirstNameError)
	}

	if result := resUser.CheckPasswordHash(password); !result {
		log.Error("Password that user entered not true.")
		errorHandler.Panic(errorHandler.PasswordNotTrueError)
	}

	token := helper.GenerateJwtToken(*resUser, s.cfg)

	return token
}
