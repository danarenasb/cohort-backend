package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// Read
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var user User
	vars := mux.Vars(r)
	db, err := dbConnection()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	result := db.First(&user, vars["id"])
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			w.WriteHeader(404)
			return
		}
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
}

// Read
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var users []User
	db, err := dbConnection()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	result := db.Find(&users).Scan(&users)
	if result.Error != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
}

// Update
func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var update User
	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	vars := mux.Vars(r)
	db, err := dbConnection()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	result := db.Model(&update).Where("id = ?", vars["id"]).Updates(User{Name: update.Name, Email: update.Email, ZipCode: update.ZipCode})
	if result.Error != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	err = json.NewEncoder(w).Encode("Record updated")
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
}
