package tencent

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/TurboHsu/munager/provider/structure"
	"github.com/TurboHsu/munager/util/logging"
	"github.com/TurboHsu/munager/util/network"
)

const (
	searchAPI = "https://c.y.qq.com/soso/fcgi-bin/search_cp?format=json&platform=yqq&new_json=1"
	lyricAPI  = "https://c.y.qq.com/lyric/fcgi-bin/fcg_query_lyric_new.fcg?platform=yqq&g_tk=5381"
)

func SearchSong(keyword string, quantity int) (result []structure.SongDetail, searchErr error) {
	params := fmt.Sprintf("&w=%s&n=%d&p=1", url.QueryEscape(keyword), quantity)
	respRaw, err := network.DoHTTPGET(searchAPI + params)
	logging.HandleErr(err)

	var response SongSearchResult
	logging.HandleErr(json.Unmarshal(respRaw, &response))

	if response.Code != 0 {
		return nil, fmt.Errorf("search failed: %s", response.Message)
	}

	for _, song := range response.Data.Song.List {
		var songDetail structure.SongDetail
		songDetail.SongID = song.Songmid
		songDetail.SongName = song.Songname
		for _, singer := range song.Singer {
			songDetail.ArtistName += singer.Name + " "
		}
		songDetail.Platform = structure.TencentPlatform
		songDetail.AlbumName = song.Albumname

		result = append(result, songDetail)
	}

	return
}
