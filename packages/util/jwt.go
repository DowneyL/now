package util

import (
	"github.com/DowneyL/now/packages/configs"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

type Auth struct {
	Token     string    `json:"token"`
	ExpiredAt *DateTime `json:"expired_at"`
}

func getJwtSecret() []byte {
	config := configs.New()
	return []byte(config.GetJwtSecret())
}

func GenerateAuth(username, password string) (Auth, error) {
	now := time.Now()
	expireTime := now.Add(3 * time.Hour)
	claims := Claims{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "now",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(getJwtSecret())
	if err != nil {
		return Auth{}, err
	}

	return Auth{token, &DateTime{expireTime}}, nil
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return getJwtSecret(), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
