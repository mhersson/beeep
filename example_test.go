package beeep

func ExampleBeep() {
	Beep(DefaultFreq, DefaultDuration)
}

func ExampleNotify() {
	Notify("Title", "MessageBody", "assets/information.png", 2)
}

func ExampleAlert() {
	Alert("Title", "MessageBody", "assets/warning.png", 2)
}
