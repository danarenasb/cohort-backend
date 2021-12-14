package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// Create
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var associate Associate
	err := json.NewDecoder(r.Body).Decode(&associate)
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
	result := db.Create(&associate)
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
	var associate Associate
	vars := mux.Vars(r)
	db, err := dbConnection()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	result := db.Where("ldap = ?", strings.ToLower(vars["ldap"])).First(&associate)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			w.WriteHeader(404)
			return
		}
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	err = json.NewEncoder(w).Encode(associate)
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
	var associates []Associate
	db, err := dbConnection()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	result := db.Raw("SELECT * FROM associates WHERE instructor = ?", false).Scan(&associates)
	if result.Error != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	err = json.NewEncoder(w).Encode(associates)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
}

// Update
func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var update Associate
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
	result := db.Model(&update).Where("LDAP = ?", vars["ldap"]).Updates(Associate{Name: update.Name, Email: update.Email, LDAP: update.LDAP})
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
	result := db.Delete(Associate{}, "ldap = ?", strings.ToLower(vars["ldap"]))
	if result.Error != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
}
