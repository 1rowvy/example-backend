package main

import (
	"back/store"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func main() {

	s := store.NewStore()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /cars", func(w http.ResponseWriter, r *http.Request) {
		// time.Sleep(2 * time.Second)
		writeJSON(w, http.StatusOK, s.GetAll())
	})

	mux.HandleFunc("GET /cars/{id}", func(w http.ResponseWriter, r *http.Request) {
		path := r.PathValue("id")

		id, err := strconv.Atoi(path)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "server error"})
			return
		}

		car, ok := s.GetById(id)
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
	if err := http.ListenAndServe(addr, cors(mux)); err != nil {
		log.Fatal(err)
	}
}
