package main

import (
	"jackHenryTest/httpWeather"
	"log"
	"os"
	"path"
)

// KeyFile denotes the expected file name for the weather api key
const KeyFile = "openweathermap.key"

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working dir: %+v", err)
	}
	keyPath := path.Join(dir, KeyFile)
	httpWeather.Run(keyPath)
}
