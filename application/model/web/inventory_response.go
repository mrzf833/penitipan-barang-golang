package web

type InventoryResponse struct {
	Id               int    `json:"id,omitempty"`
	CategoryId       int    `json:"category_id,omitempty"`
	CategoryName     string `json:"category_name,omitempty"`
	DepositAdmin     int    `json:"deposit_admin,omitempty"`
	DepositStudentId int    `json:"deposit_student_id,omitempty"`
	DepositName      string `json:"deposit_name,omitempty"`
	DepositTime      string `json:"deposit_time,omitempty"`
	ItemName         string `json:"item_name,omitempty"`
	Description      string `json:"description,omitempty"`
	Status           string `json:"status,omitempty"`
	TakeAdmin        int    `json:"take_admin,omitempty"`
	TakeStudentId    int    `json:"take_student_id,omitempty"`
	TakeName         string `json:"take_name,omitempty"`
	TakeTime         string `json:"take_time,omitempty"`
}
