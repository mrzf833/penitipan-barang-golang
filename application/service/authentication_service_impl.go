package service

import (
	"context"
	"database/sql"
	"penitipan-barang/application/exception"
	"penitipan-barang/application/helper"
	"penitipan-barang/application/middleware"
	"penitipan-barang/application/model/domain"
	"penitipan-barang/application/model/web"
	"penitipan-barang/application/repository"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticationServiceImpl struct {
	userRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewAuthenticationService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) AuthenticationService {
	return &AuthenticationServiceImpl{
		userRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *AuthenticationServiceImpl) Login(ctx context.Context, request web.Credentials) web.LoginResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrROllback(tx)

	user, err := service.userRepository.FindByUsername(ctx, tx, request.Username)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	helper.PanicIfError(err)

	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &domain.Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(domain.JwtKey)
	helper.PanicIfError(err)

	loginResponse := web.LoginResponse{
		Id:       user.Id,
		Username: user.Username,
		Name:     user.Name,
		Role:     user.Role,
		Token:    tokenString,
	}
	return loginResponse
}

func (service *AuthenticationServiceImpl) User(ctx context.Context) web.UserResponse {
	user := middleware.AuthenticationGetUser

	userResponse := web.UserResponse{
		Id:       user.Id,
		Name:     user.Name,
		Username: user.Username,
		Role:     user.Role,
	}

	return userResponse
}
