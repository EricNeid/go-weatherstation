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
	unit := NewScreenSaver()
	// action
	window.SetContent(unit.UI)
	// verify
	verify.NotNil(t, unit.clock, "clock is nil")
	verify.NotNil(t, unit.image, "image is nil")
}

func TestSetBackground(t *testing.T) {
	// arrange
	window := test.NewApp().NewWindow("TestSetBackground")
	unit := NewScreenSaver()
	window.SetContent(unit.UI)

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
	unit := NewScreenSaver()
	window.SetContent(unit.UI)
	// action
	unit.SetTime(time)
	// verify
	verify.Equals(t, time.Format("Mon 15:04"), unit.clock.Text)
}
