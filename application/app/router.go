package app

import (
	"database/sql"
	"penitipan-barang/application/controller"
	"penitipan-barang/application/exception"
	"penitipan-barang/application/middleware"
	"penitipan-barang/application/repository"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(
	db *sql.DB,
	userRepository repository.UserRepository,
	categoryController controller.CategoryController,
	studentController controller.StudentController,
	authenticationController controller.AuthenticationController,
	userController controller.UserController,
	inventoryController controller.InventoryController,
) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/login", authenticationController.Login)
	router.POST("/api/user-login", middleware.AuthMiddleware(authenticationController.User, db, userRepository))
	// router.POST("/api/get-password", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 	password := []byte("superadmin123")
	// 	hashedPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	// 	fmt.Println(string(hashedPassword))
	// })

	router.GET("/api/category", middleware.AuthMiddleware(
		middleware.SuperAdminOrAdminMiddleware(categoryController.FindAll), db, userRepository,
	))
	router.GET("/api/category/:categoryId", middleware.AuthMiddleware(
		middleware.SuperAdminOrAdminMiddleware(categoryController.FindById), db, userRepository,
	))
	router.POST("/api/category", middleware.AuthMiddleware(
		middleware.SuperAdminOrAdminMiddleware(categoryController.Create), db, userRepository,
	))
	router.PUT("/api/category/:categoryId", middleware.AuthMiddleware(
		middleware.SuperAdminOrAdminMiddleware(categoryController.Update), db, userRepository,
	))
	router.DELETE("/api/category/:categoryId", middleware.AuthMiddleware(
		middleware.SuperAdminOrAdminMiddleware(categoryController.Delete), db, userRepository,
	))

	router.GET("/api/student", middleware.AuthMiddleware(
		middleware.SuperAdminOrAdminMiddleware(studentController.FindAll), db, userRepository,
	))
	router.GET("/api/student/:studentId", middleware.AuthMiddleware(
		middleware.SuperAdminOrAdminMiddleware(studentController.FindById), db, userRepository,
	))
	router.POST("/api/student", middleware.AuthMiddleware(
		middleware.SuperAdminOrAdminMiddleware(studentController.Create), db, userRepository,
	))
	router.PUT("/api/student/:studentId", middleware.AuthMiddleware(
		middleware.SuperAdminOrAdminMiddleware(studentController.Update), db, userRepository,
	))
	router.DELETE("/api/student/:studentId", middleware.AuthMiddleware(
		middleware.SuperAdminOrAdminMiddleware(studentController.Delete), db, userRepository,
	))

	router.GET("/api/user", middleware.AuthMiddleware(
		middleware.SuperAdminMiddleware(userController.FindAll), db, userRepository,
	))
	router.GET("/api/user/:userId", middleware.AuthMiddleware(
		middleware.SuperAdminMiddleware(userController.FindById), db, userRepository,
	))
	router.POST("/api/user", middleware.AuthMiddleware(
		middleware.SuperAdminMiddleware(userController.Create), db, userRepository,
	))
	router.PUT("/api/user/:userId", middleware.AuthMiddleware(
		middleware.SuperAdminMiddleware(userController.Update), db, userRepository,
	))
	router.PUT("/api/user/:userId/password", middleware.AuthMiddleware(
		middleware.SuperAdminMiddleware(userController.UpdatePassword), db, userRepository,
	))
	router.DELETE("/api/user/:userId", middleware.AuthMiddleware(
		middleware.SuperAdminMiddleware(userController.Delete), db, userRepository,
	))

	router.GET("/api/inventory/:inventoryId", middleware.AuthMiddleware(
		middleware.SuperAdminOrAdminMiddleware(inventoryController.FindById), db, userRepository,
	))
	router.GET("/api/inventory", middleware.AuthMiddleware(
		middleware.SuperAdminOrAdminMiddleware(inventoryController.FindAll), db, userRepository,
	))
	router.POST("/api/inventory", middleware.AuthMiddleware(
		middleware.SuperAdminOrAdminMiddleware(inventoryController.Create), db, userRepository,
	))
	router.PUT("/api/inventory/:inventoryId", middleware.AuthMiddleware(
		middleware.SuperAdminOrAdminMiddleware(inventoryController.Update), db, userRepository,
	))
	router.PUT("/api/inventory/:inventoryId/take", middleware.AuthMiddleware(
		middleware.SuperAdminOrAdminMiddleware(inventoryController.TakeInventory), db, userRepository,
	))
	router.DELETE("/api/inventory/:inventoryId", middleware.AuthMiddleware(
		middleware.SuperAdminOrAdminMiddleware(inventoryController.Delete), db, userRepository,
	))

	router.PanicHandler = exception.ErrorHandler
	return router
}
