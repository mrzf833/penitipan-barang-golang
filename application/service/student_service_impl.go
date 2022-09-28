package service

import (
	"context"
	"database/sql"
	"penitipan-barang/application/exception"
	"penitipan-barang/application/helper"
	"penitipan-barang/application/model/domain"
	"penitipan-barang/application/model/web"
	"penitipan-barang/application/repository"

	"github.com/go-playground/validator/v10"
)

type StudentServiceImpl struct {
	studentRepository repository.StudentRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewStudentService(studentRepository repository.StudentRepository, DB *sql.DB, validate *validator.Validate) StudentService {
	return &StudentServiceImpl{
		studentRepository: studentRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *StudentServiceImpl) Create(ctx context.Context, request web.StudentCreateRequest) web.StudentResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrROllback(tx)

	student := domain.Student{
		Fullname:    request.Fullname,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		Address:     request.Address,
	}

	student = service.studentRepository.Save(ctx, tx, student)
	return helper.ToStudentResponse(student)
}

func (service *StudentServiceImpl) Update(ctx context.Context, request web.StudentUpdateRequest) web.StudentResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrROllback(tx)

	student, err := service.studentRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	student.Fullname = request.Fullname
	student.Email = request.Email
	student.Email = request.Email
	student.Address = request.Address

	student = service.studentRepository.Update(ctx, tx, student)

	return helper.ToStudentResponse(student)
}

func (service *StudentServiceImpl) Delete(ctx context.Context, studentId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrROllback(tx)

	student, err := service.studentRepository.FindById(ctx, tx, studentId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.studentRepository.Delete(ctx, tx, student)
}

func (service *StudentServiceImpl) FindById(ctx context.Context, studentId int) web.StudentResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrROllback(tx)

	student, err := service.studentRepository.FindById(ctx, tx, studentId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToStudentResponse(student)
}

func (service *StudentServiceImpl) FindAll(ctx context.Context) []web.StudentResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrROllback(tx)

	students := service.studentRepository.FindAll(ctx, tx)
	return helper.ToStudentResponses(students)
}
