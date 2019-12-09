package ui

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/widget"
	"github.com/EricNeid/go-weatherstation/util"
)

// ScreenSaver represents a clickable background image.
// It provides a channel to read user clicks.
type ScreenSaver struct {
	widget.Box
	image   *canvas.Image
	Touches chan *fyne.PointEvent
}

// Tapped is called automatically when this widget is clicked.
func (screenSaver *ScreenSaver) Tapped(e *fyne.PointEvent) {
	screenSaver.Touches <- e
}

// TappedSecondary is called automatically when this widget is right clicked.
func (screenSaver *ScreenSaver) TappedSecondary(e *fyne.PointEvent) {
	screenSaver.Touches <- e
}

// NewScreenSaver constructs a new instance of a ScreenSaver widget.
func NewScreenSaver() *ScreenSaver {
	s := &ScreenSaver{
		widget.Box{},
		&canvas.Image{FillMode: canvas.ImageFillOriginal},
		make(chan *fyne.PointEvent),
	}
	s.ExtendBaseWidget(s)
	s.Children = []fyne.CanvasObject{s.image}
	return s
}

// SetBackground changes the displayed background image of this screen saver.
func (screenSaver *ScreenSaver) SetBackground(filepath string) error {
	if !util.IsFilePresent(filepath) {
		return fmt.Errorf("Given file %s does not exits", filepath)
	}
	screenSaver.image.File = filepath
	screenSaver.image.Refresh()
	return nil
}
