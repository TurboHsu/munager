package sync

import (
	"testing"

	"github.com/TurboHsu/munager/cmd/sync/client"
	"github.com/TurboHsu/munager/cmd/sync/server"
)

func TestBroadcasting(t *testing.T) {
	go server.Broadcast(":10721")
	dest, port := client.WaitBroadcast(":10721")
	server.TerminateBroadcast = true
	if port != 10721 {
		t.Fatalf("Expected 10721, got %d", port)
	}
	t.Log(dest, port)
}
