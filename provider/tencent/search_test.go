package tencent

import (
	"testing"

	"github.com/TurboHsu/munager/provider/structure"
)

func TestSearch(t *testing.T) {
	v, err := SearchSong("Waiting for love", 5)
	if err != nil {
		t.Error(err)
	}
	if len(v) == 0 {
		t.Error("No result returned")
	}
}

func TestLyricFetch(t *testing.T) {
	v, err := FetchLyric(structure.SongDetail{SongID: "000lXvdX0uOg8F"})
	if err != nil {
		t.Error(err)
	}
	if len(v.RawLyric) == 0 {
		t.Error("No lyric returned")
	}
}
