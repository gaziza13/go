// main.go (inside cmd/mygoapp/)
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gaziza13/go/tsis1/api"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/music", api.GetAllMusics).Methods("GET")
	r.HandleFunc("/music/{id}", api.GetMusic).Methods("GET")
	r.HandleFunc("/health", api.HealthCheck).Methods("GET")

	log.Fatal(http.ListenAndServe(":80", r))
}
