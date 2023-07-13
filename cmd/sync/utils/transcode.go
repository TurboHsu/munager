package utils

import (
	"os/exec"
	"strings"
)

func Transcode(basePath string, ffmpegPath string, ffmpegArg string, originalFormat string, outputFormat string) error {
	args := []string{
		"-i", basePath + "." + originalFormat,
	}
	args = append(args, strings.Split(ffmpegArg, " ")...)
	args = append(args, basePath+"."+outputFormat)
	cmd := exec.Command(ffmpegPath, args...)

	// // Get a pipe to the command's stderr output
	// stderr, err := cmd.StderrPipe()
	// if err != nil {
	// 	return err
	// }

	// Start the command
	err := cmd.Start()
	if err != nil {
		return err
	}

	// // Read from the stderr pipe
	// scanner := bufio.NewScanner(stderr)
	// for scanner.Scan() {
	// 	// Handle the stderr output here
	// 	fmt.Println(scanner.Text())
	// }

	// Wait for the command to complete
	err = cmd.Wait()
	if err != nil {
		return err
	}

	// Wait for file release

	return nil
}
