package ui

import (
	"testing"
	"time"

	"EricNeid/go-weatherstation/internal/verify"
)

func TestNewScreenSaver(t *testing.T) {
	// arrange
	testWindow := newTestWindow()
	// action
	unit := NewScreenSaver()
	testWindow.SetContent(unit)
	// verify
	verify.Equals(t, 1, len(unit.Children))
}

func TestSetBackground(t *testing.T) {
	// arrange
	unit := NewScreenSaver()
	newTestWindow().SetContent(unit)

	// action
	err := unit.SetBackground("../testdata/img-1.png")
	// verify that only on children is present
	verify.Ok(t, err)
	verify.Equals(t, 1, len(unit.Children))
	verify.Equals(t, "../testdata/img-1.png", unit.image.File)

	// action
	err = unit.SetBackground("../testdata/dir/img-2.png")
	// verify that still only one children is present
	verify.Ok(t, err)
	verify.Equals(t, 1, len(unit.Children))
	verify.Equals(t, "../testdata/dir/img-2.png", unit.image.File)
}

func TestScreenSaverSetTime(t *testing.T) {
	// arrange
	time := time.Now()
	unit := NewScreenSaver()
	newTestWindow().SetContent(unit)
	// action
	unit.SetTime(time)
	// verify
	verify.Equals(t, time.Format("Mon 15:04"), unit.clock.Text)
}
