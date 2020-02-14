package main

import (
	"EricNeid/go-weatherstation"
	"EricNeid/go-weatherstation/res"
	"EricNeid/go-weatherstation/util"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
)

const city = "Berlin"

var log = util.Log{Context: "main"}

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
