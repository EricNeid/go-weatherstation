package ui

import (
	"testing"
	"time"

	"fyne.io/fyne/v2/test"
	"github.com/EricNeid/go-weatherstation/internal/verify"
)

const testDir = "../testdata"

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
	err := viewModel.SetBackground(testDir + "/img-1.png")
	// verify
	verify.Ok(t, err)
	verify.Equals(t, testDir+"/img-1.png", viewModel.image.File)

	// action
	err = viewModel.SetBackground(testDir + "/dir/img-2.png")
	// verify
	verify.Ok(t, err)
	verify.Equals(t, testDir+"/dir/img-2.png", viewModel.image.File)
}

func TestScreenSaverSetTime(t *testing.T) {
	// arrange
	window := test.NewApp().NewWindow("TestScreenSaverSetTime")
	testTime := time.Now()
	view, viewModel := NewScreenSaver(nil)
	window.SetContent(view)
	// action
	viewModel.SetTime(testTime)
	// verify
	verify.Equals(t, testTime.Format("Mon 15:04"), viewModel.clock.Text)
}

func TestIsFilePresent_shouldReturnTrue(t *testing.T) {
	// action
	result := isFilePresent(testDir + "/img-1.png")
	// verify
	verify.Equals(t, true, result)
}

func TestIsFilePresent_shouldReturnFalse(t *testing.T) {
	// action
	result := isFilePresent(testDir + "/no-file")
	// verify
	verify.Equals(t, false, result)
}

func TestIsFilePresent_shouldReturnFalseBecauseDirectory(t *testing.T) {
	// action
	result := isFilePresent(testDir + "/dir")
	// verify
	verify.Equals(t, false, result)
}
