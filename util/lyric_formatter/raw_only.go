package lyricformatter

import "github.com/TurboHsu/munager/provider/structure"

func rawOnly(lyric *structure.LyricDetail) string {
	var ret string
	for _, line := range lyric.RawLyric {
		ret += summonLRCTimestamp(&line) + line.Lyric + "\n"
	}
	return ret
}
