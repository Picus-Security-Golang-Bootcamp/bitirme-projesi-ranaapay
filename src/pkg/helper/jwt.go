package helper

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/config"
	"PicusFinalCase/src/pkg/errorHandler"
	"github.com/golang-jwt/jwt"
)

func GenerateJwtToken(user models.User, cfg config.JWTConfig) string {
	secretKey := []byte(cfg.SecretKey)
	claims := models.UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(cfg.SessionTime) * time.Second).Unix(),
		},
		UserId: user.Id,
		Role:   user.Role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		errorHandler.Panic(errorHandler.GenerateJwtError)
	}
	return tokenString
}
