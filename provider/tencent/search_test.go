package tencent

import "testing"

func TestSearch(t *testing.T) {
	_, err := SearchSong("Waiting for love", 5)
	if err != nil {
		t.Error(err)
	}
}
