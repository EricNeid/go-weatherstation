package view

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/EricNeid/go-openweather"
	"github.com/EricNeid/go-weatherstation/internal/assets"
	"github.com/EricNeid/go-weatherstation/internal/logger"
	"github.com/EricNeid/go-weatherstation/internal/weather"
)

var log = logger.Log{Context: "weather"}

// Weather represents information view for weather information.
// View property represents actual UI which can be added to a window.
type Weather struct {
	view *fyne.Container

	background         *canvas.Image
	city               *widget.Label
	currentTemperature *widget.Label
	clock              *widget.Label
	lastUpdate         *widget.Label

	today         forecast
	tomorrow      forecast
	afterTomorrow forecast
}

type forecast struct {
	header             *widget.Label
	dayTemperature     *widget.Label
	lowestTemperature  *widget.Label
	maximumTemperature *widget.Label
	icon               *canvas.Image
	layout             *fyne.Container
}

// NewWeather creates a new weather widget with the set tap handler for the close button.
func NewWeather(closeTapped func()) (view fyne.CanvasObject, viewModel *Weather) {
	w := Weather{}
	w.city = widget.NewLabel("City")
	w.city.Alignment = fyne.TextAlignCenter
	w.city.TextStyle.Bold = true

	w.currentTemperature = widget.NewLabel("Current Temperature")
	w.currentTemperature.Alignment = fyne.TextAlignCenter

	w.clock = widget.NewLabel("Clock")
	w.clock.TextStyle.Bold = true

	w.lastUpdate = widget.NewLabel("Last update")
	w.lastUpdate.Alignment = fyne.TextAlignCenter

	w.background = &canvas.Image{FillMode: canvas.ImageFillStretch}
	w.today = newForecast()
	w.tomorrow = newForecast()
	w.afterTomorrow = newForecast()

	header := container.New(layout.NewHBoxLayout(),
		layout.NewSpacer(),
		container.New(layout.NewVBoxLayout(),
			w.city,
			w.currentTemperature,
		),
		container.NewVBox(),
		layout.NewSpacer(),
	)
	footer := container.New(layout.NewHBoxLayout(),
		widget.NewButton(assets.GetLabel(assets.Close), closeTapped),
		layout.NewSpacer(),
		w.clock,
	)
	center := container.New(layout.NewVBoxLayout(),
		container.New(layout.NewGridLayout(3),
			w.today.layout,
			w.tomorrow.layout,
			w.afterTomorrow.layout,
		),
		w.lastUpdate,
	)
	w.view = container.New(layout.NewMaxLayout(),
		w.background,
		container.New(layout.NewVBoxLayout(),
			header,
			layout.NewSpacer(),
			center,
			layout.NewSpacer(),
			footer,
		),
	)

	w.today.header.SetText(assets.GetLabel(assets.Today))
	w.tomorrow.header.SetText(assets.GetLabel(assets.Tomorrow))
	w.afterTomorrow.header.SetText(assets.GetLabel(assets.AfterTomorrow))
	defaultBackground, _ := assets.GetBackgroundImage(weather.ConditionClear)
	w.SetBackground(defaultBackground)

	return w.view, &w
}

func newForecast() forecast {
	f := forecast{
		header:             widget.NewLabel("header"),
		dayTemperature:     widget.NewLabel("daytimetemperature"),
		lowestTemperature:  widget.NewLabel("lowesttemperature"),
		maximumTemperature: widget.NewLabel("maximumtemperature"),
		icon:               &canvas.Image{},
	}
	f.icon.SetMinSize(fyne.NewSize(56, 56))
	f.layout = container.New(layout.NewGridLayout(1),
		f.header,
		f.dayTemperature,
		f.lowestTemperature,
		f.maximumTemperature,
		container.New(layout.NewCenterLayout(),
			f.icon,
		),
	)
	f.header.Alignment = fyne.TextAlignCenter
	f.header.TextStyle.Bold = true
	f.dayTemperature.Alignment = fyne.TextAlignCenter
	f.lowestTemperature.Alignment = fyne.TextAlignCenter
	f.maximumTemperature.Alignment = fyne.TextAlignCenter

	return f
}

// SetBackground changes the background image of the weather screen.
func (weather *Weather) SetBackground(image fyne.Resource) error {
	weather.background.Resource = image
	weather.background.Refresh()
	return nil
}

// SetTime sets the time to be displayed.
func (weather *Weather) SetTime(t time.Time) {
	str := t.Format("Mon 15:04")
	weather.clock.SetText(str)
}

// SetCurrentTemperatureData updates header (city and current temperature) with the given information.
func (weather *Weather) SetCurrentTemperatureData(data *openweather.CurrentWeather) {
	log.D("SetCurrentTemperatureData", fmt.Sprintf("Received %+v", data))

	weather.city.SetText(data.Name)
	weather.currentTemperature.SetText(
		fmt.Sprintf(assets.GetLabel(assets.CurrentTemperature), data.Main.Temp))

	weather.lastUpdate.SetText(
		fmt.Sprintf(assets.GetLabel(assets.LastUpdate), time.Now().Format("Mon 15:04")))
}

// SetForecastTemperatureData updates the forecast displayed with the given information.
func (weather *Weather) SetForecastTemperatureData(data *openweather.DailyForecast5) {
	log.D("SetForecastTemperatureData", fmt.Sprintf("Received %+v", data))

	condition := data.List[0].Weather[0].ID
	image, err := assets.GetBackgroundImage(condition)
	if err != nil {
		log.E("SetForecastTemperatureData", err)
	} else {
		weather.SetBackground(image)
	}

	weather.today.updateInformation(
		data.List[0].Temp.Day,
		data.List[0].Temp.Min,
		data.List[0].Temp.Max,
		data.List[0].Weather[0].Icon)
	weather.tomorrow.updateInformation(
		data.List[1].Temp.Day,
		data.List[1].Temp.Min,
		data.List[1].Temp.Max,
		data.List[1].Weather[0].Icon)
	weather.afterTomorrow.updateInformation(
		data.List[2].Temp.Day,
		data.List[2].Temp.Min,
		data.List[2].Temp.Max,
		data.List[2].Weather[0].Icon)
}

func (forecast *forecast) updateInformation(
	dayTimeTemperatue float64,
	minTemperature float64,
	maxTemperature float64,
	conditionIcon string,
) {
	forecast.dayTemperature.SetText(
		fmt.Sprintf(assets.GetLabel(assets.DayTimeTemperature), dayTimeTemperatue))
	forecast.lowestTemperature.SetText(
		fmt.Sprintf(assets.GetLabel(assets.MinTemperature), minTemperature))
	forecast.maximumTemperature.SetText(
		fmt.Sprintf(assets.GetLabel(assets.MaxTemperature), maxTemperature))

	res, err := assets.GetConditionIcon(conditionIcon)
	if err != nil {
		log.E("updateInformation", err)
	} else {
		forecast.icon.Resource = res
		forecast.icon.Refresh()
	}
}

// Hide makes the ui invisible
func (weather *Weather) Hide() {
	if !weather.view.Hidden {
		weather.view.Hide()
	}
}

// Show makes the ui visible
func (weather *Weather) Show() {
	if weather.view.Hidden {
		weather.view.Show()
	}
}
