package service

import (
	"context"
	"penitipan-barang/application/model/web"
)

type AuthenticationService interface {
	Login(ctx context.Context, request web.Credentials) web.LoginResponse
	User(ctx context.Context) web.UserResponse
}
