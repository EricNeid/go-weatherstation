package services

import "github.com/EricNeid/go-openweather"

// GetWeather returns the current weather information for the given city.
func GetWeather(apiKey string, city string) (*openweather.CurrentWeather, error) {
	return openweather.NewQueryForCity(apiKey, city).Weather()
}

// GetWeatherForecast returns the 5 days weather forecast for the given city.
func GetWeatherForecast(apiKey string, city string) (*openweather.DailyForecast5, error) {
	return openweather.NewQueryForCity(apiKey, city).DailyForecast5()
}
