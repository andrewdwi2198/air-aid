package models

import "time"

// Data holds the data being sent into the server
type Data struct {
	Timestamp time.Time `json:"time"`
	Longitude float32   `json:"long" sql:"type:decimal(10,3);"`
	Latitude  float32   `json:"lat" sql:"type:decimal(10,3);"`
	CO        float32   `json:"co" sql:"type:decimal(10,3);"`
	NO2       float32   `json:"no2" sql:"type:decimal(10,3);"`
}
