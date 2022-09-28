package service

import (
	"context"
	"database/sql"
	"fmt"
	"penitipan-barang/application/exception"
	"penitipan-barang/application/helper"
	"penitipan-barang/application/middleware"
	"penitipan-barang/application/model/domain"
	"penitipan-barang/application/model/web"
	"penitipan-barang/application/repository"
	"time"

	"github.com/go-playground/validator/v10"
)

type InventoryServiceImpl struct {
	inventoryRepository repository.InventoryRepository
	DB                  *sql.DB
	Validate            *validator.Validate
}

func NewInventoryService(inventoryRepository repository.InventoryRepository, DB *sql.DB, validate *validator.Validate) InventoryService {
	return &InventoryServiceImpl{
		inventoryRepository: inventoryRepository,
		DB:                  DB,
		Validate:            validate,
	}
}

func (service *InventoryServiceImpl) Create(ctx context.Context, request web.InventoryCreateRequest) web.InventoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrROllback(tx)

	inventory := domain.Inventory{
		CategoryId: sql.NullInt64{
			Int64: int64(request.CategoryId),
		},
		ItemName: request.ItemName,
		Description: sql.NullString{
			String: request.Description,
		},
	}

	if request.DepositStudentId == 0 && request.DepositName == "" {
		panic(exception.NewNotFoundError("deposit_student_id atau deposit_name harus diisi"))
	}

	if request.DepositStudentId != 0 {
		DepositStudent, err := repository.NewStudentRepository().FindById(ctx, tx, request.DepositStudentId)
		fmt.Println("masuk sini")
		if err != nil {
			panic(exception.NewNotFoundError(err.Error()))
		}
		inventory.DepositStudentId = sql.NullInt64{
			Int64: int64(DepositStudent.Id),
		}
		depositName := DepositStudent.Fullname
		inventory.DepositName = sql.NullString{
			String: depositName,
		}
	} else {
		inventory.DepositName = sql.NullString{
			String: request.DepositName,
		}
	}

	inventory.DepositTime = time.Now()
	inventory.DepositAdmin = middleware.AuthenticationGetUser.Id

	inventory = service.inventoryRepository.Save(ctx, tx, inventory)
	return helper.ToInventoryResponse(inventory)
}

func (service *InventoryServiceImpl) Update(ctx context.Context, request web.InventoryUpdateRequest) web.InventoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrROllback(tx)

	studentRepository := repository.NewStudentRepository()

	inventory, err := service.inventoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	inventory.CategoryId = sql.NullInt64{
		Int64: int64(request.CategoryId),
	}
	if request.DepositStudentId != 0 {
		student, err := studentRepository.FindById(ctx, tx, request.DepositStudentId)
		if err != nil {
			panic(exception.NewNotFoundError(err.Error()))
		}
		inventory.DepositStudentId = sql.NullInt64{
			Int64: int64(request.DepositStudentId),
		}
		inventory.DepositName = sql.NullString{
			String: student.Fullname,
		}
	} else {
		inventory.DepositName = sql.NullString{
			String: request.DepositName,
		}
	}
	inventory.DepositTime, err = time.Parse("2006-01-02 15:04:05", request.DepositTime)
	helper.PanicIfError(err)
	inventory.ItemName = request.ItemName
	inventory.Description = sql.NullString{
		String: request.Description,
	}
	inventory.Status = request.Status
	if request.TakeStudentId != 0 {
		student, err := studentRepository.FindById(ctx, tx, request.TakeStudentId)
		if err != nil {
			panic(exception.NewNotFoundError(err.Error()))
		}
		inventory.TakeStudentId = sql.NullInt64{
			Int64: int64(request.TakeStudentId),
		}
		inventory.TakeName = sql.NullString{
			String: student.Fullname,
		}
	} else {
		inventory.TakeName = sql.NullString{
			String: request.TakeName,
		}
	}
	if request.TakeTime != "" {
		takeTime, err := time.Parse("2006-01-02 15:04:05", request.TakeTime)
		helper.PanicIfError(err)
		inventory.TakeTime = sql.NullTime{
			Time: takeTime,
		}
	}

	inventory = service.inventoryRepository.Update(ctx, tx, inventory)

	return helper.ToInventoryResponse(inventory)
}

func (service *InventoryServiceImpl) TakeInventory(ctx context.Context, request web.InventoryTakeRequest) web.InventoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrROllback(tx)

	studentRepository := repository.NewStudentRepository()

	inventory, err := service.inventoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	if inventory.Status == "take" {
		panic(exception.NewNotFoundError("this data already take"))
	}

	inventory.Status = "take"
	if request.TakeStudentId != 0 {
		student, err := studentRepository.FindById(ctx, tx, request.TakeStudentId)
		if err != nil {
			panic(exception.NewNotFoundError(err.Error()))
		}
		inventory.TakeStudentId = sql.NullInt64{
			Int64: int64(request.TakeStudentId),
		}
		inventory.TakeName = sql.NullString{
			String: student.Fullname,
		}
	} else {
		inventory.TakeName = sql.NullString{
			String: request.TakeName,
		}
	}

	takeTime := time.Now()
	inventory.TakeTime = sql.NullTime{
		Time: takeTime,
	}

	inventory = service.inventoryRepository.Update(ctx, tx, inventory)

	return helper.ToInventoryResponse(inventory)
}

func (service *InventoryServiceImpl) Delete(ctx context.Context, inventoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrROllback(tx)

	inventory, err := service.inventoryRepository.FindById(ctx, tx, inventoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.inventoryRepository.Delete(ctx, tx, inventory)
}

func (service *InventoryServiceImpl) FindById(ctx context.Context, inventoryId int) web.InventoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrROllback(tx)

	inventory, err := service.inventoryRepository.FindById(ctx, tx, inventoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToInventoryResponse(inventory)
}

func (service *InventoryServiceImpl) FindAll(ctx context.Context) []web.InventoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrROllback(tx)

	inventories := service.inventoryRepository.FindAll(ctx, tx)
	return helper.ToInventoryResponses(inventories)
}
