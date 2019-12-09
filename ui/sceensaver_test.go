package ui

import "testing"

import "github.com/EricNeid/go-weatherstation/internal/test"

func TestNew(t *testing.T) {
	// action
	screensaver := NewScreenSaver()
	// verify
	test.Equals(t, 1, len(screensaver.Children))
}

func TestSetBackground(t *testing.T) {
	// arrange
	screensaver := NewScreenSaver()

	// action
	err := screensaver.SetBackground("../testdata/img-1.png")
	// verify that only on children is present
	test.Ok(t, err)
	test.Equals(t, 1, len(screensaver.Children))
	test.Equals(t, "../testdata/img-1.png", screensaver.image.File)

	// action
	err = screensaver.SetBackground("../testdata/dir/img-2.png")
	// verify that still only one children is present
	test.Ok(t, err)
	test.Equals(t, 1, len(screensaver.Children))
	test.Equals(t, "../testdata/dir/img-2.png", screensaver.image.File)
}
