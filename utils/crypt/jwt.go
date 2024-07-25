package crypt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"pentag.kr/distimer/configs"
)

func NewJWT(userID uuid.UUID) string {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID.String()
	claims["exp"] = time.Now().Add(time.Second * time.Duration(configs.Env.JWTExpire)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(configs.Env.JWTSecret))
	return tokenString
}

func ParseIDJWT(tokenString string) (uuid.UUID, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.Env.JWTSecret), nil
	})
	if err != nil {
		return uuid.Nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.Nil, err
	}
	strID, ok := claims["user_id"].(string)
	if !ok {
		return uuid.Nil, err
	}
	id, err := uuid.Parse(strID)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
