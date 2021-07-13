package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/delgoden/internet-store/cmd/app/middleware"
	"github.com/delgoden/internet-store/pkg/types"
)

func (s *Server) buy(writer http.ResponseWriter, request *http.Request) {
	userID, err := middleware.Authentication(request.Context())
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	positions := &types.Position{}
	err = json.NewDecoder(request.Body).Decode(&positions)
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	purchases, err := s.userSvc.Buy(request.Context(), positions, userID)
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(purchases)
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
