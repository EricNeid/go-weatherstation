package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/EricNeid/go-weatherstation/internal/assets"
	"github.com/EricNeid/go-weatherstation/internal/logger"
	"github.com/EricNeid/go-weatherstation/internal/weatherstation"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
)

type args struct {
	fullscreen bool
	imageDir   string
	keyFile    string
}

const city = "Berlin"

func main() {
	assets.CurrentLocale = assets.DE
	logger.Init()
	args := parseArgs()

	a := app.New()
	// theme
	a.Settings().SetTheme(theme.DarkTheme())

	w := a.NewWindow("Weatherinformation")

	// set app icon
	appIcon := assets.GetAppIcon()
	a.SetIcon(appIcon)

	// check api key file exists
	if _, err := os.Stat(args.keyFile); os.IsNotExist(err) {
		dialog.ShowError(errors.New("could not load api key"), w)
	}

	// check if image dir exists
	if _, err := os.Stat(args.imageDir); os.IsNotExist(err) {
		dialog.ShowError(errors.New("could not find image directory"), w)
	}

	// set app size
	if args.fullscreen {
		w.SetFullScreen(true)
	} else {
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
