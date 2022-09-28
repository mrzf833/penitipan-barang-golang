package controller

import (
	"net/http"
	"penitipan-barang/application/helper"
	"penitipan-barang/application/model/web"
	"penitipan-barang/application/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userCreateRequest := web.UserCreateRequest{}

	helper.ReadFromRequestBody(r, &userCreateRequest)

	userResponse := controller.UserService.Create(r.Context(), userCreateRequest)
	webResponse := web.WebResponse{
		Message: "Success create user",
		Data:    userResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *UserControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userId, err := strconv.Atoi(params.ByName("userId"))
	helper.PanicIfError(err)

	userUpdateRequest := web.UserUpdateRequest{}

	helper.ReadFromRequestBody(r, &userUpdateRequest)

	userUpdateRequest.Id = userId

	userResponse := controller.UserService.Update(r.Context(), userUpdateRequest)
	webResponse := web.WebResponse{
		Message: "Success update user",
		Data:    userResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *UserControllerImpl) UpdatePassword(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userId, err := strconv.Atoi(params.ByName("userId"))
	helper.PanicIfError(err)

	userPasswordUpdateRequest := web.UserPasswordUpdateRequest{}

	helper.ReadFromRequestBody(r, &userPasswordUpdateRequest)

	userPasswordUpdateRequest.Id = userId

	controller.UserService.UpdatePassword(r.Context(), userPasswordUpdateRequest)
	webResponse := web.WebResponse{
		Message: "Success update password user",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *UserControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userId, err := strconv.Atoi(params.ByName("userId"))
	helper.PanicIfError(err)

	controller.UserService.Delete(r.Context(), userId)
	webResponse := web.WebResponse{
		Message: "Success delete user",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *UserControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userId, err := strconv.Atoi(params.ByName("userId"))
	helper.PanicIfError(err)

	user := controller.UserService.FindById(r.Context(), userId)
	webResponse := web.WebResponse{
		Message: "Success get user",
		Data:    user,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *UserControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	users := controller.UserService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Message: "Success get all user",
		Data:    users,
	}

	helper.WriteToResponseBody(w, webResponse)
}
