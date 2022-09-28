package service

import (
	"context"
	"database/sql"
	"fmt"
	"penitipan-barang/application/exception"
	"penitipan-barang/application/helper"
	"penitipan-barang/application/model/domain"
	"penitipan-barang/application/model/web"
	"penitipan-barang/application/repository"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		userRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrROllback(tx)

	password := []byte(request.Password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	user := domain.User{
		Name:     request.Name,
		Username: request.Username,
		Password: string(hashedPassword),
		Role:     "admin",
	}

	user = service.userRepository.Save(ctx, tx, user)
	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrROllback(tx)

	user, err := service.userRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	user.Name = request.Name
	user.Username = request.Username

	user = service.userRepository.Update(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) UpdatePassword(ctx context.Context, request web.UserPasswordUpdateRequest) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrROllback(tx)

	user, err := service.userRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	password := []byte(request.Password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	fmt.Println(string(hashedPassword))

	user.Password = string(hashedPassword)
	fmt.Println(user)

	service.userRepository.Update(ctx, tx, user)
}

func (service *UserServiceImpl) Delete(ctx context.Context, userId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrROllback(tx)

	user, err := service.userRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.userRepository.Delete(ctx, tx, user)
}

func (service *UserServiceImpl) FindById(ctx context.Context, userId int) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrROllback(tx)

	user, err := service.userRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrROllback(tx)

	users := service.userRepository.FindAll(ctx, tx)
	return helper.ToUserResponses(users)
}
