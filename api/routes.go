package api

import (
	"github.com/gorilla/mux"
	_ "net/http"
)

// InitializeRoutes menginisialisasi semua rute API
func InitializeRoutes(router *mux.Router) {
	router.HandleFunc("/items", GetItems).Methods("GET")
	router.HandleFunc("/items/{id}", GetItem).Methods("GET")
	router.HandleFunc("/items", CreateItem).Methods("POST")
	router.HandleFunc("/items/{id}", UpdateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", DeleteItem).Methods("DELETE")
}
