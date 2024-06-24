package storage

import (
	"github.com/benetis/weather-whisperer/internal/meteo"
)

func SaveForecasts(fr meteo.ForecastsResponse) error {

	place := fr.ToPlace()
	forecasts, err := fr.ToForecasts()

	if err != nil {
		return err
	}

	db := InitDB()

	db.FirstOrCreate(&place)

	for _, f := range forecasts {
		db.Create(&f)
	}

	return nil
}
