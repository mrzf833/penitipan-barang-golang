package helper

import (
	"penitipan-barang/application/model/domain"
	"penitipan-barang/application/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoriesResponse []web.CategoryResponse
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, ToCategoryResponse(category))
	}
	return categoriesResponse
}

func ToStudentResponse(student domain.Student) web.StudentResponse {
	return web.StudentResponse{
		Id:          student.Id,
		Fullname:    student.Fullname,
		Email:       student.Email,
		PhoneNumber: student.PhoneNumber,
		Address:     student.Address,
	}
}

func ToStudentResponses(students []domain.Student) []web.StudentResponse {
	var studentsResponse []web.StudentResponse
	for _, student := range students {
		studentsResponse = append(studentsResponse, ToStudentResponse(student))
	}
	return studentsResponse
}

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Id:       user.Id,
		Name:     user.Name,
		Username: user.Username,
		Role:     user.Role,
	}
}

func ToUserResponses(users []domain.User) []web.UserResponse {
	var usersResponse []web.UserResponse
	for _, user := range users {
		usersResponse = append(usersResponse, ToUserResponse(user))
	}
	return usersResponse
}

func ToInventoryResponse(inventory domain.Inventory) web.InventoryResponse {
	inventoryResponse := web.InventoryResponse{
		Id:               inventory.Id,
		CategoryId:       int(inventory.CategoryId.Int64),
		CategoryName:     inventory.CategoryName.String,
		DepositAdmin:     inventory.DepositAdmin,
		DepositStudentId: int(inventory.DepositStudentId.Int64),
		DepositName:      inventory.DepositName.String,
		DepositTime:      inventory.DepositTime.String(),
		ItemName:         inventory.ItemName,
		Description:      inventory.DepositName.String,
		Status:           inventory.Status,
		TakeAdmin:        int(inventory.TakeAdmin.Int64),
		TakeStudentId:    int(inventory.TakeStudentId.Int64),
		TakeName:         inventory.TakeName.String,
	}
	if inventory.TakeTime.Valid {
		inventoryResponse.TakeTime = inventory.TakeTime.Time.String()
	}
	return inventoryResponse
}

func ToInventoryResponses(inventories []domain.Inventory) []web.InventoryResponse {
	var inventoriesResponse []web.InventoryResponse
	for _, inventory := range inventories {
		inventoriesResponse = append(inventoriesResponse, ToInventoryResponse(inventory))
	}
	return inventoriesResponse
}
