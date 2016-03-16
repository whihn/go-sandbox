package stringutil

import (
	"testing"
)

func TestShouldReturnPassedString(t *testing.T) {

	result := reverse("foo")
	if result != "foo" {
		t.Error("failed...")
	}
}