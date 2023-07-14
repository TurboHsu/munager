package lyric

import (
	"fmt"
	"os"
	"strings"

	"github.com/TurboHsu/munager/provider"
	"github.com/TurboHsu/munager/util/logging"
	lyricformatter "github.com/TurboHsu/munager/util/lyric_formatter"
	"github.com/spf13/cobra"
)

var Query = &cobra.Command{
	Use:   "query",
	Short: "Query lyrics",
	Long: `Query lyrics for a given keyword.
You can save the result to a file by using the -o flag.`,
	Example: `munager lyric query -o "Waiting for Love.lrc" -f raw-only Waiting for Love`,
}

func RunQuery(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		logging.HandleErr(fmt.Errorf("no keyword specified, please specify some keywords in order to search something. For more information, please use the --help flag"))
		return
	}

	overwrite, err := Query.Flags().GetBool("overwrite")
	if err != nil {
		logging.HandleErr(err)
		return
	}

	p := provider.FromString(Query.Flag("provider").Value.String())
	song, err := p.SearchSong(strings.Join(args, " "), 1)
	if err != nil {
		logging.HandleErr(err)
		return
	}
	lyric, err := p.SearchLyric(song[0])
	logging.HandleErr(err)

	silent, err := Query.Flags().GetBool("silent")
	logging.HandleErr(err)

	formatter := lyricformatter.FromString(Query.Flag("format").Value.String())
	formatted := formatter.FormatLyric(&lyric)

	if Query.Flag("output").Value.String() != "" {
		// Find whether the file exists
		if _, err := os.Stat(Query.Flag("output").Value.String()); err == nil && !overwrite {
			// File exists
			logging.HandleErr(fmt.Errorf("file %s already exists, skipping", Query.Flag("output").Value.String()))
			return
		}
		// File doesn't exist
		f, err := os.Create(Query.Flag("output").Value.String())
		logging.HandleErr(err)
		defer f.Close()
		f.WriteString(formatted)
	}

	if !silent {
		fmt.Println(formatted)
	}
}

func init() {
	Query.Flags().StringP("provider", "p", "netease", "Specify a lyric provider")
	Query.Flags().StringP("output", "o", "", "Specify a file to save the result")
	Query.Flags().BoolP("silent", "s", false, "Don't print the result to stdout")
	Query.Flags().BoolP("overwrite", "O", false, "Overwrite existing lyric files")
	Query.Run = RunQuery
	appendFormattingFlags(Query)
}
