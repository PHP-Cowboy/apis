package models

import "github.com/golang-jwt/jwt"

type CustomClaims struct {
	ID          uint64
	NickName    string
	AuthorityId uint32
	jwt.StandardClaims
}
