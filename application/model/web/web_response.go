package web

type WebResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}
