package weather

import "github.com/EricNeid/go-openweather"

// ConditionClear indicates clear weather
const ConditionClear = 8

// Current returns the current weather information for the given city.
func Current(apiKey string, city string) (*openweather.CurrentWeather, error) {
	return openweather.NewQueryForCity(apiKey, city).Weather()
}

// Forecast returns the 5 days weather forecast for the given city.
func Forecast(apiKey string, city string) (*openweather.DailyForecast5, error) {
	return openweather.NewQueryForCity(apiKey, city).DailyForecast5()
}
