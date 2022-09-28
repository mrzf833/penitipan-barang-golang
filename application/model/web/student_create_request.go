package web

type StudentCreateRequest struct {
	Fullname    string `json:"fullname" validate:"required,max=225,min=1"`
	Email       string `json:"email" validate:"required,max=225,min=1,email"`
	PhoneNumber string `json:"phone_number" validate:"required,max=15,min=1"`
	Address     string `json:"address" validate:"required,max=200,min=1"`
}
