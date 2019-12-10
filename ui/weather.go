package ui

import (
	"fmt"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/EricNeid/go-openweather"
	"github.com/EricNeid/go-weatherstation/res"
	"github.com/EricNeid/go-weatherstation/util"
)

var log = util.Log{Context: "weather"}

// Weather represents information view for weather information
type Weather struct {
	widget.Box

	background         *canvas.Image
	city               *widget.Label
	currentTemperature *widget.Label
	clock              *widget.Label
	lastUpdate         *widget.Label

	CloseTouches chan bool
}

type forecast struct {
	header             *widget.Label
	dayTemperature     *widget.Label
	lowestTemperature  *widget.Label
	highestTemperature *widget.Label
	layout             *fyne.Container
}

// NewWeather constructs a new instance of a NewWeather widget.
func NewWeather() *Weather {
	w := &Weather{
		widget.Box{},
		&canvas.Image{FillMode: canvas.ImageFillOriginal},
		widget.NewLabel("City"),
		widget.NewLabel("Current Temperature"),
		widget.NewLabel("Clock"),
		widget.NewLabel("Last update"),
		make(chan bool),
	}
	w.ExtendBaseWidget(w)
	w.city.Alignment = fyne.TextAlignCenter
	w.city.TextStyle.Bold = true
	w.currentTemperature.Alignment = fyne.TextAlignCenter
	w.clock.TextStyle.Bold = true
	w.lastUpdate.Alignment = fyne.TextAlignCenter

	header := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		layout.NewSpacer(),
		fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
			w.city,
			w.currentTemperature,
		),
		widget.NewVBox(),
		layout.NewSpacer(),
	)

	footer := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		widget.NewButton(res.GetLabel("close"), func() {
			w.CloseTouches <- true
		}),
		layout.NewSpacer(),
		w.clock,
	)

	todayForecast := newForecast()
	todayForecast.header.SetText(res.GetLabel("today"))

	tomorrowForecast := newForecast()
	tomorrowForecast.header.SetText(res.GetLabel("tomorrow"))

	afterTomorrowForecast := newForecast()
	afterTomorrowForecast.header.SetText(res.GetLabel("aftertomorrow"))

	center := fyne.NewContainerWithLayout(layout.NewGridLayout(3),
		todayForecast.layout,
		fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
			tomorrowForecast.layout,
			w.lastUpdate,
		),
		afterTomorrowForecast.layout,
	)

	w.Children = []fyne.CanvasObject{
		fyne.NewContainerWithLayout(layout.NewMaxLayout(),
			fyne.NewContainerWithLayout(layout.NewMaxLayout(),
				w.background,
			),
			fyne.NewContainerWithLayout(layout.NewBorderLayout(header, footer, nil, nil), header, footer, center),
		),
	}
	return w
}

func newForecast() forecast {
	forecast := forecast{
		header:             widget.NewLabel("header"),
		dayTemperature:     widget.NewLabel("dayTemperature"),
		lowestTemperature:  widget.NewLabel("lowestTemperature"),
		highestTemperature: widget.NewLabel("highestTemperature"),
	}
	forecast.layout = fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		forecast.header,
		forecast.dayTemperature,
		forecast.lowestTemperature,
		forecast.highestTemperature,
	)
	forecast.header.Alignment = fyne.TextAlignCenter
	forecast.header.TextStyle.Bold = true
	forecast.dayTemperature.Alignment = fyne.TextAlignCenter
	forecast.lowestTemperature.Alignment = fyne.TextAlignCenter
	forecast.highestTemperature.Alignment = fyne.TextAlignCenter

	return forecast
}

// SetBackground changes the background image of the weather screen.
func (weather *Weather) SetBackground(filepath string) error {
	if !util.IsFilePresent(filepath) {
		return fmt.Errorf("Given file %s does not exits", filepath)
	}
	weather.background.File = filepath
	weather.background.Refresh()
	return nil
}

// SetTime sets the time to be displayed.
func (weather *Weather) SetTime(t time.Time) {
	str := t.Format("Mon 15:04")
	weather.clock.SetText(str)
}

// SetCurrentTemperatureData updates header (city and current temperature) with the given information.
func (weather *Weather) SetCurrentTemperatureData(data openweather.CurrentWeather) {
	log.D("SetCurrentTemperatureData", fmt.Sprintf("Received %+v", data))
	weather.city.SetText(data.Name)
	weather.currentTemperature.SetText(fmt.Sprintf(res.GetLabel("currenttemperature"), data.Main.Temp))

	weather.lastUpdate.SetText(fmt.Sprintf(res.GetLabel("lastupdate"), time.Now().Format("Mon 15:04")))
}
