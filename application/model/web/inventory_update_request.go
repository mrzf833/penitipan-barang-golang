package web

type InventoryUpdateRequest struct {
	Id               int    `json:"id"`
	CategoryId       int    `json:"category_id" validate:"required,number"`
	DepositStudentId int    `json:"deposit_student_id"`
	DepositName      string `json:"deposit_name"`
	DepositTime      string `json:"deposit_time"`
	ItemName         string `json:"item_name" validate:"required"`
	Description      string `json:"description"`
	Status           string `json:"status"`
	TakeStudentId    int    `json:"take_student_id"`
	TakeName         string `json:"take_name"`
	TakeTime         string `json:"take_time"`
}
