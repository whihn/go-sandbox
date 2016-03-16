package main

import (
	"../stringutil"
	"testing"
)

func TestShouldFail(t *testing.T) {
	result := stringutil.Reverse("foo")
	if result != "foo" {
		t.Error("expected " + result + " to be foo")
	}
}
