package httpWeather

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
)

// ApiKey is set to the openweathermap api key when Run is called
var ApiKey = ""

// Run a bare-bones http server with a single endpoint /weather.
// Accepts path to openweathermap API key
func Run(keyPath string) {
	// read and store key for use in api calls
	ApiKey = readKey(keyPath)

	// Assign handleWeather to /weather endpoint
	http.HandleFunc("/weather", handleWeather)

	// Serve until killed
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// handleWeather is the http handler for the /weather endpoint.
// It parses the requested lat/lon, requests weather data for that
// location, and parses it to a basic json conveying temperature and condition.
func handleWeather(w http.ResponseWriter, r *http.Request) {
	// Parse out url params
	lat, err := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse latitude param: %+v", err), 400)
		return
	}
	lon, err := strconv.ParseFloat(r.URL.Query().Get("lon"), 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse longitude param: %+v", err), 400)
		return
	}

	// Get weather info from api
	weatherInfo, err := GetWeather(lat, lon)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get weather from api: %+v", err), 500)
		return
	}
	// Parse weather info to bare-bones response
	response, err := ParseWeather(weatherInfo)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse weather response from api: %+v", err), 500)
		return
	}

	// Write response to http.ResponseWriter
	_, err = w.Write(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to write response to http.ResponseWriter: %+v", err), 500)
		return
	}
}

// Read api key from file and return as a string
func readKey(filePath string) string {
	// Clean file path & read key
	pathClean := path.Clean(filePath)
	keyBytes, err := os.ReadFile(pathClean)
	if err != nil {
		log.Fatalf("Failed to read OpenWeatherMap api key at path %s", pathClean)
	}
	// Remove whitespace/newlines from key string
	return strings.Trim(string(keyBytes), " \n")
}
