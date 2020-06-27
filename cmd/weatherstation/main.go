package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/EricNeid/go-weatherstation"
	"github.com/EricNeid/go-weatherstation/res"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
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
	a.SetIcon(res.GetAppIcon())
	w := a.NewWindow("Weatherinformation")
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
