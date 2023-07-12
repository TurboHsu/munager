package netease

import (
	"testing"
)

func TestSearchID(t *testing.T) {
	_, err := SearchSong("Waiting for love", 5)
	if err != nil {
		t.Error(err)
	}
}

func TestGettingLyric(t *testing.T) {
	song, err := SearchSong("secret base", 1)
	if err != nil {
		t.Error(err)
	}
	lrc, err := SearchLyric(song[0])
	if err != nil {
		t.Error(err)
	}
	if len(lrc.RawLyric) == 0 {
		t.Error("No lyric found for song \"secret base\", which is bad")
	}
}
