package user

import (
	"Artista/common"
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"golang.org/x/crypto/bcrypt"
)

const loginErrorMessage = "invalid user or password"

type Service struct {
	userRepository Repository
	config         *common.Config
}

func NewService(userRepository Repository, config *common.Config) *Service {
	return &Service{
		userRepository: userRepository,
		config:         config,
	}
}

type LoginRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoginResponse struct {
	Token string `json:"token,omitempty"`
}

type ServiceContract interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
}

func (s *Service) Login(ctx context.Context, request *LoginRequest) (*LoginResponse, error) {
	if request.Username == "" || request.Password == "" {
		return nil, fmt.Errorf(loginErrorMessage)
	}

	user, err := s.userRepository.Fetch(ctx, request.Username)
	if err != nil {
		return nil, fmt.Errorf("todo")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return nil, fmt.Errorf(loginErrorMessage)
	}
	expTime := time.Now().Add(24 * 60 * time.Minute)
	claims := &common.Claims{
		Uuid:        user.ID,
		PrivateRole: user.Role,
		StandardClaims: jwt.StandardClaims{
			Audience:  "artista",
			ExpiresAt: expTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "artista-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.config.JwtSigningSecret))
	if err != nil {
		return nil, err
	}

	return &LoginResponse{Token: tokenString}, nil

}
