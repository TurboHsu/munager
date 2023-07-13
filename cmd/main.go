package cmd

import (
	"github.com/TurboHsu/munager/cmd/lyric"
	"github.com/TurboHsu/munager/cmd/sync"
	"github.com/spf13/cobra"
)

const Version = "0.1"

var rootCmd = &cobra.Command{
	Use:   "munager",
	Short: "A handful helper for music management",
	Long: `
	 __  __                                   
	|  \/  |_   _ _ __   __ _  __ _  ___ _ __ 
	| |\/| | | | | '_ \ / _` + "`" + ` |/ _` + "`" + ` |/ _ \ '__|
	| |  | | |_| | | | | (_| | (_| |  __/ |   
	|_|  |_|\__,_|_| |_|\__,_|\__, |\___|_|   
	                          |___/           
	`,
	Version: Version,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// cobra.OnInitialize() should load some config here?
	rootCmd.AddCommand(sync.SyncRoot)
	rootCmd.AddCommand(lyric.LyricRoot)
}
