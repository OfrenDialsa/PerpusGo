package service

import (
	"PerpusGo/domain"
	"PerpusGo/dto"
	"PerpusGo/internal/config"
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	conf           *config.Config
	userRepository domain.UserRepository
}

func NewAuth(cnf *config.Config, userRepository domain.UserRepository) domain.AuthService {
	return AuthService{
		conf:           cnf,
		userRepository: userRepository,
	}
}

// Login implements domain.AuthService.
func (a AuthService) Login(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error) {
	user, err := a.userRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	if user.Id == "" {
		return dto.AuthResponse{}, errors.New("authentication failed")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		return dto.AuthResponse{}, errors.New("authentication failed")
	}

	claim := jwt.MapClaims{
		"id":  user.Id,
		"exp": time.Now().Add(time.Duration(a.conf.Jwt.Exp) * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenStr, err := token.SignedString([]byte(a.conf.Jwt.Key))
	if err != nil {
		return dto.AuthResponse{}, errors.New("authentication failed")
	}

	return dto.AuthResponse{
		UserId: user.Id,
		Token:  tokenStr,
	}, nil
}

func (a AuthService) Register(ctx context.Context, req dto.RegisterRequest) (dto.RegisterResponse, error) {

	existingUser, _ := a.userRepository.FindByEmail(ctx, req.Email)
	if existingUser.Id != "" {
		return dto.RegisterResponse{}, errors.New("email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.RegisterResponse{}, errors.New("failed to hash password")
	}

	user := domain.User{
		Id:       uuid.NewString(),
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := a.userRepository.Save(ctx, user); err != nil {
		return dto.RegisterResponse{}, errors.New("failed to save user")
	}

	return dto.RegisterResponse{
		Message: "User Successfully Registered",
	}, nil
}
