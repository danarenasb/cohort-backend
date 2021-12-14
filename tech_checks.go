package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// Create
func createTechCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tc TechCheck
	err := json.NewDecoder(r.Body).Decode(&tc)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	tc.LastUpdated = time.Now()
	db, err := dbConnection()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	result := db.Create(&tc)
	if result.Error != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
}

// Read
func getAllTC(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tc []TechCheck
	db, err := dbConnection()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	result := db.Raw("SELECT * FROM tech_checks").Scan(&tc)
	if result.Error != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	err = json.NewEncoder(w).Encode(tc)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
}

// Update
func updateTC(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var update TechCheck
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
	result := db.Model(&update).Where("id = ?", vars["id"]).Updates(TechCheck{Name: update.Name, LastUpdatedBy: update.LastUpdatedBy, LastUpdated: time.Now()})
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
func deleteTC(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	db, err := dbConnection()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	result := db.Delete(TechCheck{}, "id = ?", strings.ToLower(vars["id"]))
	if result.Error != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
}
