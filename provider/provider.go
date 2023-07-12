package provider

import (
	"github.com/TurboHsu/munager/provider/netease"
	"github.com/TurboHsu/munager/provider/structure"
)

type Provider int

const (
	Netease Provider = iota
)

func FromString(s string) Provider {
	switch s {
	case "netease":
		return Netease
	default:
		return Netease
	}
}

func (p Provider) SearchSong(keyword string, limit int) ([]structure.SongDetail, error) {
	switch p {
	case Netease:
		return netease.SearchSong(keyword, limit)
	default:
		return netease.SearchSong(keyword, limit)
	}
}

func (p Provider) SearchLyric(song structure.SongDetail) (structure.LyricDetail, error) {
	switch p {
	case Netease:
		return netease.SearchLyric(song)
	default:
		return netease.SearchLyric(song)
	}
}