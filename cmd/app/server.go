package app

import (
	"errors"
	"net/http"

	"github.com/delgoden/internet-store/cmd/app/middleware"
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

var (
	ErrInternal              = errors.New("internal error")
	ErrLoginUsed             = errors.New("login already registred")
	ErrNoSuchUser            = errors.New("no such user")
	ErrInvalidPassword       = errors.New("invalid password")
	ErrCategoryAlreadyExists = errors.New("category already exists")
	ErrCategoryDoesNotExist  = errors.New("category does not exist")
	ErrProductAlreadyExists  = errors.New("product already exists")
	ErrProductDoesNotExist   = errors.New("product does not exist")
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

	s.mux.Use(middleware.Logger)
	userAuthenticationMd := middleware.Authenticate(s.authSvc.GetIDAndRoleByToken)
	rootAccessMd := middleware.CheckAccess(middleware.CheckAccessRoot)
	adminAccess := middleware.CheckAccess(middleware.CheckAccessAdminOrAbove)

	authSubrouter := s.mux.PathPrefix("/api/auth").Subrouter()
	authSubrouter.HandleFunc("/signup", s.signUp)
	authSubrouter.HandleFunc("/signin", s.signIn)

	rootSubrouter := s.mux.PathPrefix("/api/root").Subrouter()

	rootSubrouter.Use(userAuthenticationMd)
	rootSubrouter.Handle("/role/admin/give/{id:[0-9]+}", rootAccessMd(http.HandlerFunc(s.giveRoleAdministrator))).Methods(POST)
	rootSubrouter.Handle("/role/admin/remove/{id:[0-9]+}", rootAccessMd(http.HandlerFunc(s.removeRoleAdministrator))).Methods(DELETE)

	adminSubrouter := s.mux.PathPrefix("/api/admin").Subrouter()
	adminSubrouter.Use(userAuthenticationMd)
	adminSubrouter.Handle("/category/create", adminAccess(http.HandlerFunc(s.createCategory))).Methods(POST)
	adminSubrouter.Handle("/category/update", adminAccess(http.HandlerFunc(s.updateCategory))).Methods(POST)
	adminSubrouter.Handle("/product/create", adminAccess(http.HandlerFunc(s.createProduct))).Methods(POST)
	adminSubrouter.Handle("/product/update", adminAccess(http.HandlerFunc(s.updateProduct))).Methods(POST)
	adminSubrouter.Handle("/product/remove/{id:[0-9]+}", adminAccess(http.HandlerFunc(s.removeProduct))).Methods(DELETE)
	adminSubrouter.Handle("/product/{id:[0-9]+}/photo/add", adminAccess(http.HandlerFunc(s.addFoto))).Methods(POST)
	adminSubrouter.Handle("/product/photo/{id:[0-9]+}/remove", adminAccess(http.HandlerFunc(s.removeFoto))).Methods(DELETE)

	productSubrouter := s.mux.PathPrefix("/api/product").Subrouter()
	productSubrouter.HandleFunc("/categories", s.getCategories).Methods(GET)
	productSubrouter.HandleFunc("/products", s.getAllActiveProducts).Methods(GET)
	productSubrouter.HandleFunc("/category/{id:[0-9]+}/products", s.getProductsInCategory).Methods(GET)
	productSubrouter.HandleFunc("/product/{id:[0-9]+}", s.getProductByID).Methods(GET)
	s.mux.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("images/product/"))))
}
