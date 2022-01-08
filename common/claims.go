package common

import (
	"Artista/domain"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Uuid        string      `json:"uuid"`
	PrivateRole domain.Role `json:"privateRole"`
	jwt.StandardClaims
}
