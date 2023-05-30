// Package weatherstation is a simple gui application to display either some screen saver images
// like a digital photo frame or the weather forecast for the next 3 days (from openweather).
package goweatherstation

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/EricNeid/go-weatherstation/ringlist"
	"github.com/EricNeid/go-weatherstation/ui"
	"github.com/EricNeid/go-weatherstation/weather"

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

// TimeInterval defines a time interval, determined by start and end hour
type timeInterval struct {
	StartHour int
	EndHour   int
}

// Contains checks if the given time is within the given interval. The date part is ignored.
func (interval *timeInterval) contains(t time.Time) bool {
	h := t.Hour()
	if h >= interval.StartHour && h <= interval.EndHour {
		return true
	}
	return false
}

// fixedShowWeather set the time when the weather screen is always displayed
// (not only when user touches the screen).
// For example in the morning you may always want to see the weather information.
var fixedShowWeather = timeInterval{
	StartHour: 6,
	EndHour:   9,
}

// App represents the main weatherstation application.
// It glues ui and services together.
type App struct {
	app    fyne.App
	window fyne.Window

	canvas      *fyne.Container
	screenSaver *ui.ScreenSaver
	weather     *ui.Weather

	openWeatherKey string

	currentScreen chan screen
}

// NewApp generates new instance of weatherstation application.
func NewApp(fyneApp fyne.App, window fyne.Window, city, keyFile, imageDir string) App {
	// channel currentScreen is used to changed the currently displayed view
	currentScreen := make(chan screen)

	weatherView, weatherViewModel := ui.NewWeather(func() {
		log.Println("weatherstation", "close tapped", "closing app")
		fyneApp.Quit()
	})
	screensaverView, screensaverViewModel := ui.NewScreenSaver(func() {
		log.Println("weatherstation", "screensaver tapped", "switching to weather information")
		currentScreen <- weatherinformation
		go func() {
			time.Sleep(switchBackDelay)
			log.Println("weatherstation", "screensaver tapped", "switching back to screensaver")
			currentScreen <- screensaver
		}()
	})

	app := App{
		app:    fyneApp,
		window: window,
		canvas: container.New(
			layout.NewMaxLayout(),
			screensaverView,
			weatherView,
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
	log.Println("weatherstation", "Start")
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
	log.Println("weatherstation", "startScreenSaverUpdates", imageDir)
	backgroundImages := ringlist.NewFileRingList(imageDir)
	backgroundImages.Shuffle()
	go func() {
		for {
			file, _ := backgroundImages.Next()
			log.Println("weatherstation", "startScreenSaverUpdates", fmt.Sprintf("Switching to %s", file))
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
	log.Println("weatherstation", "startCheckScreen")
	go func() {
		for {
			log.Println("weatherstation", "startCheckScreen", "checking screen to display")
			if fixedShowWeather.contains(time.Now()) {
				app.currentScreen <- weatherinformation
			} else {
				app.currentScreen <- screensaver
			}
			time.Sleep(checkScreenDelay)
		}
	}()
}

func (app *App) startClockUpdates() {
	log.Println("weatherstation", "startClockUpdates")
	go func() {
		for {
			app.weather.SetTime(time.Now())
			app.screenSaver.SetTime(time.Now())
			time.Sleep(clockUpdateDelay)
		}
	}()
}

func (app *App) startWeatherInformationUpdates(city string) {
	log.Println("weatherstation", "startWeatherInformationUpdates", city)
	go func() {
		for {
			log.Println("weatherstation", "startWeatherInformationUpdates", "update weather information")

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
	log.Println("weatherstation", "startCurrentScreenHandler")
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
	log.Println("weatherstation", "showError", err)
	dialog.ShowError(err, app.window)
}
