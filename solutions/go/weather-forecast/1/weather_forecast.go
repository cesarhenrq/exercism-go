/* Package weather provides simple utilities for tracking and reporting
the current weather condition for a given location.*/
package weather

var (
	// CurrentCondition stores the current weather condition (e.g. "sunny", "rainy", "cloudy").
	CurrentCondition string

	// CurrentLocation stores the name of the city or location for which the weather condition applies.
	CurrentLocation string
)

/* Forecast sets the current location and weather condition
   and returns a formatted weather report string. */
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}