package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	//"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	//"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
	"os"
)

func loadFileAsBytes(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error loading image:", err)
		return nil
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading image file:", err)
		return nil
	}
	return bytes
}

// MakeControls creates the core music controls (UI only for now)
func MakeControls(app fyne.App, window fyne.Window) fyne.CanvasObject {
	// Buttons
	imagePath := "assets/blank_mp3.png"
	imageResource := fyne.NewStaticResource("AlbumArt", loadFileAsBytes(imagePath))
	albumArt := canvas.NewImageFromResource(imageResource)
	albumArt.SetMinSize(fyne.NewSize(500, 500))
	playBtn := widget.NewButton("Play", func() {
		// Hook into backend later
		dialog.NewInformation("Play", "Play button clicked", window).Show()
	})

	pauseBtn := widget.NewButton("Pause", func() {
		dialog.NewInformation("Pause", "Pause button clicked", window).Show()
	})

	stopBtn := widget.NewButton("Stop", func() {
		dialog.NewInformation("Stop", "Stop button clicked", window).Show()
	})
	reverseBtn := widget.NewButton("⏮", func() {
		dialog.NewInformation("Back", "Previous track clicked", window).Show()
	})

	forwardBtn := widget.NewButton("⏭", func() {
		dialog.NewInformation("Next", "Next track clicked", window).Show()
	})

	// Volume Slider
	volumeSlider := widget.NewSlider(0, 100)
	volumeSlider.Value = 50
	volumeSlider.Step = 1
	volumeSlider.OnChanged = func(v float64) {
		// Hook into volume backend later
	}

	volumeBox := container.NewVBox(
		widget.NewLabel("Volume"),
		volumeSlider,
	)

	// Controls layout
	controlRow := container.NewHBox(
		reverseBtn,
		playBtn,
		pauseBtn,
		stopBtn,
		forwardBtn,
	)

	// Playlist name entry field
	playlistName := widget.NewEntry()
	playlistName.SetPlaceHolder("Enter new playlist name...")

	addToPlaylistBtn := widget.NewButton("Add to Playlist", func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil || reader == nil {
				return
			}
			// Placeholder: print file path
			fmt.Println("Selected file:", reader.URI().Path())
		}, window)
	})

	// Simulated list of playlists (replace with dynamic logic later)
	playlistList := []string{"Chill Vibes", "Workout", "Gaming"}

	playlistSelect := widget.NewSelect(playlistList, func(selected string) {
		fmt.Println("Selected playlist:", selected)
	})
	playlistSelect.PlaceHolder = "Select Playlist"

	// Pack all playlist UI into a vertical box
	playlistBox := container.NewVBox(
		widget.NewLabel("Playlists"),
		playlistName,
		addToPlaylistBtn,
		playlistSelect,
	)

	themeToggle := MakeThemeToggle(app)

	// Wrap all in one column
	return container.NewVBox(
		albumArt,
		controlRow,
		volumeBox,
		themeToggle,
		playlistBox,
	)
}

