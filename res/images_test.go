package res

import (
	"testing"

	"github.com/EricNeid/go-weatherstation/internal/verify"
)

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
