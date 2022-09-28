package middleware

import (
	"database/sql"
	"net/http"
	"penitipan-barang/application/helper"
	"penitipan-barang/application/model/domain"
	"penitipan-barang/application/model/web"
	"penitipan-barang/application/repository"
	"regexp"

	"github.com/golang-jwt/jwt"
	"github.com/julienschmidt/httprouter"
)

var AuthenticationGetUser web.UserAuthentication

func AuthMiddleware(n httprouter.Handle, DB *sql.DB, userRepository repository.UserRepository) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		regexToken, _ := regexp.Compile(`^Bearer\s(\S+)$`)

		tokenBearer := r.Header.Get("Authorization")
		tokens := regexToken.FindStringSubmatch(tokenBearer)
		if len(tokens) > 1 {
			token := tokens[1]
			tx, err := DB.Begin()
			helper.PanicIfError(err)
			defer helper.CommitOrROllback(tx)

			claims := &domain.Claims{}

			_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
				return domain.JwtKey, nil
			})
			if err != nil {
				helper.WriteToResponseBodyAuthorizationError(w)
				return
			}

			user, err := userRepository.FindByUsername(r.Context(), tx, claims.Username)
			helper.PanicIfError(err)

			AuthenticationGetUser = web.UserAuthentication{
				Id:       user.Id,
				Name:     user.Name,
				Username: user.Username,
				Role:     user.Role,
			}
			n(w, r, ps)
		} else {
			helper.WriteToResponseBodyAuthorizationError(w)
		}
	}
}
