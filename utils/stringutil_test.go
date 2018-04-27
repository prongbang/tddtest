package utils

import (
	"testing"
)

func TestReverse(t *testing.T) {
	if Reverse("Hello") != "Hello" {
		t.Fail()
	}
}
