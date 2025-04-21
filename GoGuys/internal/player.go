package internal

import (
	"log"
	"math"
	"os"
	"path/filepath"
	"time"

	"github.com/gopxl/beep"
	//"github.com/gopxl/beep/effects"
	"github.com/gopxl/beep/mp3"

	//"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/v2/effects"
	"github.com/gopxl/beep/v2/speaker"
	// "github.com/gopxl/beep/mp3"
	// "github.com/gopxl/beep/speaker"
)

var song string
var trackIndexer int = 0
var songsRan int = 0
var starterFormat beep.Format
var format beep.Format
var streamer beep.StreamSeekCloser
var ctrl *beep.Ctrl
var vol *effects.Volume

// type Player struct {
//     trackIndexer    int = 0
//     songsRan      int = 0
//     starterFormat beep.Format
//     format        beep.Format
//     streamer      beep.StreamSeekCloser
//     ctrl          *beep.Ctrl
//     vol           *effects.Volume
// }

type Playlist struct {
	name     string
	songList []songEntry
}

func NewPlaylist(name string) *Playlist {
	return &Playlist{
		Name:     name,
		SongList: playlistParser(name),
	}
}

type SongEntry struct {
	name string
}

// makes a playlist folder
func PlaylistMake(name string) {
	path := filepath.Join(".", "playlists", name)
	os.Mkdir(path, 0755)
}

// removes the playlist
func playlistDelete(name string) {
	path := filepath.Join(".", "playlists", name)
	os.Remove(path)
}

func PlaylistList() []Playlist {
	path := filepath.Join(".", "playlists")

	playlistList, err := os.ReadDir(path)
	if err != nil {
		return nil
	}
	//list of playlists holder
	var playlistListShow []Playlist

	//for loop that parses through a playlist list
	for entryIndex := range playlistList {

		playlist := NewPlaylist(playlistList[entryIndex].Name())
		// playlist.name = playlistList[entryIndex].Name()
		// playlise.songList = playlistParser(playlist.name)

		playlistListShow = append(playlistListShow, *playlist)
	}
}

// function to turn playlist into strings send to whatever speaker is playing
func playlistParser(playlistName string) []SongEntry {

	path := filepath.Join(".", "playlists", playlistName)

	playlistEntries, err := os.ReadDir(path)
	if err != nil {
		return nil
	}

	//stores parsed songs for playing
	var playlistToPlay []SongEntry

	//for loop that parses through a playlist
	for entryIndex := range playlistEntries {

		var songEntry SongEntry
		songEntry.name = playlistEntries[entryIndex].Name()

		playlistToPlay = append(playlistToPlay, songEntry)
	}

	return playlistToPlay
}

// plays playlist
func PlaylistPlay(playlist Playlist) (beep.Streamer, chan struct{}) {

	tracks := playlist.songList
	song = tracks[trackIndexer].name

	//opens file
	file, err := os.Open(song)

	if err != nil {
		log.Println("File: %s not found", song)
	}

	//decoding
	streamer, format, err = mp3.Decode(file)
	if err != nil {
		log.Println("Mp3 decode failed")
	}

	//initialize speaker
	if songsRan == 0 {
		starterFormat = format
		speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
		songsRan++
	} else {
		resampled := beep.Resample(4, starterFormat.SampleRate, format.SampleRate, streamer)

		streamer = resampled
	}

	ctrl = &beep.Ctrl{Streamer: streamer}
	vol = &effects.Volume{Streamer: ctrl, Base: 2, Volume: 0, Slient: false}

	speaker.Clear()

	done := make(chan struct{})
	seq := beep.Seq(vol,
		beep.Callback(func() {

			//closes at the end of the playlist or autoplays until the end
			if trackIndexer >= len(tracks) {
				close(done)
				return
			} else {
				trackIndexer++
				speaker.Play(playlistPlay(Playlist))
			}

		}))
	return seq, done

}

/////////// PLAYBACK FUNCTIONS ///////////////

func Play() {
	speaker.Play(seq)
}

func Pause() {

	//allows for pause
	speaker.Lock()
	ctrl.Paused = !ctrl.Paused
	speaker.Unlock()
}

func Volume(v float64) {
	vol.Volume = math.Log2(v / 100)
}

func Stop() {
	speaker.Clear()
	trackIndexer = 0
}

func Forward() {
	trackIndexer++
	speaker.Play(seq)
}

func Backwards() {
	if trackIndexer > 0 {
		trackIndexer--
	}
	speaker.Play(seq)
}
