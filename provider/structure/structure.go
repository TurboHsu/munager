package structure

type Platform int

const (
	NeteasePlatform Platform = iota
)

type SongDetail struct {
	SongID     int
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
