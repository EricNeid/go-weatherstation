package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/test"
)

func newTestWindow() fyne.Window {
	w := test.NewApp().NewWindow("Test")
	w.Show()
	return w
}
