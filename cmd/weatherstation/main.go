package main

import (
	"fmt"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
	"github.com/EricNeid/go-weatherstation/res"
	"github.com/EricNeid/go-weatherstation/ui"
	"github.com/EricNeid/go-weatherstation/util"
)

const screenSaverSwitchDelay = 15 * time.Second
const returnToScreenSaverDelay = 30 * time.Second

var log = util.Log{Context: "main"}

type weatherstation struct {
	app    fyne.App
	window fyne.Window

	container   *widget.Box
	screenSaver *ui.ScreenSaver
	weather     *ui.Weather
}

func main() {
	res.CurrentLocale = res.DE

	a := app.New()
	w := a.NewWindow("Weatherinformation")
	w.SetFixedSize(true)
	app := weatherstation{
		app:         a,
		window:      w,
		container:   widget.NewHBox(),
		screenSaver: ui.NewScreenSaver(),
		weather:     ui.NewWeather(),
	}

	// not working
	if err := app.weather.SetBackground("res/weather/background_clear.jpg"); err != nil {
		app.showError(err)
	}

	app.startScreenSaverUpdates()
	app.handleScreenSaverTouches()
	app.handleCloseButtonTouches()
	app.start()
}

func (app *weatherstation) startScreenSaverUpdates() {
	backgroundImages := util.NewFileRingList("images")
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
