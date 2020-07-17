package util

import (
	"testing"

	"github.com/EricNeid/go-weatherstation/internal/verify"
)

func TestIsFilePresent_shouldReturnTrue(t *testing.T) {
	// action
	result := IsFilePresent("../../test/testdata/img-1.png")
	// verify
	verify.Equals(t, true, result)
}

func TestIsFilePresent_shouldReturnFalse(t *testing.T) {
	// action
	result := IsFilePresent("../../test/testdata/no-file")
	// verify
	verify.Equals(t, false, result)
}

func TestIsFilePresent_shouldReturnFalseBecauseDirectory(t *testing.T) {
	// action
	result := IsFilePresent("../../test/testdata/dir")
	// verify
	verify.Equals(t, false, result)
}
