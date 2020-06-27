package main

import (
	"github.com/EricNeid/go-weatherstation"
	"github.com/EricNeid/go-weatherstation/res"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
)

const city = "Berlin"

func main() {
	res.CurrentLocale = res.DE

	args := parseArgs()

	a := app.New()
	a.SetIcon(res.GetAppIcon())
	w := a.NewWindow("Weatherinformation")
	w.SetFixedSize(true)
	w.SetFullScreen(args.fullscreen)
	w.Resize(fyne.NewSize(800, 480))

	app := weatherstation.NewApp(a, w, city, args.keyFile, args.imageDir)
	app.Start()
}
