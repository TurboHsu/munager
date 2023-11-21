package provider

import (
	"errors"

	"github.com/TurboHsu/munager/provider/netease"
	"github.com/TurboHsu/munager/provider/structure"
	"github.com/TurboHsu/munager/provider/tencent"
	"github.com/TurboHsu/munager/util/logging"
)

type Provider int

const (
	Netease Provider = iota
	Tencent
)

func FromString(s string) Provider {
	switch s {
	case "netease":
		return Netease
	case "tencent":
		return Tencent
	default:
		logging.Info("Unknown provider, use Netease as default")
		return Netease
	}
}

func (p Provider) SearchSong(keyword string, limit int) ([]structure.SongDetail, error) {
	switch p {
	case Netease:
		return netease.SearchSong(keyword, limit)
	case Tencent:
		return tencent.SearchSong(keyword, limit)
	default:
		err := errors.New("unexpected provider")
		return nil, err
	}
}

func (p Provider) SearchLyric(song structure.SongDetail) (structure.LyricDetail, error) {
	switch p {
	case Netease:
		return netease.SearchLyric(song)
	case Tencent:
		return tencent.FetchLyric(song)
	default:
		err := errors.New("unexpected provider")
		return structure.LyricDetail{}, err
	}
}
