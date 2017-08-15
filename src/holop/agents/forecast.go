package agents

import (
	"net/http"
	"fmt"
	"encoding/json"
)

type GeoPoint struct {
	Title string
	Latitude float64
	Longitude float64
}

type WindForecast struct {
	Place   GeoPoint
	Time    string
	WindSpeed float64
}

type RainProbabilityTO struct {
	Probability float64
}

type ForecastTO struct {
	Items []WindForecast
}

func GetCurrentWind() (f *ForecastTO, e error) {
	var forecast = ForecastTO{}

	res, err := http.Get("http://localhost:8081/wind/stats/current" )
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&forecast)

	return &forecast, nil
}

func GetRainProbability() (f *RainProbabilityTO, e error) {
	var probability = RainProbabilityTO{}

	res, err := http.Get("http://localhost:8081/rain/stats/day" )
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&probability)

	return &probability, nil
}