package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	weatherstation "github.com/EricNeid/go-weatherstation"
	"github.com/EricNeid/go-weatherstation/assets"
	"github.com/EricNeid/go-weatherstation/internal/logger"

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
	fyneApp := app.New()
	fyneApp.Settings().SetTheme(theme.DarkTheme())
	fyneApp.SetIcon(assets.GetAppIcon())
	w := fyneApp.NewWindow("Weatherinformation")

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

	app := weatherstation.NewApp(fyneApp, w, city, keyFile, imageDir)

	app.Start()
}
