package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"
)

var (
	ErrInternal         = errors.New("internal error")
	ErrTokenExpired     = errors.New("token expired")
	ErrTokenNotFound    = errors.New("token not found")
	ErrNoAuthentication = errors.New("no authentication")
	ErrNoRoles          = errors.New("no roles")

	authenticationIDKey    = &contextKey{"authentication id"}
	authenticationRolesKey = &contextKey{"authentication roles"}
)

type contextKey struct {
	name string
}

func (c *contextKey) String() string {
	return c.name
}

type IDFunc func(ctx context.Context, token string) (int64, string, error)

func Authenticate(idFunc IDFunc) func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			token := request.Header.Get("Authorization")

			id, role, err := idFunc(request.Context(), token)
			if err != nil && id == 0 && role == "" {
				log.Println(1, err)
				http.Error(writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			if err != nil && id != 0 {
				log.Println(2, err)
				http.Error(writer, http.StatusText(http.StatusLocked), http.StatusLocked)
				return
			}

			if err != nil {
				log.Println(3, err)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}

			ctx := context.WithValue(request.Context(), authenticationIDKey, id)
			ctx1 := context.WithValue(ctx, authenticationRolesKey, role)
			request = request.WithContext(ctx1)

			handler.ServeHTTP(writer, request)
		})
	}
}

func Authentication(ctx context.Context) (int64, error) {
	if value, ok := ctx.Value(authenticationIDKey).(int64); ok {
		return value, nil
	}
	return 0, ErrNoAuthentication
}

func Role(ctx context.Context) (string, error) {
	if value, ok := ctx.Value(authenticationRolesKey).(string); ok {
		return value, nil
	}
	return "", ErrNoRoles
}
