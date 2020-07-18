package ui

import (
	"testing"
	"time"

	"fyne.io/fyne/test"
	"github.com/EricNeid/go-openweather"
	"github.com/EricNeid/go-weatherstation/internal/res"
	"github.com/EricNeid/go-weatherstation/internal/verify"
)

func TestNewWeather(t *testing.T) {
	// arrange
	window := test.NewApp().NewWindow("TestNewWeather")
	// action
	unit := NewWeather()
	window.SetContent(unit.View)
	// verify
	verify.NotNil(t, unit.city, "city widget not init")
	verify.NotNil(t, unit.clock, "clock widget not init")
	verify.NotNil(t, unit.lastUpdate, "last update widget not init")
}

func TestWeatherSetTime(t *testing.T) {
	// arrange
	window := test.NewApp().NewWindow("TestWeatherSetTime")
	time := time.Now()
	unit := NewWeather()
	window.SetContent(unit.View)
	// action
	unit.SetTime(time)
	// verify
	verify.Equals(t, time.Format("Mon 15:04"), unit.clock.Text)
}

func TestSetCurrentTemperatureData(t *testing.T) {
	// arrange
	window := test.NewApp().NewWindow("TestSetCurrentTemperatureData")
	res.CurrentLocale = res.EN
	unit := NewWeather()
	window.SetContent(unit.View)
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
