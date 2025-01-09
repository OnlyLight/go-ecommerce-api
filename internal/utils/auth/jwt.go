package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/onlylight29/go-ecommerce-backend-api/global"
)

type PayloadClaims struct {
	jwt.RegisteredClaims
}

func GenTokenJWT(payload jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(global.Config.Jwt.ApiSecretKey))
}

func CreateToken(uuidToken string) (string, error) {
	// 1. Set time expire
	timeEx := global.Config.Jwt.JwtExpiration
	if timeEx == "" {
		timeEx = "1h"
	}

	expiration, err := time.ParseDuration(timeEx)
	if err != nil {
		return "", err
	}

	return GenTokenJWT(&PayloadClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.New().String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "shopdevgo",
			Subject:   uuidToken,
		},
	})
}

func ParseJwtTokenSubject(token string) (*jwt.RegisteredClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &PayloadClaims{}, func(jwtToken *jwt.Token) (interface{}, error) {
		return []byte(global.Config.Jwt.ApiSecretKey), nil
	})

	fmt.Println("tokenClaims", tokenClaims)

	if err != nil {
		fmt.Println(err)
	} else if claims, ok := tokenClaims.Claims.(*PayloadClaims); ok {
		fmt.Println(claims.RegisteredClaims.Issuer)
		return &claims.RegisteredClaims, nil
	} else {
		fmt.Println("unknown claims type, cannot proceed")
	}

	return nil, err
}

func VerifyTokenSubject(token string) (*jwt.RegisteredClaims, error) {
	claims, err := ParseJwtTokenSubject(token)

	if err != nil {
		return nil, err
	}

	return claims, nil
}
