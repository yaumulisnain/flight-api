package repository

import (
	"flight-api/app/v1/models"
)

func FetchFlight(flightCodeParam string) ([]models.Flight, error) {
	db := getDB()

	var Flight []models.Flight

	if err := db.Where("flight_number LIKE ?", flightCodeParam+"%").Find(&Flight).Error; err != nil {
		return Flight, err
	}

	return Flight, nil
}

func FetchFlightByID(Flight *models.Flight, id int32) error {
	db := getDB()
	if err := db.First(Flight, id).Error; err != nil {
		return err
	}
	return nil
}

func CreateFlight(Flight *models.Flight) error {
	db := getDB()

	if err := db.Create(Flight).Error; err != nil {
		return err
	}
	return nil
}

func UpdateFlight(Flight *models.Flight, updateField map[string]interface{}) error {
	db := getDB()

	if err := db.Model(&Flight).Updates(updateField).Error; err != nil {
		return err
	}

	return nil
}

func DeleteFlight(TypeId int32) error {
	var Flight models.Flight

	db := getDB()

	if err := db.First(&Flight, TypeId).Error; err != nil {
		return err
	}

	if err := db.Delete(&Flight).Error; err != nil {
		return err
	}

	return nil
}
