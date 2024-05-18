package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()
	w := app.NewWindow("Finance Management")

	ui := NewUI(w)
	w.SetContent(ui.MakeUI())
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
