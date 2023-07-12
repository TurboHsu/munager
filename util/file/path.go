package file

import (
	"os"
	"path/filepath"

	"github.com/TurboHsu/munager/util/logging"
)

func ListAllFiles(path string) (ret []string) {
	logging.HandleErr(filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			ret = append(ret, path)
		}
		return nil
	}))
	return
}