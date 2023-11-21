package tencent

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/TurboHsu/munager/provider/structure"
	"github.com/TurboHsu/munager/util/logging"
	"github.com/TurboHsu/munager/util/network"
)

func FetchLyric(song structure.SongDetail) (result structure.LyricDetail, fetchErr error) {
	params := fmt.Sprintf("&songmid=%s&nobase64=0", song.SongID)
	headers := [][]string{
		{"Referer", "https://y.qq.com/portal/player.html"},
	}

	respRaw, err := network.DoHTTPGETWithHeaders(lyricAPI+params, headers)
	logging.HandleErr(err)

	// Trim respRaw
	if len(respRaw) < 18 {
		return result, fmt.Errorf("fetch lyric failed: tencent lyric return too short: %s", respRaw)
	}
	respRaw = respRaw[18 : len(respRaw)-1]

	var response LyricFetchResult

	logging.HandleErr(json.Unmarshal(respRaw, &response))

	if response.Code != 0 {
		return result, fmt.Errorf("fetch lyric failed: tencent lyric ret code %d", response.Code)
	}

	decodedLyric, err := base64.StdEncoding.DecodeString(response.Lyric)
	logging.HandleErr(err)
	if err != nil {
		return result, err
	}

	decodedTranslate, err := base64.StdEncoding.DecodeString(response.Trans)
	logging.HandleErr(err)
	if err != nil {
		return result, err
	}

	if len(decodedLyric) > 0 {
		result.RawLyric = parseLyric(string(decodedLyric))
	}
	if len(decodedTranslate) > 0 {
		result.TranslatedLyric = parseLyric(string(decodedTranslate))
	}

	return
}

func parseLyric(raw string) (ret []structure.LyricLine) {
	lines := strings.Split(raw, "\n")

	for _, line := range lines {
		// escape too short line
		if len(line) < 10 {
			continue
		}

		if doesContainCopyrightStatement(line) {
			continue
		}

		var newLine structure.LyricLine

		if len(line) > 10 && line[0] == '[' && line[3] == ':' && line[6] == '.' && line[10] == ']' &&
			isNumber(line[1:3]) && isNumber(line[4:6]) && isNumber(line[7:10]) {
			// 3 digit millisecond timestamp
			var minute, second, microsecond int
			fmt.Sscanf(line, "[%d:%d.%d]", &minute, &second, &microsecond)
			newLine.Time.Minute = minute
			newLine.Time.Second = second
			newLine.Time.Microsecond = microsecond
			newLine.Lyric = line[11:]
			ret = append(ret, newLine)
		} else if len(line) > 9 && line[0] == '[' && line[3] == ':' && line[6] == '.' && line[9] == ']' &&
			isNumber(line[1:3]) && isNumber(line[4:6]) && isNumber(line[7:9]) {
			// 2 digit millisecond timestamp
			var minute, second, microsecond int
			fmt.Sscanf(line, "[%d:%d.%d]", &minute, &second, &microsecond)
			newLine.Time.Minute = minute
			newLine.Time.Second = second
			newLine.Time.Microsecond = microsecond * 10
			newLine.Lyric = line[10:]
			ret = append(ret, newLine)
		}
	}
	return
}

func isNumber(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func doesContainCopyrightStatement(s string) bool {
	return strings.Contains(s, "QQ音乐")
}
