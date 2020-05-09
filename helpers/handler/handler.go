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

type Service interface {
	Handler(w http.ResponseWriter, r *http.Request)
}

func NewHandler(db *gorm.DB) Service {
	return &service{
		DB: db,
	}
}

func (svc *service) Handler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data models.Data

	decoder.Decode(&data)
	if data.Latitude != 0 && data.Longitude != 0 {
		svc.DB.Create(&data)
	}

	w.WriteHeader(http.StatusOK)
}
