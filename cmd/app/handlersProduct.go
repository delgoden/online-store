package app

import (
	"encoding/json"
	"log"
	"net/http"
)

// GetCategories gives a list of existing categories
func (s *Server) getCategories(writer http.ResponseWriter, request *http.Request) {
	categories, err := s.productSvc.GetCategories(request.Context())
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(categories)
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

// GetProducts displays a complete list of products
func (s *Server) getAllProducts(writer http.ResponseWriter, request *http.Request) {

}

// GetProductsInCategory displays a list of products in a category
func (s *Server) getProductsInCategory(writer http.ResponseWriter, request *http.Request) {

}

// GetProductByID issues the product according to its ID
func (s *Server) getProductByID(writer http.ResponseWriter, request *http.Request) {

}
