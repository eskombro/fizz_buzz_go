package server

import (
	"encoding/json"
	"log"
	"net/http"

	fb "fizz_buzz_go/fizzbuzz"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("New HOME request")
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(
		"FIZZBUZZ: Use POST /fizzbuzz to execute, or GET /stats to see stats",
	)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("New HEALTH request")
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode("Running")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func fizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("New FIZZBUZZ request")
	w.Header().Set("Content-Type", "application/json")
	params := fb.FizzBuzzParams{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		log.Println("Error: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("Params:   ", params)
	resp := fb.Fizzbuzz(params)
	log.Println("Response: ", resp)
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("New STATS request")
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode("Some stats")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
