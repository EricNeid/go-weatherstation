package res

import "fmt"
import "strconv"
import "github.com/EricNeid/go-weatherstation/util"

import "fyne.io/fyne"

var log = util.Log{Context: "images"}

var backgroundWeather = map[int]string{
	2: "res/weather/background_thunder.jpg",
	3: "res/weather/background_drizzle.jpg",
	5: "res/weather/background_rain.jpg",
	6: "res/weather/background_snow.jpg",
	7: "res/weather/background_mist.jpg",
	8: "res/weather/background_clear.jpg",
	9: "res/weather/background_tornado.jpg",
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
