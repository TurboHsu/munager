package client

import (
	"net"
	"strconv"

	"github.com/TurboHsu/munager/cmd/sync/server"
	"github.com/TurboHsu/munager/util/logging"
)

func WaitBroadcast(addr string) (dest string) {
	// Get the listing port
	_, port, err := net.SplitHostPort(addr)
	logging.HandleErr(err)
	portInt, err := strconv.Atoi(port)
	logging.HandleErr(err)
	for {
		// Listen for broadcast message
		conn, err := net.ListenUDP("udp", &net.UDPAddr{
			Port: portInt,
		})
		logging.HandleErr(err)
		defer conn.Close()

		// Wait for UDP packet
		buf := make([]byte, 64)
		n, src, err := conn.ReadFromUDP(buf)
		logging.HandleErr(err)

		// Convert addr to string
		dest = src.IP.String() + ":" + strconv.Itoa(src.Port)

		// Check if the message is correct
		if string(buf[:n]) != server.BroadcastMessage {
			logging.Info("Received wrong broadcast message from " + dest + " , continue waiting...")
		} else {
			logging.Info("Received broadcast message from " + dest)
			return
		}
	}
}
