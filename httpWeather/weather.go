package httpWeather

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"math"
	"net/http"
)

// ApiUrl holds a formattable URL to the openweathermap api call which retrieves
// weather data for a given lat/lon
const ApiUrl = "https://api.openweathermap.org/data/2.5/weather?lat=%.6f&lon=%.6f&appid=%s&units=imperial"

// TODO: input this somewhere?

const (
	TooHot  = 75
	TooCold = 40
)

// Results struct is used for parsing out openweathermap api responses
// Note: this does not contain all the data in the json response, see
// https://openweathermap.org/current#parameter for full response contents
type Results struct {
	Coord   map[string]float64       `json:"coord"`
	Weather []map[string]interface{} `json:"weather"`
	Main    map[string]float64       `json:"main"`
}

// GetWeather retrieves the weather conditions for the passed in latitude &
// longitude from the Openweathermap API
func GetWeather(lat, lon float64) (*Results, error) {
	// Validation on lat/lon values
	if math.Abs(lat) > 90.0 {
		return nil, errors.New("latitude must be between -90 and 90")
	} else if math.Abs(lon) > 180.0 {
		return nil, errors.New("longitude must be between -180 and 180")
	}

	// Build request string & get from URL
	reqStr := fmt.Sprintf(ApiUrl, lat, lon, ApiKey)
	weatherResponse, err := http.Get(reqStr)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to GET "+reqStr)
	}

	responseBytes, err := io.ReadAll(weatherResponse.Body)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to read response body")
	}

	// Map to Results struct
	var responseMap = &Results{}
	err = json.Unmarshal(responseBytes, &responseMap)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to unmarshal response")
	}

	return responseMap, nil
}

// ParseWeather accepts a Results object from an openweathermap API query and
// parses it to the minimum useable data, returning a json containing condition
// (raining, snowing, etc.) and temperature (hot, moderate or cold)
func ParseWeather(res *Results) ([]byte, error) {
	var ret = make(map[string]string)
	conditionKey := "condition"
	temperatureKey := "temperature"

	switch temp := res.Main["feels_like"]; {
	case temp > TooHot:
		ret[temperatureKey] = "Hot"
	case temp >= TooCold:
		ret[temperatureKey] = "Moderate"
	case TooCold > temp:
		ret[temperatureKey] = "Cold"
	}

	// 'description' field might also work, would likely be more descriptive
	ret[conditionKey] = res.Weather[0]["main"].(string)

	return json.Marshal(ret)
}
