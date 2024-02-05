// Package weather contains the forecasting package of the current weather condition of various cities in Goblinocus.
package weather

// CurrentCondition represents the current weather condition.
var CurrentCondition string

// CurrentLocation represents the current corresponding location with the current weather condition.
var CurrentLocation string

// Forecast returns a string contains the current location correspond with the current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
