package file

import (
	"path/filepath"
)

func IsAudio(path string) bool {
	ext := filepath.Ext(path)
	return IsExtAudio(ext)
}

func IsExtAudio(ext string) bool {
	switch ext {
	case ".mp3", ".flac", ".wav", ".m4a", ".ogg", ".opus":
		return true
	default:
		return false
	}
}

func IsExtLyric(ext string) bool {
	switch ext {
	case ".lrc":
		return true
	default:
		return false
	}
}

func IsExtImage(ext string) bool {
	switch ext {
	case ".jpg", ".png", ".webp":
		return true
	default:
		return false
	}
}
