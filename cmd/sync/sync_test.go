package sync

import (
	"strings"
	"testing"

	"github.com/TurboHsu/munager/cmd/sync/client"
	"github.com/TurboHsu/munager/cmd/sync/server"
	"github.com/TurboHsu/munager/cmd/sync/structure"
	"github.com/TurboHsu/munager/cmd/sync/utils"
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

func TestRandomGeneration(t *testing.T) {
	str := utils.GenerateRandomString(65536)
	if strings.Contains(str, "\"") || strings.Contains(str, "\\") {
		t.Fatalf("Fingerprint contains \" or \\, which is bad for JSON")
	}
}

func TestFileFilter(t *testing.T) {
	f := []structure.FileInfo{
		{
			PathBase:  "valid",
			Extension: "lrc",
		},
		{
			PathBase:  "valid",
			Extension: "mp3",
		},
		{
			PathBase:  "valid",
			Extension: "jpg",
		},
		{
			PathBase:  "invalid",
			Extension: "txt",
		},
	}

	f = utils.FilterValidFiles(f)
	for _, a := range f {
		if a.PathBase == "invalid" {
			t.Fatalf("Invalid file is not filtered")
		}
	}
}
