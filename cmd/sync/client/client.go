package client

import (
	"fmt"

	"github.com/TurboHsu/munager/util/logging"
	"github.com/spf13/cobra"
)

var ClientCommand = &cobra.Command{
	Use:   "client",
	Short: "Run a sync client",
	Long: `Run a sync client to sync music files between devices.
By default it will look for server broadcast message and connect to it.
You can also specify server address manually.`,
}

func init() {
	ClientCommand.Flags().StringP("address", "a", ":10721", "Address to connect to")
	ClientCommand.Flags().StringP("path", "p", ".", "Path to sync")
	ClientCommand.Run = runClient
}

func runClient(cmd *cobra.Command, args []string) {
	addr, err := ClientCommand.Flags().GetString("address")
	logging.HandleErr(err)
	_, err = ClientCommand.Flags().GetString("path")
	logging.HandleErr(err)

	logging.Info("Finding sync server...")

	// Wait for broadcast
	dest, port := WaitBroadcast(addr)

	fmt.Println(dest, port)
}
