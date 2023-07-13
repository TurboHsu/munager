package client

import (
	"strings"
	"testing"
)

func TestFingerprintGeneration(t *testing.T) {
	str := generateFingerprint(65536)
	if strings.Contains(str, "\"") || strings.Contains(str, "\\") {
		t.Fatalf("Fingerprint contains \" or \\, which is bad for JSON")
	}
}
