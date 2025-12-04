package tug_test

import (
	"github.com/mcandre/tug"

	"testing"
)

func TestPlatformMarshalingSymmetric(t *testing.T) {
	platformString := "linux/arm64"
	platform, err := tug.ParsePlatform(platformString)

	if err != nil {
		t.Error(err)
	}

	platformString2 := platform.String()

	if platformString2 != platformString {
		t.Errorf("expected symmetric marshaling for platform %v", platformString)
	}
}
