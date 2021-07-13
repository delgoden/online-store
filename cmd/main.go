package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/delgoden/internet-store/cmd/app"
	"github.com/delgoden/internet-store/pkg/admin"
	"github.com/delgoden/internet-store/pkg/auth"
	"github.com/delgoden/internet-store/pkg/product"
	"github.com/delgoden/internet-store/pkg/root"
	"github.com/delgoden/internet-store/pkg/user"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/dig"
)

func main() {
	host := "0.0.0.0"
	port := "9999"
	db := "postgres://app:pass@localhost:5432/db"

	if err := excute(host, port, db); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

// excute initializes a connection to the database, initializes services and starts a server instance
func excute(host, port, db string) (err error) {
	deps := []interface{}{
		app.NewServer,
		mux.NewRouter,
		func() (*pgxpool.Pool, error) {
			ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
			return pgxpool.Connect(ctx, db)
		},

		admin.NewService,
		auth.NewService,
		product.NewService,
		root.NewService,
		user.NewService,
		func(server *app.Server) *http.Server {
			return &http.Server{
				Addr:    net.JoinHostPort(host, port),
				Handler: server,
			}
		},
	}

	container := dig.New()
	for _, dep := range deps {
		err = container.Provide(dep)
		if err != nil {
			return err
		}
	}

	err = container.Invoke(func(server *app.Server) {
		server.InitRoute()
	})
	if err != nil {
		return err
	}

	return container.Invoke(func(server *http.Server) error {
		return server.ListenAndServe()
	})
}
