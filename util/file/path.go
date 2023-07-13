package file

import (
	"os"
	"path/filepath"

	"github.com/TurboHsu/munager/util/logging"
)

func ListAllFiles(rawPath string) (ret []string) {
	logging.HandleErr(filepath.Walk(rawPath, func(path string, info os.FileInfo, err error) error {
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

// This method will return two slices, one for base name, one for extension
func ListAllFilesSplit(rawPath string) (base []string, ext []string) {
	// Use absloute path
	absRaw, err := filepath.Abs(rawPath)
	logging.HandleErr(err)

	// Convert to slash
	absRaw = filepath.ToSlash(absRaw)

	logging.HandleErr(
		filepath.Walk(rawPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				// Use abslote path
				absPath, err := filepath.Abs(path)
				logging.HandleErr(err)

				// Convert to slash
				absPath = filepath.ToSlash(absPath)

				// Trim extension and working dir
				b := absPath[len(absRaw)+1 : len(absPath)-len(filepath.Ext(absPath))]

				base = append(base, b)
				// Trim extension dot
				if filepath.Ext(path) != "" {
					ext = append(ext, filepath.Ext(path)[1:])
				} else {
					ext = append(ext, "")
				}
			}
			return nil
		},
		))

	return
}
