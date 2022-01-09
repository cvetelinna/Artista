package user

import (
	"Artista/common"
	"Artista/domain"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const loginErrorMessage = "invalid user or password"
const registerEmptyErrorMessage = "please enter username or password"
const registerUsernameError = "username should be at least 5 characters"
const registerPasswordError = "password should contain at least 5 characters"

type Service interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Register(ctx context.Context, request *RegisterRequest) (*RegisterResponse, error)
}

type service struct {
	userRepository Repository
	config         *common.Config
}

type LoginRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoginResponse struct {
	Token string `json:"token,omitempty"`
}

type RegisterRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type RegisterResponse struct {
	Token string `json:"token,omitempty"`
}

func NewService(userRepository Repository, config *common.Config) *service {
	return &service{
		userRepository: userRepository,
		config:         config,
	}
}

func (s *service) Login(ctx context.Context, request *LoginRequest) (*LoginResponse, error) {
	if request.Username == "" || request.Password == "" {
		return nil, fmt.Errorf(loginErrorMessage)
	}

	user, err := s.userRepository.Fetch(ctx, request.Username)
	if err != nil {
		return nil, err
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

func (s *service) Register(ctx context.Context, request *RegisterRequest) (*RegisterResponse, error) {
	if request.Username == "" || request.Password == "" {
		return nil, fmt.Errorf(registerEmptyErrorMessage)
	}
	if len(request.Password) < 5 {
		return nil, fmt.Errorf(registerPasswordError)
	}
	if len(request.Username) < 5 {
		return nil, fmt.Errorf(registerUsernameError)
	}

	existing, err := s.userRepository.Fetch(ctx, request.Username)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
	}
	if existing.ID != "" {
		return nil, fmt.Errorf("username %s is taken", request.Username)
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Username: request.Username,
		Password: string(hashed),
		Role:     domain.Regular,
	}
	err = s.userRepository.Insert(ctx, user)
	if err != nil {
		return nil, err
	}

	res, err := s.Login(ctx, &LoginRequest{
		Username: request.Username,
		Password: request.Password,
	})

	return &RegisterResponse{
		Token: res.Token,
	}, nil
}
