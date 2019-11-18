package beeep

import (
	"testing"
)

func TestNotify(t *testing.T) {
	err := Notify("Notify title", "Message body", "assets/information.png", 0)
	if err != nil {
		t.Error(err)
	}

	err = Notify("Notify title", "Message body", "assets/warning.png", 1)
	if err != nil {
		t.Error(err)
	}

	err = Notify("Notify title", "Message body", "assets/warning.png", 2)
	if err != nil {
		t.Error(err)
	}
}
