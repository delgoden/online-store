package middleware

import (
	"log"
	"net/http"
)

type CheckAccessFunc func(role string) bool

func CheckAccess(CheckAccessFunc CheckAccessFunc) func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			role, err := Role(request.Context())
			if err != nil {
				log.Println(err)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}

			if !CheckAccessFunc(role) {
				log.Println(err)
				http.Error(writer, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}

			handler.ServeHTTP(writer, request)
		})
	}
}

func CheckAccessAdminOrAbove(role string) (ok bool) {
	if role == "ROOT" || role == "ADMINISTRATOR" {
		ok = true
	}
	return ok
}

func CheckAccessRoot(role string) (ok bool) {

	if role == "ROOT" {
		ok = true
	}
	return ok
}
