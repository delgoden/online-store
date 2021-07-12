package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/delgoden/internet-store/pkg/types"
)

var ()

// SignUp user registration
func (s *Server) signUp(writer http.ResponseWriter, request *http.Request) {
	var regData *types.Auth
	err := json.NewDecoder(request.Body).Decode(&regData)
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	user, err := s.authSvc.SignUp(request.Context(), regData)
	if user == nil {
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
	writer.WriteHeader(http.StatusCreated)
	_, err = writer.Write(data)
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

// SignIn user authorization
func (s *Server) signIn(writer http.ResponseWriter, request *http.Request) {
	var regData *types.Auth
	err := json.NewDecoder(request.Body).Decode(&regData)
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	user, err := s.authSvc.SignIn(request.Context(), regData)
	if err == ErrNoSuchUser {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusUnauthorized), http.StatusConflict)
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
