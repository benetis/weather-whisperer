package main

import "github.com/benetis/weather-whisperer/internal/meteo"

func main() {

	forecasts, err := meteo.FetchForecasts("kaunas")

	if err != nil {
		panic(err)
	}

	for _, forecast := range forecasts.Forecasts {
		println(forecast.ForecastTimeUtc)
	}

}
