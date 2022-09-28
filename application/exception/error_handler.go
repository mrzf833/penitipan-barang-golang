package exception

import (
	"fmt"
	"net/http"
	"penitipan-barang/application/helper"
	"penitipan-barang/application/model/web"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err any) {
	if notFoundError(w, r, err) {
		return
	}

	if validationErrors(w, r, err) {
		return
	}

	if authorizationError(w, r, err) {
		return
	}

	fmt.Println(err)
	internalServerError(w, r, err)
}

func internalServerError(w http.ResponseWriter, r *http.Request, err any) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Message: "internal server error",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func notFoundError(w http.ResponseWriter, r *http.Request, err any) bool {
	exception, ok := err.(NotFoundError)

	if ok {
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Message: "not found",
			Data:    exception.Error,
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func authorizationError(w http.ResponseWriter, r *http.Request, err any) bool {
	_, ok := err.(AuthorizationError)

	if ok {
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Message: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func validationErrors(w http.ResponseWriter, r *http.Request, err any) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Message: "Bad Request",
			Data:    exception.Error(),
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}
