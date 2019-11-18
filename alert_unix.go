// +build linux freebsd netbsd openbsd

package beeep

// Alert displays a desktop notification and plays a beep.
func Alert(title, message, appIcon string, urgency int) error {
	if err := Notify(title, message, appIcon, urgency); err != nil {
		return err
	}
	return Beep(DefaultFreq, DefaultDuration)
}
