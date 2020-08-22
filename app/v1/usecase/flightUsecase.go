package usecase

import (
	"flight-api/app/v1/models"
	"flight-api/app/v1/repository"
)

func GetFlightByID(id int32) (models.Flight, error) {
	var (
		Flight models.Flight
		err    error
	)

	if err = repository.FetchFlightByID(&Flight, id); err != nil {
		return Flight, err
	}

	return Flight, nil
}

func GetFlight() ([]models.Flight, error) {
	var (
		Flight []models.Flight
		err    error
	)

	if Flight, err = repository.FetchFlight(); err != nil {
		return Flight, err
	}

	return Flight, nil
}

func CreateFlight(Flight models.Flight) (models.Flight, error) {
	var err error

	if err = repository.CreateFlight(&Flight); err != nil {
		return Flight, err
	}

	return Flight, nil
}

func UpdateFlight(id int32, updateField map[string]interface{}) (models.Flight, error) {
	var Flight models.Flight

	if err := repository.FetchFlightByID(&Flight, id); err != nil {
		return Flight, err
	}

	if err := repository.UpdateFlight(&Flight, updateField); err != nil {
		return Flight, err
	}

	return Flight, nil
}

func DeleteFlight(id int32) error {
	if err := repository.DeleteFlight(id); err != nil {
		return err
	}
	return nil
}
