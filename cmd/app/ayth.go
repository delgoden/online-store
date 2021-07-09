package app

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/delgoden/internet-store/pkg/types"
)

var (
	ErrInternal  = errors.New("internal error")
	ErrLoginUsed = errors.New("login already registred")
)

// SignUp user registration
func (s *Server) SignUp(writer http.ResponseWriter, request *http.Request) {
	var regData *types.Auth
	err := json.NewDecoder(request.Body).Decode(&regData)
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	user, err := s.authSvc.SignUp(request.Context(), regData)
	if err == ErrLoginUsed {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusConflict), http.StatusConflict)
		return
	}
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(data)
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

// SignIn user authorization
func (s *Server) SignIn(writer http.ResponseWriter, request *http.Request) {

}
