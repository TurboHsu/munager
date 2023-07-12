package file

import "path/filepath"

func IsAudio(path string) bool {
	ext := filepath.Ext(path)
	switch ext {
	case ".mp3", ".flac", ".wav", ".m4a", ".ogg", ".opus":
		return true
	default:
		return false
	}
}
