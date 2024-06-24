package meteo

type ForecastsResponse struct {
	Place struct {
		Code                   string `json:"code"`
		Name                   string `json:"name"`
		AdministrativeDivision string `json:"administrativeDivision"`
		Country                string `json:"country"`
		CountryCode            string `json:"countryCode"`
		Coordinates            struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"coordinates"`
	} `json:"place"`
	ForecastType            string `json:"forecastType"`
	ForecastCreationTimeUtc string `json:"forecastCreationTimeUtc"`
	Forecasts               []struct {
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
	} `json:"forecastTimestamps"`
}
