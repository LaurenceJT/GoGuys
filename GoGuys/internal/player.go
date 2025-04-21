package internal

import (
	"log"
	"math"
	"os"
	"path/filepath"
	"time"

	"github.com/gopxl/beep"

	"github.com/gopxl/beep/effects"
	"github.com/gopxl/beep/mp3"

	"github.com/gopxl/beep/speaker"
	//"github.com/gopxl/beep/v2/effects"
	//"github.com/gopxl/beep/v2/speaker"
)

var song string
var trackIndexer int = 0
var songsRan int = 0
var starterFormat beep.Format
var format beep.Format

// var streamer beep.StreamSeekCloser
var ctrl *beep.Ctrl
var vol *effects.Volume

// var seq beep.Streamer
var raw beep.StreamSeekCloser
var resamp beep.Streamer

// type Player struct {
//     trackIndexer    int = 0
//     songsRan      int = 0
//     starterFormat beep.Format
//     format        beep.Format
//     streamer      beep.StreamSeekCloser
//     ctrl          *beep.Ctrl
//     vol           *effects.Volume
// }

type SongEntry struct {
	name string
}
type Playlist struct {
	Name     string
	SongList []SongEntry
}

func NewPlaylist(name string) *Playlist {
	return &Playlist{
		Name:     name,
		SongList: PlaylistParser(name),
	}
}

// makes a playlist folder
func PlaylistMake(name string) {
	path := filepath.Join(".", "playlists", name)
	os.Mkdir(path, 0755)
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
		// playlise.S = PlaylistParser()(playlist.name)

		playlistListShow = append(playlistListShow, *playlist)
	}

	return playlistListShow
}

// function to turn playlist into strings send to whatever speaker is playing
func PlaylistParser(playlistName string) []SongEntry {

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

	tracks := playlist.SongList
	song = tracks[trackIndexer].name

	//opens file
	file, err := os.Open(song)

	if err != nil {
		log.Printf("File: %s not found\n", song)

	}

	//decoding
	raw, format, err = mp3.Decode(file)
	if err != nil {
		log.Println("Mp3 decode failed")
	}

	//initialize speaker
	if songsRan == 0 {
		resamp = raw
		starterFormat = format
		speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
		songsRan++
	} else {
		resamp = beep.Resample(4, format.SampleRate, starterFormat.SampleRate, raw)
	}

	ctrl = &beep.Ctrl{Streamer: resamp}
	vol = &effects.Volume{Streamer: ctrl, Base: 2, Volume: 0, Silent: false}

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
				nextSeq, _ := PlaylistPlay(playlist)

				speaker.Play(nextSeq)
			}

		}))
	return seq, done

}

/////////// PLAYBACK FUNCTIONS ///////////////

func Play(b beep.Streamer) {
	speaker.Play(b)
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

}

func Backwards() {
	if trackIndexer > 0 {
		trackIndexer--
	}

}
