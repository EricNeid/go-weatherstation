package ui

import (
	"fmt"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/EricNeid/go-openweather"
	"github.com/EricNeid/go-weatherstation/assets"
	"github.com/EricNeid/go-weatherstation/weather"
)

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
func (w *Weather) SetBackground(image fyne.Resource) error {
	w.background.Resource = image
	w.background.Refresh()
	return nil
}

// SetTime sets the time to be displayed.
func (w *Weather) SetTime(t time.Time) {
	str := t.Format("Mon 15:04")
	w.clock.SetText(str)
}

// SetCurrentTemperatureData updates header (city and current temperature) with the given information.
func (w *Weather) SetCurrentTemperatureData(data *openweather.CurrentWeather) {
	log.Println("weather", "SetCurrentTemperatureData", fmt.Sprintf("%+v\n", data))

	w.city.SetText(data.Name)
	w.currentTemperature.SetText(
		fmt.Sprintf(assets.GetLabel(assets.CurrentTemperature), data.Main.Temp))

	w.lastUpdate.SetText(
		fmt.Sprintf(assets.GetLabel(assets.LastUpdate), time.Now().Format("Mon 15:04")))
}

// SetForecastTemperatureData updates the forecast displayed with the given information.
func (w *Weather) SetForecastTemperatureData(data *openweather.DailyForecast5) {
	log.Println("weather", "SetForecastTemperatureData", fmt.Sprintf("%+v\n", data))

	condition := data.List[0].Weather[0].ID
	image, err := assets.GetBackgroundImage(condition)
	if err != nil {
		log.Panicln("weather", "SetForecastTemperatureData", "error while getting background image", err)
	} else {
		w.SetBackground(image)
	}

	w.today.updateInformation(
		data.List[0].Temp.Day,
		data.List[0].Temp.Min,
		data.List[0].Temp.Max,
		data.List[0].Weather[0].Icon)
	w.tomorrow.updateInformation(
		data.List[1].Temp.Day,
		data.List[1].Temp.Min,
		data.List[1].Temp.Max,
		data.List[1].Weather[0].Icon)
	w.afterTomorrow.updateInformation(
		data.List[2].Temp.Day,
		data.List[2].Temp.Min,
		data.List[2].Temp.Max,
		data.List[2].Weather[0].Icon)
}

func (forecast *forecast) updateInformation(
	dayTimeTemperature float64,
	minTemperature float64,
	maxTemperature float64,
	conditionIcon string,
) {
	forecast.dayTemperature.SetText(
		fmt.Sprintf(assets.GetLabel(assets.DayTimeTemperature), dayTimeTemperature))
	forecast.lowestTemperature.SetText(
		fmt.Sprintf(assets.GetLabel(assets.MinTemperature), minTemperature))
	forecast.maximumTemperature.SetText(
		fmt.Sprintf(assets.GetLabel(assets.MaxTemperature), maxTemperature))

	res, err := assets.GetConditionIcon(conditionIcon)
	if err != nil {
		log.Panicln("weather", "SetForecastTemperatureData", "error while getting condition icon", err)
	} else {
		forecast.icon.Resource = res
		forecast.icon.Refresh()
	}
}

// Hide makes the ui invisible
func (w *Weather) Hide() {
	if !w.view.Hidden {
		w.view.Hide()
	}
}

// Show makes the ui visible
func (w *Weather) Show() {
	if w.view.Hidden {
		w.view.Show()
	}
}
