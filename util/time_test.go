package util

import "testing"
import "time"
import "github.com/EricNeid/go-weatherstation/internal/test"

func TestIntervalContains_shouldReturnTrue(t *testing.T) {
	// arrange
	unit := TimeInterval{
		4,
		9,
	}
	testData := time.Date(1, 1, 1, 5, 0, 0, 0, time.UTC)
	// action
	result := unit.Contains(testData)
	// verify
	test.Equals(t, true, result)
}

func TestIntervalContains_shouldReturnFalse(t *testing.T) {
	// arrange
	unit := TimeInterval{
		4,
		9,
	}
	testData := time.Date(1, 1, 1, 10, 0, 0, 0, time.UTC)
	// action
	result := unit.Contains(testData)
	// verify
	test.Equals(t, false, result)
}
