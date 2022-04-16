package helper

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/config"
	"PicusFinalCase/src/pkg/errorHandler"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
	"time"
)

// GenerateJwtToken Generates tokens based on incoming user model and secret key.
func GenerateJwtToken(user models.User, cfg config.JWTConfig) string {

	secretKey := []byte(cfg.SecretKey)

	//Custom type User Claims declared.
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
		log.Error("Couldn't get signed token : %v", err)
		errorHandler.Panic(errorHandler.GenerateJwtError)
	}

	log.Info("The token was successfully generated.")
	return tokenString
}

// VerifyToken The incoming encrypted token is checked according to the hash method and
//secret key. If token is valid, UserClaims is returned.
func VerifyToken(token string, secret string) *models.UserClaims {

	secretKey := []byte(secret)

	decodedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return secretKey, nil
	})
	if err != nil {
		log.Error("Jwt token parse error : %v", err)
		return nil
	}
	if !decodedToken.Valid {
		log.Error("Jwt token not valid ")
		return nil
	}

	//Gets the carried claims in the token
	claims, ok := decodedToken.Claims.(jwt.MapClaims)
	if !ok {
		log.Error("Jwt token not valid ")
		return nil
	}

	//Convert claims to UserClaims type
	var userClaims models.UserClaims
	jsonString, _ := json.Marshal(claims)
	json.Unmarshal(jsonString, &userClaims)

	log.Info("Successfully generated UserClaims from the token.")
	return &userClaims
}
