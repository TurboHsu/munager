package netease

import (
	"encoding/json"
	"fmt"

	"github.com/TurboHsu/munager/provider/structure"
	"github.com/TurboHsu/munager/util/logging"
	"github.com/TurboHsu/munager/util/network"
)

const (
	searchAPI = "https://music.163.com/weapi/search/get"
)

// SearchSong search song by key words
func SearchSong(key string, quantity int) (result []structure.SongDetail, searchErr error) {
	// Param details:
	// s: key words
	// type: 1: Song, 10: Album, 100: Singer, 1000: Playlist, 1002: User,
	// 1004: MV, 1006: Lyric, 1009: Radio, 1014: Vids
	// offset: offset
	// limit: limit
	data := map[string]interface{}{
		"s":      key,
		"type":   1,
		"offset": 0,
		"limit":  quantity,
	}
	var httpHeaders = [][]string{
		{"Origin", "https://music.163.com"},
		{"Referer", "https://music.163.com"},
	}
	var response NeteaseSearchResult
	respRaw, err := network.DoHTTPPostWithHeaders(searchAPI,
		convertParams(data, WEAPI), httpHeaders)
	logging.HandleErr(err)
	logging.HandleErr(json.Unmarshal(respRaw, &response))
	if response.Code != 200 {
		searchErr = fmt.Errorf("search [%s] failed, code: %d", key, response.Code)
		return
	}
	if len(response.Result.Songs) == 0 {
		searchErr = fmt.Errorf("search failed, no result returned")
		return
	}
	for _, song := range response.Result.Songs {
		var artist string
		for _, a := range song.Artists {
			artist += a.Name + " "
		}
		result = append(result, structure.SongDetail{
			SongID:     song.ID,
			SongName:   song.Name,
			ArtistName: artist,
			AlbumName:  song.Album.Name,
			Platform:   structure.NeteasePlatform,
		})
	}
	return
}
