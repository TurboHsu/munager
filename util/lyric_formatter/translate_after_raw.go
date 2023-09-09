package lyricformatter

import (
	"github.com/TurboHsu/munager/provider/structure"
)

func translateAfterRaw(lyric *structure.LyricDetail) string {
	raw := make(map[int]string)
	translated := make(map[int]string)
	var timestamp []int
	for _, line := range lyric.RawLyric {
		time := calcTime(&line)
		timestamp = append(timestamp, time)
		raw[time] = summonLRCTimestamp(&line) + line.Lyric
	}

	for _, line := range lyric.TranslatedLyric {
		if line.Lyric != "" {
			translated[calcTime(&line)] = summonLRCTimestamp(&line) + "「" + line.Lyric + "」"
		} else {
			translated[calcTime(&line)] = summonLRCTimestamp(&line) + line.Lyric
		}
	}

	// Merge two maps
	var ret string
	for _, time := range timestamp {
		ret += raw[time]
		if translated[time] != "" {
			ret += "\n" + translated[time]
		}
		ret += "\n"
	}
	return ret
}
