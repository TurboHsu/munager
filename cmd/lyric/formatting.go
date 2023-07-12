package lyric

import "github.com/spf13/cobra"

func appendFormattingFlags(target *cobra.Command) {
	target.Flags().StringP("format", "f", "translate-after-raw",
	 "The way to format the lyric, there are several options: raw-only, translate-only, translate-after-raw")
}