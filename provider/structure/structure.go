package structure

type Platform int

const (
	NeteasePlatform Platform = iota
	TencentPlatform
)

type SongDetail struct {
	SongID     string
	SongName   string
	ArtistName string
	AlbumName  string
	Platform   Platform
}

type LyricDetail struct {
	RawLyric        []LyricLine
	TranslatedLyric []LyricLine
}

type LyricLine struct {
	Time struct {
		Minute      int
		Second      int
		Microsecond int
	}
	Lyric string
}
