package res

import "testing"
import verify "github.com/EricNeid/go-weatherstation/internal/test"

func TestGetBackgroundImage(t *testing.T) {
	// action
	result, err := GetBackgroundImage(6)
	// verify
	verify.Ok(t, err)
	verify.Equals(t, "assets/weather/background_snow.jpg", result)

	// action
	result, err = GetBackgroundImage(601)
	// verify
	verify.Ok(t, err)
	verify.Equals(t, "assets/weather/background_snow.jpg", result)

	// action
	result, err = GetBackgroundImage(615)
	// verify
	verify.Ok(t, err)
	verify.Equals(t, "assets/weather/background_snow.jpg", result)
}
