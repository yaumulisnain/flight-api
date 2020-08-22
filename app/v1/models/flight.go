package models

import "time"

type Flight struct {
	ID            int32      `json:"id" gorm:"primary_key;AUTO_INCREMENT;unique_index"`
	FlightNumber  string     `json:"flightNumber" validate:"required" gorm:"unique_index;not_null"`
	DeparturePort string     `json:"departurePort" validate:"required" gorm:"not_null"`
	ArrivalPort   string     `json:"arrivalPort" validate:"required" gorm:"not_null"`
	DepartureTime string     `json:"departureTime" validate:"required" gorm:"not_null"`
	ArrivalTime   string     `json:"arrivalTime" validate:"required" gorm:"not_null"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
	DeletedAt     *time.Time `json:"-"`
}

func (Flight) TableName() string {
	return "flight"
}
