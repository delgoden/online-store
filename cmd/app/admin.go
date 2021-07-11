package app

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/delgoden/internet-store/pkg/types"
)

var (
	ErrCategoryAlreadyExists = errors.New("category already exists")
	ErrCategoryDoesNotExist  = errors.New("category does not exist")
)

// CreateCategory creates a new category
func (s *Server) createCategory(writer http.ResponseWriter, request *http.Request) {
	category := &types.Category{}
	err := json.NewDecoder(request.Body).Decode(&category)
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	category, err = s.adminSvc.CreateCategory(request.Context(), category)
	if category == nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusConflict), http.StatusConflict)
		return
	}
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(category)
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

// UpdateCategory updates an existing category
func (s *Server) updateCategory(writer http.ResponseWriter, request *http.Request) {
	category := &types.Category{}
	err := json.NewDecoder(request.Body).Decode(&category)
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	category, err = s.adminSvc.UpdateCategory(request.Context(), category)
	if err != nil && category.ID == 0 {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(category)
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

// CreateProduct creates a new product
func (s *Server) createProduct(writer http.ResponseWriter, request *http.Request) {

}

// UpdateProduct updates existing products
func (s *Server) updateProduct(writer http.ResponseWriter, request *http.Request) {

}

// RemoveProduct removes product
func (s *Server) removeProduct(writer http.ResponseWriter, request *http.Request) {

}

// AddFoto adds a new photo
func (s *Server) addFoto(writer http.ResponseWriter, request *http.Request) {

}

// RemoveFoto deletes photo
func (s *Server) removeFoto(writer http.ResponseWriter, request *http.Request) {

}
