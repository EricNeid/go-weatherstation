// Package weatherstation is a simple gui application to display either some screen saver images
// like a digitail photo frame or the weather forecast for the next 3 days (from openweather).
package weatherstation

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/EricNeid/go-weatherstation/internal/ui"
	"github.com/EricNeid/go-weatherstation/internal/util"
	"github.com/EricNeid/go-weatherstation/internal/weather"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
)

const screenSaverSwitchDelay = 30 * time.Second
const clockUpdateDelay = 1 * time.Minute
const weatherUpdateDelay = 1 * time.Hour
const checkScreenDelay = 1 * time.Minute

// enum for currently displayed screen.
type screen int

const (
	screensaver        screen = iota
	weatherinformation screen = iota
)

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

	container   *fyne.Container
	screenSaver *ui.ScreenSaver
	weather     *ui.Weather

	openWeatherKey string

	currentScreen chan screen
}

// NewApp generates new instance of weatherstation application.
func NewApp(fyneApp fyne.App, window fyne.Window, city string, keyFile string, imageDir string) App {
	// channel currentScreen is used to changed the currently displayed view
	currentScreen := make(chan screen)

	uiWeather := ui.NewWeather(func() {
		fyneApp.Quit()
	})
	uiScreenSaver := ui.NewScreenSaver(func() {
		currentScreen <- weatherinformation
	})

	app := App{
		app:    fyneApp,
		window: window,
		container: fyne.NewContainerWithLayout(
			layout.NewMaxLayout(),
			uiScreenSaver.View,
			uiWeather.View,
		),
		weather:       uiWeather,
		screenSaver:   uiScreenSaver,
		currentScreen: currentScreen,
	}

	app.loadKey(keyFile)

	app.startClockUpdates()
	app.startCheckScreen()
	app.startScreenSaverUpdates(imageDir)
	app.startWeatherInformationUpdates(city)
	app.startCurrentScreenHandler()

	return app
}

// Start starts the ui lifecycle of this weatherstation.
func (app *App) Start() {
	log.D("start", "")
	app.window.SetContent(app.container)
	app.currentScreen <- screensaver // initial view is the screensaver
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

func (app *App) startScreenSaverUpdates(imageDir string) {
	log.D("startScreenSaverUpdates", "")
	backgroundImages := util.NewFileRingList(imageDir)
	backgroundImages.Shuffle()
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

// startCheckScreen checks which screen should be displayed on a regular basis.
// During fixedShowWeather the weather screen should be display by default.0
func (app *App) startCheckScreen() {
	log.D("startCheckScreen", "")
	go func() {
		for {
			log.D("startCheckScreen", "Check screen to display")
			if fixedShowWeather.Contains(time.Now()) {
				app.currentScreen <- weatherinformation
			} else {
				app.currentScreen <- screensaver
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

			current, err := weather.Current(app.openWeatherKey, city)
			if err != nil {
				app.showError(err)
			} else {
				app.weather.SetCurrentTemperatureData(current)
			}

			forecast, err := weather.Forecast(app.openWeatherKey, city)
			if err != nil {
				app.showError(err)
			} else {
				app.weather.SetForecastTemperatureData(forecast)
			}

			time.Sleep(weatherUpdateDelay)
		}
	}()
}

func (app *App) startCurrentScreenHandler() {
	log.D("startCurrentScreenHandler", "")
	go func() {
		for {
			currentScreen := <-app.currentScreen
			if currentScreen == weatherinformation {
				app.weather.Show()
				app.screenSaver.Hide()
			} else {
				app.screenSaver.Show()
				app.weather.Hide()
			}
			app.container.Refresh()
		}
	}()
}

func (app *App) showError(err error) {
	log.D("showError", fmt.Sprintf("%s", err.Error()))
	dialog.ShowError(err, app.window)
}
