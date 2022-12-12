package middleware

import (
	"context"
	"log"
	"time"

	"github.com/doge-verse/easy-upgrade-backend/api/request"
	"github.com/doge-verse/easy-upgrade-backend/api/response"
	"github.com/doge-verse/easy-upgrade-backend/internal/cache"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token")
	TokenIsEmpty     = errors.New("Couldn't find the token")
)

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(viper.GetString("session.keyPairs")),
	}
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := CheckoutSession(c)
		if err != nil {
			c.Abort()
			log.Printf("JWTAuth error: %+v \n", err)
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

func CheckoutSession(c *gin.Context) (*request.CustomClaims, error) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		response.Fail(c, TokenIsEmpty)
		return nil, TokenIsEmpty
	}
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		response.Fail(c, err)
		return nil, err
	}
	err, userJwt := GetRedisJWT(claims.UserID)
	if userJwt != token {
		return nil, TokenInvalid
	}
	return claims, nil
}

// ParseToken .
func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{},
		func(token *jwt.Token) (i interface{}, e error) {
			return j.SigningKey, nil
		})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, TokenInvalid
}

// CreateToken .
func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func SignJwt(userID uint) (string, int64, error) {
	j := NewJWT()
	claims := request.CustomClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Unix() + viper.GetInt64("session.expires-time"),
			Issuer:    "easy-upgrade",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		return "", 0, errors.New("create token failed")
	}
	if err := SetRedisJWT(token, userID); err != nil {
		err = errors.Wrap(err, "set login status failed")
		return "", 0, err
	}
	return token, claims.StandardClaims.ExpiresAt * 1000, nil
}

// GetRedisJWT .
func GetRedisJWT(userID uint) (err error, redisJWT string) {
	redisJWT, err = cache.Redis.Get(context.Background(), cast.ToString(userID)).Result()
	return err, redisJWT
}

// SetRedisJWT .
func SetRedisJWT(jwt string, userID uint) (err error) {
	timer := time.Duration(viper.GetInt64("session.expires-time")) * time.Second
	return cache.Redis.Set(context.Background(), cast.ToString(userID), jwt, timer).Err()
}
