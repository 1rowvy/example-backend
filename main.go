package main

import (
	"back/store"
	"encoding/json"
	"log"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func main() {

	s := store.NewStore()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /cars", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, s.GetAll())
	})

	mux.HandleFunc("GET /cars/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")

		car, ok := s.GetByName(name)
		if !ok {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "not found"})
			return
		}

		writeJSON(w, http.StatusOK, car)
	})

	mux.HandleFunc("POST /cars", func(w http.ResponseWriter, r *http.Request) {
		var car store.Car

		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		if err := dec.Decode(&car); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid json"})
			return
		}

		createdCar := s.Add(car)
		writeJSON(w, http.StatusCreated, createdCar)
	})

	addr := ":8080"
	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
