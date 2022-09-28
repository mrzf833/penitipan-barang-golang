package web

type InventoryCreateRequest struct {
	CategoryId       int    `json:"category_id" validate:"required,number"`
	DepositStudentId int    `json:"deposit_student_id"`
	DepositName      string `json:"deposit_name"`
	ItemName         string `json:"item_name" validate:"required"`
	Description      string `json:"description" validate:"required"`
}
