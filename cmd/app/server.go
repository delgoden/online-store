package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"
)

type Server struct {
	mux *mux.Router
}

func NewServer(mux *mux.Router) *Server {
	return &Server{mux: mux}
}

func (s *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.mux.ServeHTTP(writer, request)
}

func (s *Server) InitRoute() {
	authSubrouter := s.mux.PathPrefix("/api/auth").Subrouter()
	authSubrouter.HandleFunc("/signup", s.SignUp)
	authSubrouter.HandleFunc("/signin", s.SignIn)

	rootSubrouter := s.mux.PathPrefix("/api/root").Subrouter()
	rootSubrouter.HandleFunc("/role/admin/give", s.GiveRoleAdministrator).Methods(POST)
	rootSubrouter.HandleFunc("role/admin/remove", s.RemoveRoleAdministrator).Methods(DELETE)

	adminSubrouter := s.mux.PathPrefix("/api/admin").Subrouter()
	adminSubrouter.HandleFunc("/category/create", s.CreateCategory).Methods(POST)
	adminSubrouter.HandleFunc("/category/update", s.UpdateCategory).Methods(POST)
	adminSubrouter.HandleFunc("/product/create", s.CreateProduct).Methods(POST)
	adminSubrouter.HandleFunc("/product/update", s.UpdateProduct).Methods(POST)
	adminSubrouter.HandleFunc("/product/remove", s.RemoveProduct).Methods(DELETE)
	adminSubrouter.HandleFunc("/product/{id:[0-9]+}/foto/add", s.AddFoto).Methods(POST)
	adminSubrouter.HandleFunc("/product/foto/{id:[0-9]+}/remove", s.RemoveFoto).Methods(DELETE)

	productSubrouter := s.mux.PathPrefix("/api/product").Subrouter()
	productSubrouter.HandleFunc("/categories", s.GetCategories).Methods(GET)
	productSubrouter.HandleFunc("/products", s.GetProducts).Methods(GET)
	productSubrouter.HandleFunc("/category/{id:[0-9]+}/products", s.GetProductsInCategory).Methods(GET)
	productSubrouter.HandleFunc("/product/{id:[0-9]+}", s.GetProductByID).Methods(GET)

}
