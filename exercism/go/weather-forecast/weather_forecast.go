
// Package weather implements a function Forecast to do weather forecasting.
package weather 


// CurrentCondition indicates the current weather condition.
var CurrentCondition string

// CurrentLocation indicates where the weather forecast has been carried out.
var CurrentLocation string

// Forecast calculates the likelihood of the weather in the coming up time.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
