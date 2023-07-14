package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/TurboHsu/munager/cmd/sync/structure"
	"github.com/TurboHsu/munager/cmd/sync/utils"
	fileprocessing "github.com/TurboHsu/munager/util/file"
	"github.com/TurboHsu/munager/util/logging"
	"github.com/schollz/progressbar/v3"
)

var Fingerprint string

func connectServer(dest string, port int) {
	// Handshakes with server
	serverAddr := dest + ":" + strconv.Itoa(port)
	err := performHandshake(serverAddr)
	if err != nil {
		logging.HandleErr(err)
		return
	}

	// Trim files if needed
	needTrim, err := ClientCommand.Flags().GetBool("trim")
	logging.HandleErr(err)
	if needTrim {
		err = trimFiles(serverAddr)
		if err != nil {
			logging.HandleErr(err)
			return
		}
	}

	// Sends local file list
	files, err := sendFileList(serverAddr)
	if err != nil {
		logging.HandleErr(err)
		return
	}

	// Gets the files
	getFiles(serverAddr, files)

	// Suicide
	err = suicide(serverAddr)
	if err != nil {
		logging.HandleErr(err)
		return
	}
	logging.Info("Done!")
}

func trimFiles(addr string) error {
	// Read silent flag
	isSilent, err := ClientCommand.Flags().GetBool("silent")
	logging.HandleErr(err)

	// Construct API addr
	addr = "http://" + addr + "/api/get-file"

	// Sends nothing to get the full list
	sendFileListMsg, err := json.Marshal(structure.ListRequest{
		Fingerprint: Fingerprint,
	})
	if err != nil {
		return err
	}

	resp, err := http.Post(addr, "application/json", bytes.NewBuffer(sendFileListMsg))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result structure.ListResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return err
	}

	// Get local file list
	path := ClientCommand.Flag("path").Value.String()
	files := utils.GetFiles(path)
	files = utils.FilterValidFiles(files)

	// Get the trim list
	for _, f := range files {
		burn := true
		for _, r := range result.Files {
			if f.PathBase == r.PathBase {
				burn = false
				break
			}
		}

		if burn {
			if !isSilent {
				logging.Info("Deleting " + f.PathBase + "." + f.Extension + "...")
			}
			filePath := ClientCommand.Flag("path").Value.String() + f.PathBase + "." + f.Extension
			err = os.Remove(filePath)
			if err != nil {
				logging.HandleErr(err)
				return err
			}
		}
	}

	return nil
}

func suicide(addr string) error {
	suicideMessage, err := json.Marshal(structure.Suicide{
		Fingerprint: Fingerprint,
	})
	logging.HandleErr(err)

	addr = "http://" + addr + "/api/suicide"

	if err != nil {
		return err
	}

	resp, err := http.Post(addr, "application/json", bytes.NewBuffer(suicideMessage))
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
		return fmt.Errorf("suicide failed, code: %d, msg: %s", result.Code, result.Msg)
	}

	return nil
}

func getFiles(addr string, files []structure.FileInfo) {
	// Read thread limits
	threadLimit, err := ClientCommand.Flags().GetInt("thread")
	logging.HandleErr(err)

	// Read transcode flag
	transcode, err := ClientCommand.Flags().GetBool("transcode")
	logging.HandleErr(err)

	// Read FFmpeg path
	ffmpegPath, err := ClientCommand.Flags().GetString("ffmpeg-path")
	logging.HandleErr(err)

	// Read FFmpeg arguments
	ffmpegArg, err := ClientCommand.Flags().GetString("ffmpeg-arg")
	logging.HandleErr(err)

	// Read output format
	outputFormat, err := ClientCommand.Flags().GetString("output-format")
	logging.HandleErr(err)

	// Read silent flag
	isSilent, err := ClientCommand.Flags().GetBool("silent")
	logging.HandleErr(err)

	// Construct API address
	addr = "http://" + addr + "/api/get-file"

	var wg sync.WaitGroup
	jobs := make(chan bool, threadLimit)
	bar := progressbar.Default(int64(len(files)))

	for _, file := range files {
		jobs <- true
		wg.Add(1)
		go func(f structure.FileInfo) {
			// Construct request body
			requestBody, err := json.Marshal(structure.FileServeRequest{
				PathBase:    f.PathBase,
				Extension:   f.Extension,
				Fingerprint: Fingerprint,
			})
			if err != nil {
				logging.HandleErr(err)
				bar.Add(1)
				wg.Done()
				<-jobs
				return
			}

			client := &http.Client{}
			req, err := http.NewRequest("POST", addr, bytes.NewBuffer(requestBody))
			if err != nil {
				logging.HandleErr(err)
				bar.Add(1)
				wg.Done()
				<-jobs
				return
			}
			resp, err := client.Do(req)
			if err != nil {
				logging.HandleErr(err)
				bar.Add(1)
				wg.Done()
				<-jobs
				return
			}
			defer resp.Body.Close()

			filePath := ClientCommand.Flag("path").Value.String() + f.PathBase + "." + f.Extension
			// Checks if file exists
			if _, err := os.Stat(filePath); err == nil {
				logging.Info("File " + filePath + " already exists, skipping...")
				bar.Add(1)
				wg.Done()
				<-jobs
				return
			}

			// Checks whether parent directory exists
			parentPath := filepath.Dir(filePath)
			if _, err := os.Stat(parentPath); os.IsNotExist(err) {
				err = os.MkdirAll(parentPath, 0755)
				if err != nil {
					logging.HandleErr(err)
					bar.Add(1)
					wg.Done()
					<-jobs
					return
				}
			}

			// Creates the file
			downloadedFile, err := os.Create(filePath)
			if err != nil {
				logging.HandleErr(err)
				bar.Add(1)
				wg.Done()
				<-jobs
				return
			}

			_, err = io.Copy(downloadedFile, resp.Body)
			if err != nil {
				logging.HandleErr(err)
				bar.Add(1)
				wg.Done()
				<-jobs
				return
			}

			// Closes the file
			err = downloadedFile.Close()
			if err != nil {
				logging.HandleErr(err)
				bar.Add(1)
				wg.Done()
				<-jobs
				return
			}

			// Checks whether the user wants to run ffmpeg
			if transcode {
				if fileprocessing.IsExtAudio("." + f.Extension) {
					if !isSilent {
						logging.Info("Transcoding " + filePath + "...")
					}
					// Run transcoding
					err := utils.Transcode(ClientCommand.Flag("path").Value.String()+f.PathBase, ffmpegPath, ffmpegArg, f.Extension, outputFormat)
					if err != nil {
						logging.HandleErr(err)
					}

					// Deletes the original file
					err = os.Remove(filePath)
					if err != nil {
						logging.HandleErr(err)
					}
				}
			}

			bar.Add(1)
			wg.Done()
			<-jobs
		}(file)
	}

	wg.Wait()
}

func sendFileList(addr string) ([]structure.FileInfo, error) {
	// Generates local file list
	path := ClientCommand.Flag("path").Value.String()
	files := utils.GetFiles(path)
	files = utils.FilterValidFiles(files)
	addr = "http://" + addr + "/api/get-list"

	sendFileListMsg, err := json.Marshal(structure.ListRequest{
		Files:       files,
		Fingerprint: Fingerprint,
	})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(addr, "application/json", bytes.NewBuffer(sendFileListMsg))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result structure.ListResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Files, nil
}

// This function performs handshake with server
func performHandshake(addr string) error {
	addr = "http://" + addr + "/api/handshake"

	Fingerprint = utils.GenerateRandomString(32)
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
