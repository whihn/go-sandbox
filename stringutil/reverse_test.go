package stringutil

import (
	"testing"
)

func TestShouldReturnPassedString(t *testing.T) {

	result := Reverse("foo")
	if result != "foo" {
		t.Error("failed...")
	}
}
