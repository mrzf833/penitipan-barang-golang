package service

import (
	"context"
	"penitipan-barang/application/model/web"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse
	UpdatePassword(ctx context.Context, request web.UserPasswordUpdateRequest)
	Delete(ctx context.Context, userId int)
	FindById(ctx context.Context, userId int) web.UserResponse
	FindAll(ctx context.Context) []web.UserResponse
}
