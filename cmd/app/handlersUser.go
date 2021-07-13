package app

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/delgoden/internet-store/cmd/app/middleware"
	"github.com/delgoden/internet-store/pkg/types"
	"github.com/gorilla/mux"
)

// buy buy the selected item
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

// addProductIntoCart add item to cart
func (s *Server) addProductIntoCart(writer http.ResponseWriter, request *http.Request) {
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

	purchases, err := s.userSvc.AddProductIntoCart(request.Context(), positions, userID)
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

// removeProductFromCart remove item from cart
func (s *Server) removeProductFromCart(writer http.ResponseWriter, request *http.Request) {
	userID, err := middleware.Authentication(request.Context())
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	vars := mux.Vars(request)
	idParam := vars["id"]

	productID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	status, err := s.userSvc.RemoveProductFromCart(request.Context(), productID, userID)
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(status)
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

// buyFromCart buy goods from the cart
func (s *Server) buyFromCart(writer http.ResponseWriter, request *http.Request) {
	userID, err := middleware.Authentication(request.Context())
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	purchases, err := s.userSvc.BuyFromCart(request.Context(), userID)
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
