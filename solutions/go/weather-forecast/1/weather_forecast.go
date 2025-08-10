// Package weather provides a reading of the current weather conditions.
package weather

// CurrentCondition describes the current weather condition.
var CurrentCondition string

// CurrentLocation describes the current location that the weather is being reported on.
var CurrentLocation string

// Forecast This function outputs the current weather conditions for the given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
