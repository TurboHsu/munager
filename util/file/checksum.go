package file

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"

	"github.com/TurboHsu/munager/util/logging"
)

func CalculateSHA1(f *os.File) string {
	h := sha1.New()
	_, err := io.Copy(h, f)
	logging.HandleErr(err)
	return fmt.Sprintf("%x", h.Sum(nil))
}
