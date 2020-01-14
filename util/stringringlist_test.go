package util

import "testing"
import "EricNeid/go-weatherstation/internal/verify"

func TestNext(t *testing.T) {
	// arrange
	unit := StringRingList{
		Items: []string{"a", "b"},
	}

	// action & verify
	result, _ := unit.Next()
	verify.Equals(t, "a", result)
	result, _ = unit.Next()
	verify.Equals(t, "b", result)
	result, _ = unit.Next()
	verify.Equals(t, "a", result)
}

func TestNewFileRingList(t *testing.T) {
	// action
	result := NewFileRingList("../testdata")

	// verify
	verify.Equals(t, 2, len(result.Items))
}
