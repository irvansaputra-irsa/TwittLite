package middlewares

import (
	"errors"
	"strings"
	"time"
	"twittlite/helpers/common"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var secretKey = []byte(viper.GetString("jwt.secret_key"))

func CreateToken(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": userId,
			"exp":     time.Now().Add(time.Hour * 1).Unix(), //set 1 hour token expired time
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := GetJwtTokenFromHeader(c)
		if err != nil {
			common.GenerateErrorResponse(c, err.Error())
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil {
			common.GenerateErrorResponse(c, err.Error())
			return
		}

		if !token.Valid {
			common.GenerateErrorResponse(c, "token is not valid, please try again")
			return
		}

		tokenClaim, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			common.GenerateErrorResponse(c, "token is not valid, please try again")
			return
		}

		if exp, ok := tokenClaim["exp"].(float64); ok {
			expTime := time.Unix(int64(exp), 0)
			if time.Now().After(expTime) {
				common.GenerateErrorResponse(c, "token expired, please log in again")
				return
			}
		} else {
			common.GenerateErrorResponse(c, "token does not have a valid expiration time")
			return
		}

		c.Set("auth", tokenClaim)

		c.Next()
	}
}

func GetJwtTokenFromHeader(c *gin.Context) (tokenString string, err error) {
	authHeader := c.Request.Header.Get("Authorization")

	if common.IsEmptyField(authHeader) {
		return tokenString, errors.New("authorization header is required")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return tokenString, errors.New("invalid Authorization header format")
	}

	return parts[1], nil
}

func EncryptToken(c *gin.Context) (userId int, err error) {
	auth, exists := c.Get("auth")
	if !exists {
		return 0, errors.New("token does not exist")
	}

	data, ok := auth.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("token is not valid")
	}
	return int(data["user_id"].(float64)), nil
}
