package meteo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchForecasts(placeCode string) (ForecastsResponse, error) {
	url := fmt.Sprintf("https://api.meteo.lt/v1/places/%s/forecasts/long-term", placeCode)

	resp, err := http.Get(url)
	if err != nil {
		return ForecastsResponse{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	var forecasts ForecastsResponse
	err = json.NewDecoder(resp.Body).Decode(&forecasts)
	if err != nil {
		return ForecastsResponse{}, err
	}

	return forecasts, nil
}
