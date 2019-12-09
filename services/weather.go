package services

import "github.com/EricNeid/go-openweather"

func getWeather(apiKey string, city string) (*openweather.CurrentWeather, error) {
	return openweather.NewQueryForCity(apiKey, city).Weather()
}

func getWeatherForecast(apiKey string, city string) (*openweather.DailyForecast5, error) {
	return openweather.NewQueryForCity(apiKey, city).DailyForecast5()
}
