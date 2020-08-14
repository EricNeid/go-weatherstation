package ui

import (
	"testing"
	"time"

	"fyne.io/fyne/test"
	"github.com/EricNeid/go-weatherstation/internal/verify"
)

const testDir = "../../test"

func TestNewScreenSaver(t *testing.T) {
	// arrange
	window := test.NewApp().NewWindow("TestNewScreenSaver")
	view, viewModel := NewScreenSaver(nil)
	// action
	window.SetContent(view)
	// verify
	verify.NotNil(t, viewModel.clock, "clock is nil")
	verify.NotNil(t, viewModel.image, "image is nil")
}

func TestSetBackground(t *testing.T) {
	// arrange
	window := test.NewApp().NewWindow("TestSetBackground")
	view, viewModel := NewScreenSaver(nil)
	window.SetContent(view)

	// action
	err := viewModel.SetBackground(testDir + "/testdata/img-1.png")
	// verify
	verify.Ok(t, err)
	verify.Equals(t, testDir+"/testdata/img-1.png", viewModel.image.File)

	// action
	err = viewModel.SetBackground(testDir + "/testdata/dir/img-2.png")
	// verify
	verify.Ok(t, err)
	verify.Equals(t, testDir+"/testdata/dir/img-2.png", viewModel.image.File)
}

func TestScreenSaverSetTime(t *testing.T) {
	// arrange
	window := test.NewApp().NewWindow("TestScreenSaverSetTime")
	time := time.Now()
	view, viewModel := NewScreenSaver(nil)
	window.SetContent(view)
	// action
	viewModel.SetTime(time)
	// verify
	verify.Equals(t, time.Format("Mon 15:04"), viewModel.clock.Text)
}
