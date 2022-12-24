package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("we are up and running...")
	r := mux.NewRouter()

	u := r.PathPrefix("/users").Subrouter()
	u.HandleFunc("/", createUser).Methods("POST", "OPTIONS")
	u.HandleFunc("/", getAllUsers).Methods("GET", "OPTIONS")
	u.HandleFunc("/{id}", getUser).Methods("GET", "POST", "OPTIONS")
	u.HandleFunc("/{id}", deleteUser).Methods("DELETE", "OPTIONS")
	u.HandleFunc("/{id}", updateUser).Methods("PUT", "OPTIONS")

	r.Use(mux.CORSMethodMiddleware(r))

	muxWithMiddlewares := http.TimeoutHandler(r, time.Second*10, "Timeout!")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, muxWithMiddlewares)
}
