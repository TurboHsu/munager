package cmd

import (
	"github.com/TurboHsu/munager/cmd/lyric"
	"github.com/TurboHsu/munager/cmd/sync"
	"github.com/spf13/cobra"
)

const Version = "0.1.0"

var rootCmd = &cobra.Command{
	Use:   "munager",
	Short: "A handful helper for music management",
	Long: `
▄▄▄  ▄▄▄  
███  ███                               
████████ ▐▌ ▐▌▐▙██▖ ▟██▖ ▟█▟▌ ▟█▙  █▟█▌
██ ██ ██ ▐▌ ▐▌▐▛ ▐▌ ▘▄▟▌▐▛ ▜▌▐▙▄▟▌ █▘  
██ ▀▀ ██ ▐▌ ▐▌▐▌ ▐▌▗█▀▜▌▐▌ ▐▌▐▛▀▀▘ █   
██    ██ ▐▙▄█▌▐▌ ▐▌▐▙▄█▌▝█▄█▌▝█▄▄▌ █   
▀▀    ▀▀  ▀▀▝▘▝▘ ▝▘ ▀▀▝▘ ▞▀▐▌ ▝▀▀  ▀   
                         ▜█▛▘                  
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
