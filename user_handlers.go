package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// Create
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var user Users
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	db, err := dbConnection()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	result := db.Create(&user)
	if result.Error != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
}

// Read
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var user Users
	vars := mux.Vars(r)
	db, err := dbConnection()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	db.First(&user, 10)
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
	var users []Users
	db, err := dbConnection()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	result := db.Raw("SELECT * FROM users WHERE admin = ?", false).Scan(&users)
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
	var update Users
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
	result := db.Model(&update).Where("id = ?", vars["id"]).Updates(Users{Name: update.Name, Email: update.Email, ZipCode: update.ZipCode})
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

// Delete
func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	db, err := dbConnection()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	result := db.Delete(Users{}, "id = ?", vars["ldap"])
	if result.Error != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
}
