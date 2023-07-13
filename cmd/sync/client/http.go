package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/TurboHsu/munager/cmd/sync/structure"
	"github.com/TurboHsu/munager/util/logging"
)

var Fingerprint string

func connectServer(dest string, port int) {
	serverAddr := dest + ":" + strconv.Itoa(port)
	err := performHandshake(serverAddr)
	if err != nil {
		logging.HandleErr(err)
		return
	}
}

func generateFingerprint(len int) (ret string) {
	// Summon random string
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := 0; i < len; i++ {
		char := r.Intn(87)
		// Had better not to use backslash
		if char == 52 {
			char = -2
		}
		ret += fmt.Sprintf("%c", rune(char+40))
	}
	return
}

func performHandshake(addr string) error {
	addr = "http://" + addr + "/api/handshake"

	Fingerprint = generateFingerprint(32)
	handshakeMsg, err := json.Marshal(structure.Handshake{
		Fingerprint: Fingerprint,
		MagicWord:   structure.HandshakeMagicWord,
	})
	if err != nil {
		return err
	}

	resp, err := http.Post(addr, "application/json", bytes.NewBuffer(handshakeMsg))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result structure.BasicResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return err
	}
	if result.Code != 200 {
		return fmt.Errorf("handshake failed, code: %d, msg: %s", result.Code, result.Msg)
	}

	logging.Info("Handshake success, client fingerprint: " + Fingerprint)
	return nil
}
