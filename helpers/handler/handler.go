package handler

import (
	"encoding/json"
	"net/http"

	"github.com/andrewdwi2198/air-aid/helpers/models"
	"github.com/jinzhu/gorm"
)

type service struct {
	DB *gorm.DB
}

// Service interface contains the functions of the Service
type Service interface {
	MutateHandler(w http.ResponseWriter, r *http.Request)
	RequestHandler(w http.ResponseWriter, r *http.Request)
}

// NewHandler will pass the handler service implementation of Service interface
func NewHandler(db *gorm.DB) Service {
	return &service{
		DB: db,
	}
}

// MutateHandler inserts data into the DB
func (svc *service) MutateHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data models.Data

	decoder.Decode(&data)
	if data.Latitude != 0 && data.Longitude != 0 {
		svc.DB.Create(&data)
	}

	w.WriteHeader(http.StatusOK)
}

// RequestHandler retrieves data from the DB
func (svc *service) RequestHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: add retrieve data according to criteria
	w.WriteHeader(http.StatusOK)
}
