package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"os"
)

func greeting() {
	a := app.New()
	w := a.NewWindow("Hello World")

	var btn = widget.NewButton("Поехали!", func() {
		os.Exit(1)
	})

	w.SetContent(btn)
	w.ShowAndRun()

}
