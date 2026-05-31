// Package weather is use for forecast the current weather condition of various cities in Goblinocus.
package weather

var (
    // CurrentCondition variable tell what is the current weather condition.
	CurrentCondition string
    // CurrentLocation variable tell what is the current location.
	CurrentLocation  string
)
// Forecast function take city and conditon and tell us what is the weather condition of that city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
