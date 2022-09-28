package middleware

import (
	"net/http"
	"penitipan-barang/application/helper"

	"github.com/julienschmidt/httprouter"
)

func SuperAdminOrAdminMiddleware(n httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		role := AuthenticationGetUser.Role
		if role == "admin" || role == "super_admin" {
			n(w, r, ps)
		} else {
			helper.WriteToResponseBodyForbiddenError(w)
		}
	}
}
