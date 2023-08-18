package config

import "github.com/golang-jwt/jwt/v4"

var JWT_KEY = []byte("rqje9riu320thqt08gfhqe09ghq90hg904gh90h")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}
