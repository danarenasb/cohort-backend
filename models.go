package main

import "time"

type Associate struct {
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	LDAP          string    `json:"ldap"`
	Instructor    bool      `json:"instructor"`
	LastUpdatedBy string    `json:"last_updated_by"`
	LastUpdated   time.Time `json:"last_updated"`
}

type TechCheck struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	LastUpdatedBy string    `json:"last_updated_by"`
	LastUpdated   time.Time `json:"last_updated"`
}

type Score struct {
	ID            int       `json:"id"`
	StudentLDAP   string    `json:"ldap"`
	TechCheckID   int       `json:"tech_check"`
	Score         int       `json:"score"`
	GradedBy      string    `json:"graded_by"`
	LastUpdatedBy string    `json:"last_updated_by"`
	LastUpdated   time.Time `json:"last_updated"`
}
