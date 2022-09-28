package service

import (
	"context"
	"penitipan-barang/application/model/web"
)

type StudentService interface {
	Create(ctx context.Context, request web.StudentCreateRequest) web.StudentResponse
	Update(ctx context.Context, request web.StudentUpdateRequest) web.StudentResponse
	Delete(ctx context.Context, studentId int)
	FindById(ctx context.Context, studentId int) web.StudentResponse
	FindAll(ctx context.Context) []web.StudentResponse
}
