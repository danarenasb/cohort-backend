package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var header = r.Header.Get("Authorization")
		header = strings.TrimSpace(header)
		if header == "" || header != "Bearer supersecrettoken" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("Missing auth token")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()
	a := r.PathPrefix("/api").Subrouter()
	u := a.PathPrefix("/users").Subrouter()
	u.HandleFunc("/", getAllUsers).Methods("GET", "OPTIONS")
	u.HandleFunc("/{id}", getUser).Methods("GET", "OPTIONS")
	u.HandleFunc("/{id}", updateUser).Methods("PUT", "OPTIONS")
	r.Use(VerifyToken)
	r.Use(mux.CORSMethodMiddleware(r))

	muxWithMiddlewares := http.TimeoutHandler(r, time.Second*10, "Timeout!")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	http.ListenAndServe(":"+port, muxWithMiddlewares)
}
