package app

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/delgoden/internet-store/pkg/types"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var (
	ErrCategoryAlreadyExists = errors.New("category already exists")
	ErrCategoryDoesNotExist  = errors.New("category does not exist")
	ErrProductAlreadyExists  = errors.New("product already exists")
	ErrProductDoesNotExist   = errors.New("product does not exist")
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
	product := &types.Product{}
	err := json.NewDecoder(request.Body).Decode(&product)
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	product, err = s.adminSvc.CreateProduct(request.Context(), product)
	if product == nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusConflict), http.StatusConflict)
		return
	}

	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(product)
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

// UpdateProduct updates existing products
func (s *Server) updateProduct(writer http.ResponseWriter, request *http.Request) {
	product := &types.Product{}
	err := json.NewDecoder(request.Body).Decode(&product)
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	status, err := s.adminSvc.UpdateProduct(request.Context(), product)
	if status.Status == false {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

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

// RemoveProduct removes product
func (s *Server) removeProduct(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	idParam := vars["id"]

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	status, err := s.adminSvc.RemoveProduct(request.Context(), id)
	if err == ErrCategoryDoesNotExist && status.Status == false {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
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

// AddFoto adds a new photo
func (s *Server) addFoto(writer http.ResponseWriter, request *http.Request) {
	foto := &types.Photo{}
	vars := mux.Vars(request)
	idParam := vars["id"]

	product_id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	file, fileHead, err := request.FormFile("image")
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	foto.File = file
	fileName := fileHead.Filename
	nameSls := strings.Split(fileName, ".")
	foto.Name = uuid.New().String() + "." + nameSls[1]

	status, err := s.adminSvc.AddPhoto(request.Context(), foto, product_id)
	if err != nil {
		log.Print(err)
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

// RemoveFoto deletes photo
func (s *Server) removeFoto(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	idParam := vars["id"]
	photoID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	status, err := s.adminSvc.RemovePhoto(request.Context(), photoID)
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
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
