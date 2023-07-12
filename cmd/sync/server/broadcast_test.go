package server

import (
	"net"
	"testing"
)

func TestBroadcasting(t *testing.T) {
	go broadcast(":10721")

	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: 10721,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	buf := make([]byte, 64)
	n, _, err := conn.ReadFromUDP(buf)
	if err != nil {
		t.Fatal(err)
	}

	if string(buf[:n]) != BroadcastMessage {
		t.Fatalf("Expected %s, got %s", BroadcastMessage, string(buf[:n]))
	}
}
