package meteo

import (
	"fmt"
	"github.com/benetis/weather-whisperer/internal/entities"
	"time"
)

type ForecastDto struct {
	ForecastTimeUtc      string  `json:"forecastTimeUtc"`
	AirTemperature       float64 `json:"airTemperature"`
	FeelsLikeTemperature float64 `json:"feelsLikeTemperature"`
	WindSpeed            int     `json:"windSpeed"`
	WindGust             int     `json:"windGust"`
	WindDirection        int     `json:"windDirection"`
	CloudCover           int     `json:"cloudCover"`
	SeaLevelPressure     int     `json:"seaLevelPressure"`
	RelativeHumidity     int     `json:"relativeHumidity"`
	TotalPrecipitation   float64 `json:"totalPrecipitation"`
	ConditionCode        string  `json:"conditionCode"`
}

type PlaceDto struct {
	Code                   string `json:"code"`
	Name                   string `json:"name"`
	AdministrativeDivision string `json:"administrativeDivision"`
	Country                string `json:"country"`
	CountryCode            string `json:"countryCode"`
	Coordinates            struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"coordinates"`
}

type ForecastsResponse struct {
	Place                   PlaceDto      `json:"place"`
	ForecastType            string        `json:"forecastType"`
	ForecastCreationTimeUtc string        `json:"forecastCreationTimeUtc"`
	Forecasts               []ForecastDto `json:"forecastTimestamps"`
}

func convertSingleForecast(forecastDto ForecastDto, placeCode string, forecastCreationTimeUtc string) (entities.Forecast, error) {
	createdFor, err := parseForecastTime(forecastDto.ForecastTimeUtc)
	if err != nil {
		fmt.Println("Error parsing forecast time:", err)
		return entities.Forecast{}, err
	}
	createdAt, err := parseForecastTime(forecastCreationTimeUtc)
	if err != nil {
		fmt.Println("Error parsing forecast creation time:", err)
		return entities.Forecast{}, err
	}

	return entities.Forecast{
		PlaceCode:            entities.PlaceCode(placeCode),
		CreatedFor:           createdFor,
		CreatedAt:            createdAt,
		AirTemperature:       forecastDto.AirTemperature,
		FeelsLikeTemperature: forecastDto.FeelsLikeTemperature,
		WindSpeed:            float64(forecastDto.WindSpeed),
		WindGust:             float64(forecastDto.WindGust),
		WindDirection:        float64(forecastDto.WindDirection),
		CloudCover:           float64(forecastDto.CloudCover),
		SeaLevelPressure:     float64(forecastDto.SeaLevelPressure),
		RelativeHumidity:     float64(forecastDto.RelativeHumidity),
		TotalPrecipitation:   forecastDto.TotalPrecipitation,
		ConditionCode:        forecastDto.ConditionCode,
	}, nil
}

func (fr *ForecastsResponse) ToForecasts() ([]entities.Forecast, error) {
	forecasts := make([]entities.Forecast, len(fr.Forecasts))

	for i, forecastDto := range fr.Forecasts {
		forecast, err := convertSingleForecast(forecastDto, fr.Place.Code, fr.ForecastCreationTimeUtc)
		if err != nil {
			return nil, err
		}
		forecasts[i] = forecast
	}

	return forecasts, nil
}

func parseForecastTime(timeStr string) (time.Time, error) {
	const layout = "2006-01-02 15:04:05"
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func (fr *ForecastsResponse) ToPlace() entities.Place {
	return entities.Place{
		Code:                   entities.PlaceCode(fr.Place.Code),
		Name:                   fr.Place.Name,
		AdministrativeDivision: fr.Place.AdministrativeDivision,
		Country:                fr.Place.Country,
		CountryCode:            fr.Place.CountryCode,
		Latitude:               fr.Place.Coordinates.Latitude,
		Longitude:              fr.Place.Coordinates.Longitude,
	}
}
