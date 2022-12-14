package util

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

var (
	expires     = 30 * 365 * 24 * time.Hour
	tokenKey    = "user-token"
	tokenSecret = "3a7f3df7-50b3-4699-abaa-c0d191e18b2d"
)

// Token .
type Token struct {
	UserID   uint      `json:"userID"`
	Token    string    `json:"token"`
	ExpireAt time.Time `json:"expireAt"`
}

// Sign token
func Sign(userID uint) (*Token, error) {
	now := time.Now()
	expiresAt := now.Add(expires)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		Id:        fmt.Sprintf("%d", userID),
		ExpiresAt: expiresAt.Unix(),
		IssuedAt:  now.Unix(),
		Issuer:    tokenKey,
	})
	tokenString, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return nil, err
	}
	return &Token{
		UserID:   userID,
		Token:    tokenString,
		ExpireAt: expiresAt,
	}, nil
}

// Unsign token
func Unsign(tokenString string) (uint, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(tokenSecret), nil
		})
	if err != nil {
		logrus.Infof(" %+v err with %+v", err, token)
	}
	userID, err := strconv.Atoi(token.Claims.(*jwt.StandardClaims).Id)
	if err != nil {
		return 0, err
	}
	return uint(userID), nil
}
