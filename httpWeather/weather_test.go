package httpWeather

import (
	"testing"
)

func TestGetWeather(t *testing.T) {
	weather, err := GetWeather(42.371920, -71.084880)
	if err != nil {
		t.Fatalf("Failed to get weather: %+v", err)
	}
	t.Logf("Basic weather request success, received: %+v", weather)

	// Test lat/lon bounds
	boundWeather, err := GetWeather(-90, -180)
	if err != nil {
		t.Fatalf("-90/-180 should be within bounds for weather, instead got: %+v", err)
	}
	t.Log(boundWeather)
	_, err = GetWeather(-90.1, -180)
	if err == nil {
		t.Fatal("Expected error did not occur when latitude under -90")
	}
	_, err = GetWeather(-90, -180.1)
	if err == nil {
		t.Fatal("Expected error did not occur when longitude under -180")
	}
}
