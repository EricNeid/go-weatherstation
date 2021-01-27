package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// TransparentButton widget has a func when clicked and no further layout
type TransparentButton struct {
	widget.BaseWidget
	OnTapped func()
}

// NewTransparentButton creates a new transparent button widget with the set tap handler
func NewTransparentButton(tapped func()) *TransparentButton {
	button := &TransparentButton{
		OnTapped: tapped,
	}
	return button
}

// Tapped is called automatically when this widget is clicked.
func (b *TransparentButton) Tapped(e *fyne.PointEvent) {
	if b.OnTapped != nil {
		b.OnTapped()
	}
}

// TappedSecondary is called automatically when this widget is right clicked.
func (b *TransparentButton) TappedSecondary(e *fyne.PointEvent) {
}
