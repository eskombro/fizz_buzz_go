package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler).Methods("GET", "POST")
	router.HandleFunc("/health", healthHandler).Methods("GET", "POST")
	router.HandleFunc("/fizzbuzz", fizzBuzzHandler).Methods("POST")
	router.HandleFunc("/stats", statsHandler).Methods("GET", "POST")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal("Unable to run server")
	}
}
