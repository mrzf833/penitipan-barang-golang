package web

type InventoryTakeRequest struct {
	Id            int    `json:"id"`
	TakeStudentId int    `json:"take_student_id"`
	TakeName      string `json:"take_name"`
}
