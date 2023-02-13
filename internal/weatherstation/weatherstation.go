// Package weatherstation is a simple gui application to display either some screen saver images
// like a digitail photo frame or the weather forecast for the next 3 days (from openweather).
package weatherstation

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/EricNeid/go-weatherstation/internal/logger"
	"github.com/EricNeid/go-weatherstation/internal/util"
	"github.com/EricNeid/go-weatherstation/internal/view"
	"github.com/EricNeid/go-weatherstation/internal/weather"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
)

// switch displayed screensaver image
const screenSaverSwitchDelay = 30 * time.Second

// update displayed clock
const clockUpdateDelay = 1 * time.Minute

// retrieve new weather data
const weatherUpdateDelay = 1 * time.Hour

// check wether to display screensaver or weather information
const checkScreenDelay = 30 * time.Minute

// switch back to screensaver after user requested weather information
const switchBackDelay = 30 * time.Second

// enum for currently displayed screen.
type screen int

const (
	screensaver screen = iota
	weatherinformation
)

// fixedShowWeather set the time when the weather screen is always displayed
// (not only when user touches the screen).
// For example in the morning you may always want to see the weather information.
var fixedShowWeather = util.TimeInterval{
	StartHour: 6,
	EndHour:   9,
}

var log = logger.Log{Context: "weatherstation"}

// App represents the main weatherstation application.
// It glues ui and services together.
type App struct {
	app    fyne.App
	window fyne.Window

	canvas      *fyne.Container
	screenSaver *view.ScreenSaver
	weather     *view.Weather

	openWeatherKey string

	currentScreen chan screen
}

// NewApp generates new instance of weatherstation application.
func NewApp(fyneApp fyne.App, window fyne.Window, city, keyFile, imageDir string) App {
	// channel currentScreen is used to changed the currently displayed view
	currentScreen := make(chan screen)

	uiWeather, weatherViewModel := view.NewWeather(func() {
		log.D("closeTapped", "closing app")
		fyneApp.Quit()
	})
	uiScreensaver, screensaverViewModel := view.NewScreenSaver(func() {
		log.D("screensaver tapped", "switching to weather information")
		currentScreen <- weatherinformation
		go func() {
			time.Sleep(switchBackDelay)
			log.D("screensaver tapped", "switching back to screensaver")
			currentScreen <- screensaver
		}()
	})

	app := App{
		app:    fyneApp,
		window: window,
		canvas: container.New(
			layout.NewMaxLayout(),
			uiScreensaver,
			uiWeather,
		),
		weather:       weatherViewModel,
		screenSaver:   screensaverViewModel,
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
	app.window.SetContent(app.canvas)
	app.currentScreen <- screensaver // initial view is the screensaver
	app.window.ShowAndRun()
}

func (app *App) loadKey(keyFile string) {
	key, err := os.ReadFile(keyFile)
	if err != nil {
		app.showError(err)
	} else {
		// cleanup api key
		apiKey := strings.TrimSpace(string(key))
		apiKey = strings.ReplaceAll(apiKey, "\n", "")
		apiKey = strings.ReplaceAll(apiKey, "\r", "")

		app.openWeatherKey = apiKey
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
				app.weather.SetCurrentTemperatureData(&current)
			}

			forecast, err := weather.Forecast(app.openWeatherKey, city)
			if err != nil {
				app.showError(err)
			} else {
				app.weather.SetForecastTemperatureData(&forecast)
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
			app.canvas.Refresh()
		}
	}()
}

func (app *App) showError(err error) {
	log.D("showError", err.Error())
	dialog.ShowError(err, app.window)
}
