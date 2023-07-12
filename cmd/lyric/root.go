package lyric

import "github.com/spf13/cobra"

var LyricRoot = &cobra.Command{
	Use:   "lyric",
	Short: "Lyric manager",
	Long:  "Munager lyric is a utility for automatic lyric download and management.",
}

func init() {
	LyricRoot.AddCommand(Fetch)
	LyricRoot.AddCommand(Query)
}
