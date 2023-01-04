package config

import "github.com/golang-jwt/jwt/v4"

var JWT_KEY = []byte("dskmdjsnjn2372i323jui2")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}
