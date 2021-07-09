package app

import (
	"net/http"

	"github.com/delgoden/internet-store/pkg/admin"
	"github.com/delgoden/internet-store/pkg/auth"
	"github.com/delgoden/internet-store/pkg/product"
	"github.com/delgoden/internet-store/pkg/root"
	"github.com/gorilla/mux"
)

const (
	GET    = "GET"    //
	POST   = "POST"   //
	DELETE = "DELETE" //
)

// Server ...
type Server struct {
	mux        *mux.Router
	adminSvc   *admin.Service
	authSvc    *auth.Service
	productSvc *product.Service
	rootSvc    *root.Service
}

// NewServer constructor function to create the server
func NewServer(mux *mux.Router, adminSvc *admin.Service, authSvc *auth.Service, productSvc *product.Service, rootSvc *root.Service) *Server {
	return &Server{
		mux:        mux,
		adminSvc:   adminSvc,
		authSvc:    authSvc,
		productSvc: productSvc,
		rootSvc:    rootSvc,
	}
}

func (s *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.mux.ServeHTTP(writer, request)
}

// InitRoute registration of routers
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
