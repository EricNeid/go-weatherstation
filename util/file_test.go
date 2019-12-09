package util

import "testing"

import "github.com/EricNeid/go-weatherstation/internal/test"

func TestIsFilePresent_shouldReturnTrue(t *testing.T) {
	// action
	result := IsFilePresent("../testdata/img-1.png")
	// verify
	test.Equals(t, true, result)
}

func TestIsFilePresent_shouldReturnFalse(t *testing.T) {
	// action
	result := IsFilePresent("../testdata/no-file")
	// verify
	test.Equals(t, false, result)
}

func TestIsFilePresent_shouldReturnFalseBecauseDirectory(t *testing.T) {
	// action
	result := IsFilePresent("../testdata/dir")
	// verify
	test.Equals(t, false, result)
}
