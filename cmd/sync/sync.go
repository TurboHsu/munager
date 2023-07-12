package sync

import (
	"github.com/TurboHsu/munager/cmd/sync/client"
	"github.com/TurboHsu/munager/cmd/sync/server"
	"github.com/spf13/cobra"
)

var SyncRoot = &cobra.Command{
	Use:   "sync",
	Short: "Sync music files between devices",
	Long: `Munager sync is a tool to sync music files between devices.
You should run both server side and client side to get it working.
See 'munager sync server --help' and 'munager sync client --help' for more information.`,
}

func init() {
	SyncRoot.AddCommand(server.ServerCommand)
	SyncRoot.AddCommand(client.ClientCommand)
}
