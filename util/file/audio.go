package file

import (
	"os"

	"github.com/TurboHsu/munager/structure"
	"github.com/TurboHsu/munager/util/logging"
	"github.com/dhowden/tag"
)

func ReadMetadata(path string) (ret structure.Song, err error) {
	f, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		logging.HandleErr(err)
		return
	}
	defer f.Close()
	m, err := tag.ReadFrom(f)
	if err != nil {
		logging.HandleErr(err)
		return
	}
	ret = structure.Song{
		Title:  m.Title(),
		Artist: m.Artist(),
	}
	return
}
