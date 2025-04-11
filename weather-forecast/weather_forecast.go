// Package weather provides tools to forecast weather in Goblinocus.
package weather

// CurrentCondition represents the current weather condition in a given location.
var CurrentCondition string
// CurrentLocation represents the current location for which we're forecasting the weather.
var CurrentLocation string

// Forecast takes a city (string) and condition (string) as input values and returns a string detailing the current weather condition at the current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
