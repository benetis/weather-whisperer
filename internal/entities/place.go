package entities

type PlaceCode string

type Place struct {
	Code                   PlaceCode `gorm:"primaryKey"`
	Name                   string
	AdministrativeDivision string
	Country                string
	CountryCode            string
	Latitude               float64
	Longitude              float64
	Forecasts              []Forecast `gorm:"foreignKey:PlaceCode"`
}
