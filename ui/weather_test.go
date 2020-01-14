package ui

import (
	"testing"
	"time"

	"EricNeid/go-weatherstation/internal/verify"
	"EricNeid/go-weatherstation/res"
	"github.com/EricNeid/go-openweather"
)

func TestNewWeather(t *testing.T) {
	// arrange
	testWindow := newTestWindow()
	// action
	unit := NewWeather()
	testWindow.SetContent(unit)
	// verify
	verify.Equals(t, 1, len(unit.Children))
}

func TestWeatherSetTime(t *testing.T) {
	// arrange
	time := time.Now()
	unit := NewWeather()
	newTestWindow().SetContent(unit)
	// action
	unit.SetTime(time)
	// verify
	verify.Equals(t, time.Format("Mon 15:04"), unit.clock.Text)
}

func TestSetCurrentTemperatureData(t *testing.T) {
	// arrange
	res.CurrentLocale = res.EN
	unit := NewWeather()
	newTestWindow().SetContent(unit)
	testData := openweather.CurrentWeather{
		Name: "TestCity",
		Main: struct {
			Temp     float64 "json:\"temp\""
			Pressure int     "json:\"pressure\""
			Humidity int     "json:\"humidity\""
			TempMin  float64 "json:\"temp_min\""
			TempMax  float64 "json:\"temp_max\""
		}{
			Temp: 42.034,
		},
	}
	// action
	unit.SetCurrentTemperatureData(testData)
	// verify
	verify.Equals(t, "TestCity", unit.city.Text)
	verify.Equals(t, "current temperature: 42.03°", unit.currentTemperature.Text)
}
