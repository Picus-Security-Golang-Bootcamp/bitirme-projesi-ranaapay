package repository

import (
	models "PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/errorHandler"
	log "github.com/sirupsen/logrus"
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

// CreateUser Creates a user in the database. Returns user.
func (r *AuthRepository) CreateUser(user models.User) *models.User {

	result := r.db.Create(&user)
	if result.Error != nil {
		log.Errorf("Create User Error : %s", result.Error.Error())
		return nil
	}

	return &user
}

// FindUser Finds the user based on the entered name parameter. Returns user.
func (r *AuthRepository) FindUser(name string) *models.User {
	var user models.User

	result := r.db.Where(&models.User{FirstName: name}).First(&user)
	if result.Error != nil {
		log.Errorf("Find User Error : %s", result.Error.Error())
		return nil
	}

	return &user
}

func (r *AuthRepository) migrations() {
	if err := r.db.AutoMigrate(&models.User{}); err != nil {
		log.Errorf("User Migration Error : %v", err)
		errorHandler.Panic(errorHandler.DBMigrateError)
	}
}

//A statically maintained admin is created in the database.
func (r *AuthRepository) insertSampleAdminData() {
	user := models.User{
		FirstName: "admin",
		LastName:  "admin",
		Email:     "admin",
		Password:  "12345",
		Role:      0,
	}
	user.HashPassword()
	r.db.Create(&user)
	log.Info("Admin user created in database.")
}
