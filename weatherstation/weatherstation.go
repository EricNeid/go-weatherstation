// Package weatherstation is a simple gui application to display either some screen saver images
// like a digitail photo frame or the weather forecast for the next 3 days (from openweather).
package weatherstation

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/EricNeid/go-weatherstation/services"
	"github.com/EricNeid/go-weatherstation/ui"
	"github.com/EricNeid/go-weatherstation/util"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

const screenSaverSwitchDelay = 30 * time.Second
const clockUpdateDelay = 1 * time.Minute
const weatherUpdateDelay = 1 * time.Hour
const checkScreenDelay = 1 * time.Minute

// fixedShowWeather set the time when the weather screen is always displayed
// (not only when user touches the screen).
// For example in the morning you may always want to see the weather information.
var fixedShowWeather = util.TimeInterval{
	StartHour: 6,
	EndHour:   9,
}

var log = util.Log{Context: "weatherstation"}

// App represents the main weatherstation application.
// It glues ui and services together.
type App struct {
	app    fyne.App
	window fyne.Window

	container   *widget.Box
	screenSaver *ui.ScreenSaver
	weather     *ui.Weather

	openWeatherKey string
}

// NewApp generates new instance of weatherstation application.
func NewApp(fyneApp fyne.App, window fyne.Window, city string, keyFile string, imageDir string) App {
	app := App{
		app:         fyneApp,
		window:      window,
		container:   widget.NewHBox(),
		screenSaver: ui.NewScreenSaver(),
		weather:     ui.NewWeather(),
	}

	app.loadKey(keyFile)
	app.startClockUpdates()
	app.startCurrentScreenHandler()
	app.startScreenSaverUpdates(imageDir)
	app.startWeatherInformationUpdates(city)
	app.handleScreenSaverTouches()
	app.handleCloseButtonTouches()

	return app
}

// Start starts the ui lifecycle of this weatherstation.
func (app *App) Start() {
	log.D("start", "")
	app.window.SetContent(app.container)
	app.showScreenSaver()
	app.window.ShowAndRun()
}

func (app *App) loadKey(keyFile string) {
	key, err := ioutil.ReadFile(keyFile)
	if err != nil {
		app.showError(err)
	} else {
		// cleanup api key
		apiKey := strings.TrimSpace(string(key))
		apiKey = strings.ReplaceAll(apiKey, "\n", "")
		apiKey = strings.ReplaceAll(apiKey, "\r", "")

		app.openWeatherKey = string(apiKey)
	}
}

func (app *App) handleScreenSaverTouches() {
	go func() {
		for {
			<-app.screenSaver.Touches
			log.D("handleScreenSaverTouches", "Switching to weather")
			app.showWeatherInfo()
		}
	}()
}

func (app *App) handleCloseButtonTouches() {
	go func() {
		<-app.weather.CloseTouches
		log.D("handleCloseButtonTouches", "Closing app")
		app.app.Quit()
	}()
}

func (app *App) startScreenSaverUpdates(imageDir string) {
	log.D("startScreenSaverUpdates", "")
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

// startCurrentScreenHandler check which screen should be displayed on a regular basis.
// During fixedShowWeather the weather screen should be display by default.
func (app *App) startCurrentScreenHandler() {
	log.D("startCurrentScreenHandler", "")
	go func() {
		for {
			log.D("startCurrentScreenHandler", "Check screen to display")
			if fixedShowWeather.Contains(time.Now()) {
				app.showWeatherInfo()
			} else {
				app.showScreenSaver()
			}
			time.Sleep(checkScreenDelay)
		}
	}()
}

func (app *App) startClockUpdates() {
	log.D("startClockUpdates", "")
	go func() {
		for {
			app.weather.SetTime(time.Now())
			app.screenSaver.SetTime(time.Now())
			time.Sleep(clockUpdateDelay)
		}
	}()
}

func (app *App) startWeatherInformationUpdates(city string) {
	log.D("startWeatherInformationUpdates", "")
	go func() {
		for {
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
		}
	}()
}

func (app *App) showScreenSaver() {
	log.D("showScreenSaver", "")
	if len(app.container.Children) > 0 && app.container.Children[0] == app.screenSaver {
		return
	}
	app.container.Children = []fyne.CanvasObject{app.screenSaver}
	app.container.Refresh()
}

func (app *App) showWeatherInfo() {
	log.D("showWeatherInfo", "")
	if len(app.container.Children) > 0 && app.container.Children[0] == app.weather {
		return
	}
	app.container.Children = []fyne.CanvasObject{app.weather}
	app.container.Refresh()
}

func (app *App) showError(err error) {
	log.D("showError", fmt.Sprintf("%s", err.Error()))
	dialog.ShowError(err, app.window)
}
