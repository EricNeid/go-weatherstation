package main

import (
	"fmt"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
	"github.com/EricNeid/go-netconvert/ui"
	"github.com/EricNeid/go-netconvert/util"
)

const screenSaverSwitchDelay = 15 * time.Second
const returnToScreenSaverDelay = 30 * time.Second

type weatherstation struct {
	window fyne.Window

	container   *widget.Box
	screenSaver *ui.ScreenSaver
	weather     *ui.Weather
}

func main() {
	w := app.New().NewWindow("Weatherinformation")
	app := weatherstation{
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
	app.start()
}

func (app *weatherstation) startScreenSaverUpdates() {
	backgroundImages := util.NewFileRingList("images")
	go func() {
		for {
			file, _ := backgroundImages.Next()
			fmt.Printf("Switching to %s\n", file)
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
			app.showWeatherInfo()
			time.Sleep(returnToScreenSaverDelay)
			fmt.Printf("Switching back to screensaver\n")
			app.showScreenSaver()
		}
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
