package view

import (
	"testing"
	"time"

	"fyne.io/fyne/test"
	"github.com/EricNeid/go-openweather"
	"github.com/EricNeid/go-weatherstation/internal/assets"
	"github.com/EricNeid/go-weatherstation/internal/verify"
)

func TestNewWeather(t *testing.T) {
	// arrange
	window := test.NewApp().NewWindow("TestNewWeather")
	// action
	view, viewModel := NewWeather(nil)
	window.SetContent(view)
	// verify
	verify.NotNil(t, viewModel.city, "city widget not init")
	verify.NotNil(t, viewModel.clock, "clock widget not init")
	verify.NotNil(t, viewModel.lastUpdate, "last update widget not init")
}

func TestWeatherSetTime(t *testing.T) {
	// arrange
	window := test.NewApp().NewWindow("TestWeatherSetTime")
	time := time.Now()
	view, viewModel := NewWeather(nil)
	window.SetContent(view)
	// action
	viewModel.SetTime(time)
	// verify
	verify.Equals(t, time.Format("Mon 15:04"), viewModel.clock.Text)
}

func TestSetCurrentTemperatureData(t *testing.T) {
	// arrange
	window := test.NewApp().NewWindow("TestSetCurrentTemperatureData")
	assets.CurrentLocale = assets.EN
	view, viewModel := NewWeather(nil)
	window.SetContent(view)
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
	viewModel.SetCurrentTemperatureData(testData)
	// verify
	verify.Equals(t, "TestCity", viewModel.city.Text)
	verify.Equals(t, "current temperature: 42.03°", viewModel.currentTemperature.Text)
}
