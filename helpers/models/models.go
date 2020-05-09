package models

import "time"

// Data holds the data being sent into the server
type Data struct {
	Timestamp time.Time `json:"time"`
	Longitude float32   `json:"long"`
	Latitude  float32   `json:"lat"`
	CO        float32   `json:"co"`
	NO2       float32   `json:"no2"`
}
