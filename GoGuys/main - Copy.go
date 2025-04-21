package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	//"github.com/Jvargas40/GoGuys/internal"//
	"github.com/Jvargas40/GoGuys/ui"
)

// Entry point of the app
func main() {
	myApp := app.NewWithID("com.goguys.player")
	myApp.Settings().SetTheme(theme.DarkTheme()) // Default to dark theme

	window := myApp.NewWindow("GoGuys Music Player")
	window.Resize(ui.DefaultWindowSize())

	// Load controls and layout
	controls := ui.MakeControls(myApp, window)
	// player := internal.NewPlayer() // placeholder, will hook in later
	mainLayout := ui.MakeMainLayout(controls)

	window.SetContent(mainLayout)
	window.ShowAndRun()
}
