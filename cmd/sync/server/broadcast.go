package server

import (
	"net"
	"strconv"
	"time"

	"github.com/TurboHsu/munager/cmd/sync/structure"
	"github.com/TurboHsu/munager/util/logging"
)

var TerminateBroadcast bool

func Broadcast(addr string) {
	TerminateBroadcast = false

	// Parse listening port
	_, port, err := net.SplitHostPort(addr)
	logging.HandleErr(err)
	portInt, err := strconv.Atoi(port)
	logging.HandleErr(err)

	logging.Info("Start broadcasting on port " + port + "...")

	for !TerminateBroadcast {
		conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
			IP:   net.IPv4(255, 255, 255, 255),
			Port: portInt,
		})
		logging.HandleErr(err)
		defer conn.Close()

		// Send broadcast message
		_, err = conn.Write([]byte(structure.BroadcastMessage + port))
		logging.HandleErr(err)

		// Sleep
		time.Sleep(5 * time.Second)
	}
}
