package lyricformatter

import "github.com/TurboHsu/munager/provider/structure"

func translateOnly(lyric *structure.LyricDetail) string {
	var ret string
	for _, line := range lyric.TranslatedLyric {
		ret += summonLRCTimestamp(&line) + line.Lyric + "\n"
	}
	return ret
}
