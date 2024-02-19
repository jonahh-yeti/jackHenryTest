# HTTP WEATHER SERVER

This is a basic http server which connects with the openweathermap API to
return the weather for a given lat/lon, written as a coding test for Jack Henry

## Running instructions
Clone the repo and add a file in the root directory called `openweathermap.key`.  
This should contain a single line with your openweathermap API key, and will be read in on start.

Run the server with `go run main.go` from the root directory.

## Available endpoints

### `/weather`
The weather endpoint accepts the url params `lat` and `lon`, returning a simple json with the temperature
and the condition of the passed in location.  Ex: `{"condition":"Clouds","temperature":"Cold"}`

Example call: `localhost:8080/weather?lat=37.4221&lon=122.0852`
