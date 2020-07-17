package weather

import "github.com/EricNeid/go-openweather"

// ConditionClear indicates clear weather
const ConditionClear = 8

// Current returns the current weather information for the given city.
func Current(apiKey string, city string) (openweather.CurrentWeather, error) {
	res, err := openweather.NewQueryForCity(apiKey, city).Weather()
	return *res, err
}

// Forecast returns the 5 days weather forecast for the given city.
func Forecast(apiKey string, city string) (openweather.DailyForecast5, error) {
	res, err := openweather.NewQueryForCity(apiKey, city).DailyForecast5()
	return *res, err
}
