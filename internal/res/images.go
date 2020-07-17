package res

import (
	"fmt"
	"strconv"

	"fyne.io/fyne"
	"github.com/EricNeid/go-weatherstation/internal/util"
	"github.com/EricNeid/go-weatherstation/internal/weather"
)

var log = util.Log{Context: "images"}

var backgroundWeather = map[int]fyne.Resource{
	2:                      resourceBackgroundthunderJpg,
	3:                      resourceBackgrounddrizzleJpg,
	5:                      resourceBackgroundrainJpg,
	6:                      resourceBackgroundsnowJpg,
	7:                      resourceBackgroundmistJpg,
	weather.ConditionClear: resourceBackgroundclearJpg,
	9:                      resourceBackgroundtornadoJpg,
}

// GetBackgroundImage returns file path to background image matching the
// given condition code. See https://openweathermap.org/weather-conditions for conditions.
// Only the first digit is used to display the primary weather condition (all types of snow are returned with same image).
func GetBackgroundImage(weatherConditionID int) (fyne.Resource, error) {
	log.D("GetBackgroundImage", fmt.Sprintf("Condition id is: %d", weatherConditionID))

	cond := strconv.Itoa(weatherConditionID)
	primaryCond, err := strconv.Atoi(string(cond[0]))
	if err != nil {
		return nil, err
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
	return resourceAppiconPng
}
