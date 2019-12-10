package ui

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/EricNeid/go-weatherstation/res"
	"github.com/EricNeid/go-weatherstation/util"
	"time"
)

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
	weather := &Weather{
		widget.Box{},
		&canvas.Image{FillMode: canvas.ImageFillOriginal},
		widget.NewLabel("Clock"),
		widget.NewLabel("City"),
		widget.NewLabel("Current Temperature"),
		make(chan bool),
	}
	weather.ExtendBaseWidget(weather)

	header := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		layout.NewSpacer(),
		widget.NewVBox(weather.city, weather.currentTemperature),
		layout.NewSpacer(),
	)

	footer := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		widget.NewButton(res.GetLabel("close"), func() {
			weather.CloseTouches <- true
		}),
		layout.NewSpacer(),
		weather.clock,
	)

	center := widget.NewLabel("Center")

	weather.Children = []fyne.CanvasObject{
		fyne.NewContainerWithLayout(layout.NewMaxLayout(),
			fyne.NewContainerWithLayout(layout.NewMaxLayout(),
				weather.background,
			),
			fyne.NewContainerWithLayout(layout.NewBorderLayout(header, footer, nil, nil), header, footer, center),
		),
	}
	return weather
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
