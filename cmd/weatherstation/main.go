package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/EricNeid/go-weatherstation/internal/res"
	"github.com/EricNeid/go-weatherstation/internal/logger"
	"github.com/EricNeid/go-weatherstation/internal/weatherstation"

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
	logger.Init()
	args := parseArgs()

	a := app.New()
	w := a.NewWindow("Weatherinformation")

	// set app icon
	appIcon := res.GetAppIcon()
	a.SetIcon(appIcon)

	// check api key file exists
	if _, err := os.Stat(args.keyFile); os.IsNotExist(err) {
		dialog.ShowError(errors.New("Could not load api key"), w)
	}

	// check if image dir exists
	if _, err := os.Stat(args.imageDir); os.IsNotExist(err) {
		dialog.ShowError(errors.New("Could not find image directory"), w)
	}

	// set app size
	if args.fullscreen {
		w.SetFullScreen(true)
	} else {
		w.SetFixedSize(true)
		w.Resize(fyne.NewSize(800, 480))
	}

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
	flag.StringVar(&args.imageDir, "screensaver", "screensaver", "directory, containing images for the screensave")
	flag.StringVar(&args.keyFile, "key", "api.key", "file, containing api key for openweather")

	flag.Parse()

	return args
}
