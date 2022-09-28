package web

type UserResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
