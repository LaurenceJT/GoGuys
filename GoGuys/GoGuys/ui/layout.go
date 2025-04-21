package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// MakeMainLayout builds the overall screen layout
func MakeMainLayout(controls fyne.CanvasObject) fyne.CanvasObject {
	// Placeholder: later we’ll add track info & playlist sections
	mainContent := container.NewVBox(
		controls,
	)

	return container.NewBorder(nil, nil, nil, nil, mainContent)
}

// DefaultWindowSize returns default app window size
func DefaultWindowSize() fyne.Size {
	return fyne.NewSize(800, 500)
}
