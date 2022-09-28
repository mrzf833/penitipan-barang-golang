package web

type UserCreateRequest struct {
	Name     string `json:"name" validate:"required,max=100,min=1"`
	Username string `json:"username" validate:"required,max=100,min=1"`
	Password string `json:"password" validate:"required,max=100,min=1"`
}
