package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// MakeThemeToggle returns a toggle button to switch between light and dark themes
func MakeThemeToggle(app fyne.App) fyne.CanvasObject {
	isDark := true // default is dark

	toggle := widget.NewButton("Switch Theme", func() {
		if isDark {
			app.Settings().SetTheme(theme.LightTheme())
		} else {
			app.Settings().SetTheme(theme.DarkTheme())
		}
		isDark = !isDark
	})

	return container.NewVBox(toggle)
}
