package netease

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/TurboHsu/munager/provider/structure"
	"github.com/TurboHsu/munager/util/logging"
	"github.com/TurboHsu/munager/util/network"
)

const (
	LinuxAPIURL = "https://music.163.com/api/linux/forward"
)

// Uses LinuxAPI to fetch song lyric
// IDK why Muget uses LinuxAPI to do so, but it works anyways
func SearchLyric(song structure.SongDetail) (structure.LyricDetail, error) {
	data := map[string]interface{}{
		"method": "POST",
		"url":    "https://music.163.com/api/song/lyric?lv=-1&kv=-1&tv=-1",
		"params": map[string]int{
			"id": song.SongID,
		},
	}
	params := convertParams(data, LinuxAPI)
	var httpHeaders = [][]string{
		{"Origin", "https://music.163.com"},
		{"Referer", "https://music.163.com"},
	}
	respRaw, err := network.DoHTTPPostWithHeaders(LinuxAPIURL,
		params, httpHeaders)
	logging.HandleErr(err)

	var resp NeteaseLyricResult
	logging.HandleErr(json.Unmarshal(respRaw, &resp))

	if resp.Code != 200 {
		return structure.LyricDetail{}, fmt.Errorf("search lyric failed, code: %d", resp.Code)
	}

	// Parses the result
	var ret structure.LyricDetail
	if len(resp.Lrc.Lyric) != 0 {
		ret.RawLyric = parseLyric(resp.Lrc.Lyric)
	}
	if len(resp.Tlyric.Lyric) != 0 {
		ret.TranslatedLyric = parseLyric(resp.Tlyric.Lyric)
	}
	return ret, nil
}

func parseLyric(rawLyric string) (ret []structure.LyricLine) {
	lines := strings.Split(rawLyric, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		var lyricLine structure.LyricLine
		// Checks whether the line is a timestamp

		if len(line) > 10 && line[0] == '[' && line[3] == ':' && line[6] == '.' && line[10] == ']' {
			// 3 digit millisecond timestamp
			var minute, second, microsecond int
			fmt.Sscanf(line, "[%d:%d.%d]", &minute, &second, &microsecond)
			lyricLine.Time.Minute = minute
			lyricLine.Time.Second = second
			lyricLine.Time.Microsecond = microsecond
			lyricLine.Lyric = line[11:]
			ret = append(ret, lyricLine)
		} else if len(line) > 9 && line[0] == '[' && line[3] == ':' && line[6] == '.' && line[9] == ']' {
			// 2 digit millisecond timestamp
			var minute, second, microsecond int
			fmt.Sscanf(line, "[%d:%d.%d]", &minute, &second, &microsecond)
			lyricLine.Time.Minute = minute
			lyricLine.Time.Second = second
			lyricLine.Time.Microsecond = microsecond * 10
			lyricLine.Lyric = line[10:]
			ret = append(ret, lyricLine)
		}
	}
	return
}