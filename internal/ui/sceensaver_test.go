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
	unit := NewScreenSaver(nil)
	// action
	window.SetContent(unit.View)
	// verify
	verify.NotNil(t, unit.clock, "clock is nil")
	verify.NotNil(t, unit.image, "image is nil")
}

func TestSetBackground(t *testing.T) {
	// arrange
	window := test.NewApp().NewWindow("TestSetBackground")
	unit := NewScreenSaver(nil)
	window.SetContent(unit.View)

	// action
	err := unit.SetBackground(testDir + "/testdata/img-1.png")
	// verify
	verify.Ok(t, err)
	verify.Equals(t, testDir+"/testdata/img-1.png", unit.image.File)

	// action
	err = unit.SetBackground(testDir + "/testdata/dir/img-2.png")
	// verify
	verify.Ok(t, err)
	verify.Equals(t, testDir+"/testdata/dir/img-2.png", unit.image.File)
}

func TestScreenSaverSetTime(t *testing.T) {
	// arrange
	window := test.NewApp().NewWindow("TestScreenSaverSetTime")
	time := time.Now()
	unit := NewScreenSaver(nil)
	window.SetContent(unit.View)
	// action
	unit.SetTime(time)
	// verify
	verify.Equals(t, time.Format("Mon 15:04"), unit.clock.Text)
}
