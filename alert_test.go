package beeep

import (
	"testing"
)

func TestAlert(t *testing.T) {
	err := Alert("Alert title", "Message body", "assets/information.png", 0)
	if err != nil {
		t.Error(err)
	}

	err = Alert("Alert title", "Message body", "assets/warning.png", 1)
	if err != nil {
		t.Error(err)
	}

	err = Alert("Alert title", "Message body", "assets/warning.png", 2)
	if err != nil {
		t.Error(err)
	}
}
