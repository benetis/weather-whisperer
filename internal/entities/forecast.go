package entities

import (
	"time"
)

type Forecast struct {
	Id                   uint      `gorm:"primaryKey"`
	PlaceCode            PlaceCode `gorm:"index"`
	CreatedFor           time.Time `gorm:"index"`
	CreatedAt            time.Time // ForecastCreationTimeUtc
	AirTemperature       float64
	FeelsLikeTemperature float64
	WindSpeed            float64
	WindGust             float64
	WindDirection        float64
	CloudCover           float64
	SeaLevelPressure     float64
	RelativeHumidity     float64
	TotalPrecipitation   float64
	ConditionCode        string
}
