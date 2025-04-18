package internal

import "fmt"

// Player is the struct for managing playback (placeholder for backend logic)
type Player struct {
	CurrentTrack string
	Volume       float64
	IsPlaying    bool
}

// NewPlayer returns a new player instance
func NewPlayer() *Player {
	return &Player{
		CurrentTrack: "",
		Volume:       50,
		IsPlaying:    false,
	}
}

// Play simulates playing a track
func (p *Player) Play() {
	p.IsPlaying = true
	fmt.Println("Playing:", p.CurrentTrack)
}

// Pause simulates pausing playback
func (p *Player) Pause() {
	p.IsPlaying = false
	fmt.Println("Paused")
}

// Stop simulates stopping playback
func (p *Player) Stop() {
	p.IsPlaying = false
	fmt.Println("Stopped")
}
