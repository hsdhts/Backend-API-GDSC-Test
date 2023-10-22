package main

import (
	"backend-gdsc/api"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	// Inisialisasi rute-rute API
	api.InitializeRoutes(router)

	port := 8080
	fmt.Printf("Server berjalan di port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
