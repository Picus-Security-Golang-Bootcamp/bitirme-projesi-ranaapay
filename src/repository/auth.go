package repository

import (
	models "PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/errorHandler"
	"errors"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	authRepo := AuthRepository{db: db}
	authRepo.migrations()
	authRepo.insertSampleAdminData()
	return &authRepo
}

func (r *AuthRepository) CreateUser(user models.User) *models.User {
	result := r.db.Create(&user)
	if result.Error != nil {
		errorHandler.Panic(errorHandler.DBCreateError)
	}
	return &user
}

func (r *AuthRepository) FindUser(name string) *models.User {
	var user models.User
	result := r.db.Where(&models.User{FirstName: name}).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		errorHandler.Panic(errorHandler.NotFoundError)
	}
	if result.Error != nil {
		errorHandler.Panic(errorHandler.InternalServerError)
	}
	return &user
}

func (r *AuthRepository) migrations() {
	if err := r.db.AutoMigrate(&models.User{}); err != nil {
		errorHandler.Panic(errorHandler.DBMigrateError)
	}
}

func (r *AuthRepository) insertSampleAdminData() {
	user := models.User{
		FirstName: "admin",
		LastName:  "admin",
		Email:     "admin",
		Password:  "12345",
		Role:      0,
	}
	r.db.Create(&user)
}
