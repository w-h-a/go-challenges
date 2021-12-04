// Package weather provides forecasting tools.
package weather

// CurrentCondition represents the current weather conditions as a string.
var CurrentCondition string
// CurrentLocation represents the current location as a string.
var CurrentLocation string

// Forecast returns a string representing the current weather condition
// in the current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
