package res

import (
	"fmt"
	"strconv"

	"fyne.io/fyne"
	"github.com/EricNeid/go-weatherstation/util"
)

var log = util.Log{Context: "images"}

var backgroundWeather = map[int]string{
	2: "assets/weather/background_thunder.jpg",
	3: "assets/weather/background_drizzle.jpg",
	5: "assets/weather/background_rain.jpg",
	6: "assets/weather/background_snow.jpg",
	7: "assets/weather/background_mist.jpg",
	8: "assets/weather/background_clear.jpg",
	9: "assets/weather/background_tornado.jpg",
}

// GetBackgroundImage returns file path to background image matching the
// given condition code. See https://openweathermap.org/weather-conditions for conditions.
// Only the first digit is used to display the primary weather condition (all types of snow are returned with same image).
func GetBackgroundImage(weatherConditionID int) (string, error) {
	log.D("GetBackgroundImage", fmt.Sprintf("Condition id is: %d", weatherConditionID))

	cond := strconv.Itoa(weatherConditionID)
	primaryCond, err := strconv.Atoi(string(cond[0]))
	if err != nil {
		return "", err
	}

	return backgroundWeather[primaryCond], nil
}

// GetConditionIcon returns a new resource for given weather icon.
// The given icon is converted to an URL and the resource is retrieved.
// See https://openweathermap.org/weather-conditions for more details.
func GetConditionIcon(weatherConditionIcon string) (fyne.Resource, error) {
	url := fmt.Sprintf("http://openweathermap.org/img/w/%s.png", weatherConditionIcon)
	return fyne.LoadResourceFromURLString(url)
}

// GetAppIcon returns the application icon.
func GetAppIcon() fyne.Resource {
	res, err := fyne.LoadResourceFromPath("assets/ic-sunny.png")
	if err != nil {
		log.E("GetAppIcon", err)
	}
	return res
}
