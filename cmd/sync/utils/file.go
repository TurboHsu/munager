package utils

import (
	"github.com/TurboHsu/munager/cmd/sync/structure"
	"github.com/TurboHsu/munager/util/file"
	"github.com/TurboHsu/munager/util/logging"
	"github.com/spf13/cobra"
)

func FilterValidFiles(f []structure.FileInfo) []structure.FileInfo {
	var ret []structure.FileInfo
	for _, n := range f {
		if isAllowedExtension("." + n.Extension) {
			ret = append(ret, n)
		}
	}
	return ret
}

func isAllowedExtension(ext string) bool {
	return file.IsExtAudio(ext) || file.IsExtLyric(ext) || file.IsExtImage(ext)
}

func GetFiles(path string) (f []structure.FileInfo) {
	p, e := file.ListAllFilesSplit(path)
	for i := 0; i < len(p); i++ {
		f = append(f, structure.FileInfo{
			PathBase:  p[i],
			Extension: e[i],
		})
	}
	return
}

func FixPath(cmd *cobra.Command) {
	// Check if path have slash in the end
	path, err := cmd.Flags().GetString("path")
	logging.HandleErr(err)
	if len(path) == 0 {
		cmd.Flag("path").Value.Set("./")
	} else if path[len(path)-1] != '/' {
		cmd.Flag("path").Value.Set(path + "/")
	}
}
