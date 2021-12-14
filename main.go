package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("we are up and running...")
	r := mux.NewRouter()

	u := r.PathPrefix("/users").Subrouter()
	u.HandleFunc("/", createUser).Methods("POST", "OPTIONS")
	u.HandleFunc("/", getAllUsers).Methods("GET", "OPTIONS")
	u.HandleFunc("/{ldap}", getUser).Methods("GET", "POST", "OPTIONS")
	u.HandleFunc("/{ldap}", deleteUser).Methods("DELETE", "OPTIONS")
	u.HandleFunc("/{ldap}", updateUser).Methods("PUT", "OPTIONS")

	s := r.PathPrefix("/scores").Subrouter()
	s.HandleFunc("/", createScore).Methods("POST")
	s.HandleFunc("/", getAllScores).Methods("GET")
	s.HandleFunc("/{ldap}", getScoresPerUser).Methods("GET", "POST")
	s.HandleFunc("/{id}", deleteScore).Methods("DELETE")
	s.HandleFunc("/{id}", updateScore).Methods("PUT")

	t := r.PathPrefix("/techchecks").Subrouter()
	t.HandleFunc("/", createTechCheck).Methods("POST")
	t.HandleFunc("/", getAllTC).Methods("GET")
	t.HandleFunc("/{id}", deleteTC).Methods("DELETE")
	t.HandleFunc("/{id}", updateTC).Methods("PUT")
	r.Use(mux.CORSMethodMiddleware(r))

	muxWithMiddlewares := http.TimeoutHandler(r, time.Second*10, "Timeout!")

	http.ListenAndServe(":8080", muxWithMiddlewares)
}
