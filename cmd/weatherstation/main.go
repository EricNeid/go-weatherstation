package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
	"github.com/EricNeid/go-weatherstation/res"
	"github.com/EricNeid/go-weatherstation/services"
	"github.com/EricNeid/go-weatherstation/ui"
	"github.com/EricNeid/go-weatherstation/util"
)

const screenSaverSwitchDelay = 30 * time.Second
const returnToScreenSaverDelay = 45 * time.Second
const clockUpdateDelay = 1 * time.Minute
const weatherUpdateDelay = 1 * time.Hour

const city = "Berlin"

var log = util.Log{Context: "main"}

type weatherstation struct {
	app    fyne.App
	window fyne.Window

	container   *widget.Box
	screenSaver *ui.ScreenSaver
	weather     *ui.Weather

	openWeatherKey string
}

func main() {
	res.CurrentLocale = res.DE

	args := parseArgs()

	a := app.New()
	a.SetIcon(res.GetAppIcon())
	w := a.NewWindow("Weatherinformation")
	w.SetFixedSize(true)
	w.SetFullScreen(args.fullscreen)
	w.Resize(fyne.NewSize(800, 480))
	app := weatherstation{
		app:         a,
		window:      w,
		container:   widget.NewHBox(),
		screenSaver: ui.NewScreenSaver(),
		weather:     ui.NewWeather(),
	}
	app.loadKey()

	app.startClockUpdates()
	app.startScreenSaverUpdates(args.imageDir)
	app.startWeatherInformationUpdates()
	app.handleScreenSaverTouches()
	app.handleCloseButtonTouches()
	app.start()
}

func (app *weatherstation) loadKey() {
	key, err := ioutil.ReadFile("api.key")
	if err != nil {
		app.showError(err)
	} else {
		app.openWeatherKey = string(key)
	}
}

func (app *weatherstation) startScreenSaverUpdates(imageDir string) {
	backgroundImages := util.NewFileRingList(imageDir)
	go func() {
		for {
			file, _ := backgroundImages.Next()
			log.D("startScreenSaverUpdates", fmt.Sprintf("Switching to %s", file))
			if err := app.screenSaver.SetBackground(file); err != nil {
				app.showError(err)
			}
			time.Sleep(screenSaverSwitchDelay)
		}
	}()
}

func (app *weatherstation) handleScreenSaverTouches() {
	go func() {
		for {
			<-app.screenSaver.Touches
			log.D("handleScreenSaverTouches", "Switching to weather")
			app.showWeatherInfo()
			time.Sleep(returnToScreenSaverDelay)
			log.D("handleScreenSaverTouches", "Switching back to screensaver")
			app.showScreenSaver()
		}
	}()
}

func (app *weatherstation) handleCloseButtonTouches() {
	go func() {
		<-app.weather.CloseTouches
		log.D("handleCloseButtonTouches", "Closing app")
		app.app.Quit()
	}()
}

func (app *weatherstation) startClockUpdates() {
	go func() {
		app.weather.SetTime(time.Now())
		app.screenSaver.SetTime(time.Now())
		time.Sleep(clockUpdateDelay)
	}()
}

func (app *weatherstation) startWeatherInformationUpdates() {
	go func() {
		log.D("startWeatherInformationUpdates", "Update weather information")

		current, err := services.GetWeather(app.openWeatherKey, city)
		if err != nil {
			app.showError(err)
		} else {
			app.weather.SetCurrentTemperatureData(*current)
		}

		forecast, err := services.GetWeatherForecast(app.openWeatherKey, city)
		if err != nil {
			app.showError(err)
		} else {
			app.weather.SetForecastTemperatureData(*forecast)
		}

		time.Sleep(weatherUpdateDelay)
	}()
}

func (app *weatherstation) start() {
	app.window.SetFullScreen(false)
	app.window.SetContent(app.container)
	app.showScreenSaver()
	app.window.ShowAndRun()
}

func (app *weatherstation) showScreenSaver() {
	app.container.Children = []fyne.CanvasObject{app.screenSaver}
	app.container.Refresh()
}

func (app *weatherstation) showWeatherInfo() {
	app.container.Children = []fyne.CanvasObject{app.weather}
	app.container.Refresh()
	app.weather.Refresh()
}

func (app *weatherstation) showError(err error) {
	dialog.ShowError(err, app.window)
}

func (app *weatherstation) showInfo(msg string) {
	dialog.ShowInformation("foo", msg, app.window)
}
