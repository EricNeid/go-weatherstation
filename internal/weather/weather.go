package weather

import "github.com/EricNeid/go-openweather"

// ConditionClear indicates clear weather
const ConditionClear = 8

// Current returns the current weather information for the given city.
func Current(apiKey, city string) (openweather.CurrentWeather, error) {
	result, err := openweather.NewQueryForCity(apiKey, city).Weather()
	return *result, err
}

// Forecast returns the 5 days weather forecast for the given city.
func Forecast(apiKey, city string) (openweather.DailyForecast5, error) {
	result, err := openweather.NewQueryForCity(apiKey, city).DailyForecast5()
	return *result, err
}
