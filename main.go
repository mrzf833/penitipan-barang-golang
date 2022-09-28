package main

import (
	"net/http"
	"penitipan-barang/application/app"
	"penitipan-barang/application/controller"
	"penitipan-barang/application/helper"
	"penitipan-barang/application/repository"
	"penitipan-barang/application/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	validate := validator.New()
	db := app.NewDB()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	studentRepository := repository.NewStudentRepository()
	studentService := service.NewStudentService(studentRepository, db, validate)
	studentController := controller.NewStudentController(studentService)

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	inventoryRepository := repository.NewInventoryRepository()
	inventoryService := service.NewInventoryService(inventoryRepository, db, validate)
	inventoryController := controller.NewInventoryController(inventoryService)

	authenticationService := service.NewAuthenticationService(userRepository, db, validate)
	authenticationController := controller.NewAuthenticationController(authenticationService)

	router := app.NewRouter(
		db,
		userRepository, categoryController, studentController,
		authenticationController, userController, inventoryController,
	)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
