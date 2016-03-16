package main 

import (
	"testing"
	"../stringutil"
)

func TestShouldFail(t *testing.T) {
	result := stringutil.Reverse("foo")
	if result != "foo" {
		t.Error("expected " + result + " to be foo")
	}
}