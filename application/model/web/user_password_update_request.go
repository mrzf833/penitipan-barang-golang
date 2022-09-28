package web

type UserPasswordUpdateRequest struct {
	Id       int    `validate:"required" json:"id"`
	Password string `validate:"required,max=100,min=1" json:"password"`
}
