package client

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/TurboHsu/munager/cmd/sync/utils"
	"github.com/TurboHsu/munager/util/logging"
	"github.com/spf13/cobra"
)

var ClientCommand = &cobra.Command{
	Use:   "client",
	Short: "Run a sync client",
	Long: `Run a sync client to sync music files between devices.
By default it will look for server broadcast message and connect to it.
You can also specify server address manually.`,
	Example: `munager sync client --transcode -p "/home/me/Music" --thread=8 --trim`,
}

func init() {
	ClientCommand.Flags().StringP("address", "a", ":10721", "Address to connect to")
	ClientCommand.Flags().StringP("path", "p", ".", "Path to sync")
	ClientCommand.Flags().IntP("thread", "t", 5, "Number of threads to use")
	ClientCommand.Flags().BoolP("transcode", "T", false, "Transcode the file received using ffmpeg")
	ClientCommand.Flags().StringP("ffmpeg-path", "F", "ffmpeg", "Where ffmpeg is located")
	ClientCommand.Flags().StringP("ffmpeg-arg", "A", "-vn -ar 48000 -b:a 512k -n", "Arguments to ffmpeg")
	ClientCommand.Flags().StringP("output-format", "O", "opus", "Output format")
	ClientCommand.Flags().BoolP("silent", "s", true, "Stop spamming transcoding log. It is on by default, to disable it, run --silent=false")
	ClientCommand.Flags().BoolP("trim", "m", false, "Delete audio related files that are not in the server")
	ClientCommand.Run = runClient
}

func runClient(cmd *cobra.Command, args []string) {
	addr, err := ClientCommand.Flags().GetString("address")
	logging.HandleErr(err)
	_, err = ClientCommand.Flags().GetString("path")
	logging.HandleErr(err)

	utils.FixPath(ClientCommand)
	if len(strings.Split(addr, ":")[0]) == 0 {
		logging.Info("Finding sync server...")

		// Wait for broadcast
		dest, port := WaitBroadcast(addr)

		logging.Info("Found server at " + dest + ":" + fmt.Sprint(port) + "...")
		connectServer(dest, port)
	} else {
		logging.Info("Connecting to " + addr + "...")
		port, err := strconv.Atoi(strings.Split(addr, ":")[1])
		logging.HandleErr(err)
		connectServer(strings.Split(addr, ":")[0], port)
	}
}
