package helper

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/config"
	"fmt"
	"github.com/golang-jwt/jwt"
)

func GenerateJwtToken(user models.User, cfg config.JWTConfig) (string, error) {
	secretKey := []byte(cfg.SecretKey)
	claims := models.UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: cfg.SessionTime,
		},
		UserId: user.Id,
		Role:   user.Role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println(secretKey)
		return "", err
	}
	return tokenString, nil
}
