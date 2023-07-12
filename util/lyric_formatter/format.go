package lyricformatter

import (
	"strconv"

	"github.com/TurboHsu/munager/provider/structure"
)

type Format int

const (
	TranslateAfterRaw Format = iota
	RawOnly
	TranslateOnly
)

func FromString(s string) Format {
	switch s {
	case "translate-after-raw":
		return TranslateAfterRaw
	case "raw-only":
		return RawOnly
	case "translate-only":
		return TranslateOnly
	default:
		return TranslateAfterRaw
	}
}

func (f Format) FormatLyric(lyric *structure.LyricDetail) string {
	switch f {
	case TranslateAfterRaw:
		return translateAfterRaw(lyric)
	case RawOnly:
		return rawOnly(lyric)
	case TranslateOnly:
		return translateOnly(lyric)
	default:
		return translateAfterRaw(lyric)
	}
}

// calcTime calculates the time in microseconds
func calcTime(line *structure.LyricLine) int {
	return line.Time.Minute*60000000 + line.Time.Second*1000000 + line.Time.Microsecond
}

func summonLRCTimestamp(line *structure.LyricLine) string {
	var min, sec, microsec string
	if line.Time.Minute < 10 {
		min = "0" + strconv.Itoa(line.Time.Minute)
	} else {
		min = strconv.Itoa(line.Time.Minute)
	}
	if line.Time.Second < 10 {
		sec = "0" + strconv.Itoa(line.Time.Second)
	} else {
		sec = strconv.Itoa(line.Time.Second)
	}
	if line.Time.Microsecond < 10 {
		microsec = "00" + strconv.Itoa(line.Time.Microsecond)
	} else if line.Time.Microsecond < 100 {
		microsec = "0" + strconv.Itoa(line.Time.Microsecond)
	} else {
		microsec = strconv.Itoa(line.Time.Microsecond)
	}
	return "[" + min + ":" + sec + "." + microsec + "]"
}
