package lyric

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/TurboHsu/munager/provider"
	"github.com/TurboHsu/munager/util/file"
	"github.com/TurboHsu/munager/util/logging"
	lyricformatter "github.com/TurboHsu/munager/util/lyric_formatter"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

var Fetch = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch lyrics",
	Long: `This command will search and download lyrics from selected lyric providers.
By running this command, Munager will search for lyrics of all songs in the current directory, if they don't have a lyric file.`,
}

func RunFetch(cmd *cobra.Command, args []string) {
	fileList := file.ListAllFiles(Fetch.Flag("path").Value.String())
	overwrite, err := Fetch.Flags().GetBool("overwrite")
	logging.HandleErr(err)
	silent, err := Fetch.Flags().GetBool("silent")
	logging.HandleErr(err)
	var queue []string

	if !silent {
		logging.Info("There are " + strconv.Itoa(len(fileList)) + " songs in total.")
	}

	for _, f := range fileList {
		if !file.IsAudio(f) {
			continue
		}
		// Check whether the lyric file exists
		lrcFile := f[:len(f)-len(filepath.Ext(f))] + ".lrc"

		if _, err := os.Stat(lrcFile); err == nil && !overwrite {
			// File exists
			continue
		}
		queue = append(queue, f)
	}

	logging.Info("Found " + strconv.Itoa(len(queue)) + " songs without lyrics")

	// Start routines to do that
	thread, err := Fetch.Flags().GetInt("thread")
	logging.HandleErr(err)
	provider := provider.FromString(Fetch.Flag("provider").Value.String())
	formatter := lyricformatter.FromString(Fetch.Flag("format").Value.String())

	// jobs is a thread controller
	jobs := make(chan bool, thread)
	var wg sync.WaitGroup
	bar := progressbar.Default(int64(len(queue)))

	// TODO: If searching in one provider failed, try another one
	for _, f := range queue {
		jobs <- true
		wg.Add(1)
		go func(f string) {
			metadata, err := file.ReadMetadata(f)
			var keyword string
			if err != nil {
				logging.HandleErr(fmt.Errorf("failed to read metadata for %s, using filename as keywords", f))
				keyword = filepath.Base(f)
			} else {
				keyword = metadata.Title + " " + metadata.Artist
			}

			song, err := provider.SearchSong(keyword, 1)
			if err != nil {
				logging.HandleErr(err)
				bar.Add(1)
				wg.Done()
				<-jobs
				return
			}
			lyric, err := provider.SearchLyric(song[0])
			logging.HandleErr(err)

			formatted := formatter.FormatLyric(&lyric)

			file, err := os.OpenFile(f[:len(f)-len(filepath.Ext(f))]+".lrc", os.O_CREATE|os.O_WRONLY, 0644)
			logging.HandleErr(err)
			defer file.Close()
			file.WriteString(formatted)
			if !silent {
				logging.Info("Fetching lyric for " + f)
			}
			bar.Add(1)
			wg.Done()
			<-jobs
		}(f)
	}

	wg.Wait()
	logging.Info("Done!")
}

func init() {
	Fetch.Flags().BoolP("overwrite", "o", false, "Overwrite existing lyric files")
	Fetch.Flags().StringP("provider", "p", "netease", "Specify a lyric provider")
	Fetch.Flags().StringP("path", "P", ".", "Specify a path to search for songs, single file is also supported")
	Fetch.Flags().IntP("thread", "t", 5, "Specify the number of threads to use")
	Fetch.Flags().BoolP("silent", "s", true, "Don't print detailed information to stdout")
	appendFormattingFlags(Fetch)
	Fetch.Run = RunFetch
}
