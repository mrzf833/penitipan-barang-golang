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

type CategoryServiceImpl struct {
	categoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		categoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrROllback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.categoryRepository.Save(ctx, tx, category)
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrROllback(tx)

	category, err := service.categoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category.Name = request.Name

	category = service.categoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrROllback(tx)

	category, err := service.categoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.categoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrROllback(tx)

	category, err := service.categoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrROllback(tx)

	categories := service.categoryRepository.FindAll(ctx, tx)
	return helper.ToCategoryResponses(categories)
}
