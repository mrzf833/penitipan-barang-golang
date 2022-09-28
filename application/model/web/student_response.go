package web

type StudentResponse struct {
	Id          int    `json:"id"`
	Fullname    string `json:"fullname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}
