package controller

import (
	"net/http"
	"penitipan-barang/application/helper"
	"penitipan-barang/application/model/web"
	"penitipan-barang/application/service"

	"github.com/julienschmidt/httprouter"
)

type AuthenticationControllerImpl struct {
	AuthenticationService service.AuthenticationService
}

func NewAuthenticationController(authenticationService service.AuthenticationService) AuthenticationController {
	return &AuthenticationControllerImpl{
		AuthenticationService: authenticationService,
	}
}

func (controller *AuthenticationControllerImpl) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	loginRequest := web.Credentials{}

	helper.ReadFromRequestBody(r, &loginRequest)

	authenticationService := controller.AuthenticationService.Login(r.Context(), loginRequest)
	webResponse := web.WebResponse{
		Message: "Success login",
		Data:    authenticationService,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *AuthenticationControllerImpl) User(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userResponse := controller.AuthenticationService.User(r.Context())
	webResponse := web.WebResponse{
		Message: "success get user",
		Data:    userResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
