package repository

import (
	models "PicusFinalCase/src/models"
	"errors"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	authRepo := AuthRepository{db: db}
	authRepo.migrations()
	return &authRepo
}

func (r *AuthRepository) CreateUser(user models.User) (*models.User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *AuthRepository) CheckUserNamePassword(name string) (*models.User, error) {
	var user models.User
	result := r.db.Where(&models.User{FirstName: name}).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("UserNotFound")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *AuthRepository) migrations() {
	r.db.AutoMigrate(&models.User{})
}
