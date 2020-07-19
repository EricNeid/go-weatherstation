package ui

import (
	"fmt"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/EricNeid/go-weatherstation/internal/util"
)

// ScreenSaver represents a clickable background image.
// View property represents actual UI which can be added to a window.
// Read from Taps to detect clicks on the whole screen.
type ScreenSaver struct {
	View  *fyne.Container
	image *canvas.Image
	clock *widget.Label
	Taps  <-chan bool
}

// NewScreenSaver constructs a new instance of a ScreenSaver widget.
func NewScreenSaver() *ScreenSaver {
	taps := make(chan bool)
	s := ScreenSaver{
		Taps: taps,
	}
	s.clock = widget.NewLabel("clock")
	s.clock.TextStyle.Bold = true

	s.image = &canvas.Image{FillMode: canvas.ImageFillStretch}

	footer := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		layout.NewSpacer(),
		s.clock,
	)
	s.View = fyne.NewContainerWithLayout(layout.NewMaxLayout(),
		s.image,
		NewTransparentButton(func() {
			taps <- true
		}),
		fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
			layout.NewSpacer(),
			footer,
		),
	)

	return &s
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

// SetTime sets the time to be displayed.
func (screenSaver *ScreenSaver) SetTime(t time.Time) {
	str := t.Format("Mon 15:04")
	screenSaver.clock.SetText(str)
}

// Hide makes the ui invisible
func (screenSaver *ScreenSaver) Hide() {
	if !screenSaver.View.Hidden {
		screenSaver.View.Hide()
	}
}

// Show makes the ui visible
func (screenSaver *ScreenSaver) Show() {
	if screenSaver.View.Hidden {
		screenSaver.View.Show()
	}
}
