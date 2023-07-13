package server

import (
	"github.com/TurboHsu/munager/cmd/sync/structure"
	"github.com/TurboHsu/munager/util/file"
)

func getFiles(path string) (f []structure.FileInfo) {
	p, e := file.ListAllFilesSplit(path)
	for i := 0; i < len(p); i++ {
		f = append(f, structure.FileInfo{
			PathBase:  p[i],
			Extension: e[i],
		})
	}
	return
}
