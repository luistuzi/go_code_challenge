package api

import (
	"database/sql"
	"go_code_challenge/services/device"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1/go_code_challenge").Subrouter()

	deviceController := device.NewController()
	deviceController.RegisterRoutes(subrouter)

	log.Println("Starting on endppoint: ", s.addr)

	return http.ListenAndServe(s.addr, router)
}
