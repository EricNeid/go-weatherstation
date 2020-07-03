package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/EricNeid/go-weatherstation/res"
	"github.com/EricNeid/go-weatherstation/weatherstation"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
)

type args struct {
	fullscreen bool
	imageDir   string
	keyFile    string
}

const city = "Berlin"

func main() {
	res.CurrentLocale = res.DE

	args := parseArgs()

	a := app.New()
	w := a.NewWindow("Weatherinformation")

	appIcon, err := res.GetAppIcon()
	if err != nil {
		dialog.ShowError(errors.New("Could not load app icon"), w)
	}
	if _, err := os.Stat(args.keyFile); os.IsNotExist(err) {
		dialog.ShowError(errors.New("Could not load api key"), w)
	}
	if _, err := os.Stat(args.imageDir); os.IsNotExist(err) {
		dialog.ShowError(errors.New("Could not find image directory"), w)
	}

	a.SetIcon(appIcon)
	w.SetFixedSize(true)
	w.SetFullScreen(args.fullscreen)
	w.Resize(fyne.NewSize(800, 480))

	app := weatherstation.NewApp(a, w, city, args.keyFile, args.imageDir)

	app.Start()
}

func parseArgs() args {
	flag.Usage = func() {
		fmt.Printf("Usage: %s\n", os.Args[0])
		flag.PrintDefaults()
	}

	var args args
	flag.BoolVar(&args.fullscreen, "fullscreen", false, "show app in fullscreen mode")
	flag.StringVar(&args.imageDir, "images", "images", "directory, containing images for the screensave")
	flag.StringVar(&args.keyFile, "key", "api.key", "file, containing api key for openweather")

	flag.Parse()

	return args
}
