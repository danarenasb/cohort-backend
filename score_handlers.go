package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// Create
func createScore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var score Score
	err := json.NewDecoder(r.Body).Decode(&score)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	score.LastUpdated = time.Now()
	db, err := dbConnection()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	result := db.Create(&score)
	if result.Error != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
}

// Read
func getScoresPerUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var score []Score
	vars := mux.Vars(r)
	db, err := dbConnection()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	result := db.Where("student_ldap = ?", strings.ToLower(vars["ldap"])).Find(&score)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			w.WriteHeader(404)
			return
		}
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	err = json.NewEncoder(w).Encode(score)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
}

// Read
func getAllScores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var scores []Score
	db, err := dbConnection()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	result := db.Raw("SELECT * FROM scores").Scan(&scores)
	if result.Error != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	err = json.NewEncoder(w).Encode(scores)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
}

// Update
func updateScore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var update Score
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
	result := db.Model(&update).Where("id = ?", vars["id"]).Updates(Score{Score: update.Score, LastUpdated: time.Now()})
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
func deleteScore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	db, err := dbConnection()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	result := db.Delete(Score{}, "id = ?", strings.ToLower(vars["id"]))
	if result.Error != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
}
