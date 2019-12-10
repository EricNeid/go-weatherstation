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

	CloseTouches chan bool
}

// NewWeather constructs a new instance of a NewWeather widget.
func NewWeather() *Weather {
	w := &Weather{
		widget.Box{},
		&canvas.Image{FillMode: canvas.ImageFillOriginal},
		widget.NewLabel("City"),
		widget.NewLabel("Current Temperature"),
		widget.NewLabel("Clock"),
		make(chan bool),
	}
	w.ExtendBaseWidget(w)
	w.city.Alignment = fyne.TextAlignCenter
	w.city.TextStyle.Bold = true
	w.currentTemperature.Alignment = fyne.TextAlignCenter

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

	center := widget.NewLabel("Center")

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
	weather.currentTemperature.SetText(fmt.Sprintf(res.GetLabel("currentTemperature"), data.Main.Temp))
}
