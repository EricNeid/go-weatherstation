package ui

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/widget"
	"github.com/EricNeid/go-weatherstation/res"
	"github.com/EricNeid/go-weatherstation/util"
)

// Weather represents information view for weather information
type Weather struct {
	widget.Box
	background *canvas.Image
}

// NewWeather constructs a new instance of a NewWeather widget.
func NewWeather() *Weather {
	weather := &Weather{
		widget.Box{},
		&canvas.Image{FillMode: canvas.ImageFillOriginal},
	}
	weather.ExtendBaseWidget(weather)

	weather.Children = []fyne.CanvasObject{
		fyne.NewContainer(
			widget.NewVBox(
				weather.background,
			),
			widget.NewVBox(
				widget.NewLabel("Hello"),
				widget.NewHBox(
					widget.NewLabel(res.GetLabel("today")),
					widget.NewLabel(res.GetLabel("tomorrow")),
					widget.NewLabel(res.GetLabel("aftertomorrow")),
				),
			),
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
