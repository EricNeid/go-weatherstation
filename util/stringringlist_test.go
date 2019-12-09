package util

import "testing"

import "github.com/EricNeid/go-netconvert/internal/test"

func TestNext(t *testing.T) {
	// arrange
	unit := StringRingList{
		Items: []string{"a", "b"},
	}

	// action & verify
	result, _ := unit.Next()
	test.Equals(t, "a", result)
	result, _ = unit.Next()
	test.Equals(t, "b", result)
	result, _ = unit.Next()
	test.Equals(t, "a", result)
}

func TestNewFileRingList(t *testing.T) {
	// action
	result := NewFileRingList("../testdata")

	// verify
	test.Equals(t, 2, len(result.Items))
}
