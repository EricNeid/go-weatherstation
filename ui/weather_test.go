package ui

import (
	"testing"
	"time"

	verify "github.com/EricNeid/go-weatherstation/internal/test"
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
	unit := NewScreenSaver()
	newTestWindow().SetContent(unit)
	// action
	unit.SetTime(time)
	// verify
	verify.Equals(t, time.Format("Mon 15:04"), unit.clock.Text)
}
