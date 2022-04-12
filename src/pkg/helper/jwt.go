package helper

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/config"
	"PicusFinalCase/src/pkg/errorHandler"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
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

func VerifyToken(token string, secret string) *models.UserClaims {
	secretKey := []byte(secret)
	decodedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return secretKey, nil
	})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if !decodedToken.Valid {
		fmt.Println(5)
		return nil
	}
	claims, ok := decodedToken.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println(4)
		return nil
	}

	var userClaims models.UserClaims
	jsonString, _ := json.Marshal(claims)
	json.Unmarshal(jsonString, &userClaims)

	return &userClaims
}
