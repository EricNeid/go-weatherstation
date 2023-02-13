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

var (
	fullScreen = false
	imageDir   = "screensaver"
	keyFile    = "api.key"
)

const city = "Berlin"

func main() {
	// init
	assets.CurrentLocale = assets.DE
	logger.Init()
	// read cli arguments
	flag.Usage = func() {
		fmt.Printf("Usage: %s\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.BoolVar(&fullScreen, "fullscreen", fullScreen, "show app in fullscreen mode")
	flag.StringVar(&imageDir, "screensaver", imageDir, "directory, containing images for the screensaver")
	flag.StringVar(&keyFile, "key", keyFile, "file, containing api key for openweather")
	flag.Parse()

	// create application
	a := app.New()
	a.Settings().SetTheme(theme.DarkTheme())
	// set app icon
	appIcon := assets.GetAppIcon()
	a.SetIcon(appIcon)

	w := a.NewWindow("Weatherinformation")

	// check api key file exists
	if _, err := os.Stat(keyFile); os.IsNotExist(err) {
		dialog.ShowError(errors.New("could not load api key"), w)
	}

	// check if image dir exists
	if _, err := os.Stat(imageDir); os.IsNotExist(err) {
		dialog.ShowError(errors.New("could not find image directory"), w)
	}

	// set app size
	if fullScreen {
		w.SetFullScreen(true)
	} else {
		w.Resize(fyne.NewSize(800, 480))
	}

	app := weatherstation.NewApp(a, w, city, keyFile, imageDir)

	app.Start()
}
