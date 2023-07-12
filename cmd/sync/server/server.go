package server

import (
	"time"

	"github.com/TurboHsu/munager/util/logging"
	"github.com/spf13/cobra"
)

var ServerCommand = &cobra.Command{
	Use:   "server",
	Short: "Run a sync server",
	Long: `Run a sync server to sync music files between devices.
Server side will broadcast its IP address and port to the local network.
Client side will connect to the server and sync music files.`,
}

func init() {
	ServerCommand.Flags().StringP("address", "a", ":10721", "Address to listen on")
	ServerCommand.Flags().BoolP("broadcast", "b", true, "Broadcast server address to local network")
	ServerCommand.Flags().StringP("path", "p", ".", "Path to sync")
	ServerCommand.Flags().BoolP("keep-broadcasting", "k", false, "Keep broadcasting server address to local network even client handshaked with server")
	ServerCommand.Run = runServer
}

func runServer(cmd *cobra.Command, args []string) {
	doBroadcast, err := ServerCommand.Flags().GetBool("broadcast")
	logging.HandleErr(err)
	addr, err := ServerCommand.Flags().GetString("address")
	logging.HandleErr(err)

	if doBroadcast {
		go broadcast(addr)
	}

	listen(addr)

	for {
		time.Sleep(1 * time.Second)
	}
}
