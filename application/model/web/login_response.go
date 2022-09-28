package web

type LoginResponse struct {
	Id       int    `josn:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}
