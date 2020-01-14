package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
	"EricNeid/go-weatherstation/res"
	"EricNeid/go-weatherstation/services"
	"EricNeid/go-weatherstation/ui"
	"EricNeid/go-weatherstation/util"
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
	app.loadKey(args.keyFile)

	app.startClockUpdates()
	app.startCurrentScreenHandler()
	app.startScreenSaverUpdates(args.imageDir)
	app.startWeatherInformationUpdates()
	app.handleScreenSaverTouches()
	app.handleCloseButtonTouches()
	app.start()
}

func (app *weatherstation) loadKey(keyFile string) {
	key, err := ioutil.ReadFile("api.key")
	if err != nil {
		app.showError(err)
	} else {
		app.openWeatherKey = string(key)
	}
}

func (app *weatherstation) handleScreenSaverTouches() {
	go func() {
		for {
			<-app.screenSaver.Touches
			log.D("handleScreenSaverTouches", "Switching to weather")
			app.showWeatherInfo()
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

func (app *weatherstation) startScreenSaverUpdates(imageDir string) {
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
func (app *weatherstation) startCurrentScreenHandler() {
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

func (app *weatherstation) startClockUpdates() {
	log.D("startClockUpdates", "")
	go func() {
		for {
			app.weather.SetTime(time.Now())
			app.screenSaver.SetTime(time.Now())
			time.Sleep(clockUpdateDelay)
		}
	}()
}

func (app *weatherstation) startWeatherInformationUpdates() {
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

func (app *weatherstation) start() {
	log.D("start", "")
	app.window.SetContent(app.container)
	app.showScreenSaver()
	app.window.ShowAndRun()
}

func (app *weatherstation) showScreenSaver() {
	log.D("showScreenSaver", "")
	if len(app.container.Children) > 0 && app.container.Children[0] == app.screenSaver {
		return
	}
	app.container.Children = []fyne.CanvasObject{app.screenSaver}
	app.container.Refresh()
}

func (app *weatherstation) showWeatherInfo() {
	log.D("showWeatherInfo", "")
	if len(app.container.Children) > 0 && app.container.Children[0] == app.weather {
		return
	}
	app.container.Children = []fyne.CanvasObject{app.weather}
	app.container.Refresh()
}

func (app *weatherstation) showError(err error) {
	log.D("showError", fmt.Sprintf("%s", err.Error()))
	dialog.ShowError(err, app.window)
}
