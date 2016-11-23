package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	humanize "github.com/dustin/go-humanize"
	"github.com/fsufitch/filebounce-cli/api"
)

func getPath() (string, error) {
	path := flag.Arg(0)
	if len(path) == 0 {
		return "", errors.New("No file specified")
	}
	fileInfo, err := os.Stat(path)
	if err != nil {
		return "", err
	}
	if fileInfo.IsDir() {
		return "", errors.New("Cannot upload a directory")
	}
	return path, nil
}

func waitToUpload(delay int, cb func()) {
	if delay < 1 {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Hit Enter to start upload. ")
		_, _ = reader.ReadString('\n')
	} else {
		time.Sleep(time.Duration(delay) * time.Second)
	}
	cb()
}

func main() {
	auth := flag.String("auth", "XXX", "authentication key to send to transfer node")
	host := flag.String("host", "localhost:8888", "host of transfer node")
	delay := flag.Int("delay", 0, "delay before upload is triggered (0 = wait for user input)")
	help := flag.Bool("h", false, "help")
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		return
	}

	path, err := getPath()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	url := fmt.Sprintf("ws://%s/client_ws", *host)
	fmt.Printf("Connecting to: %s\n", url)
	fmt.Printf("Using auth: %s\n", *auth)
	fmt.Printf("Uploading file: %s\n", path)

	conn, doneChan := api.ConnectTransferNode(url, *auth)

	if conn.Error != nil {
		fmt.Printf("Error connecting! %s\n", conn.Error)
		return
	}

	conn.Authenticate()

	if conn.Error != nil {
		fmt.Fprintf(os.Stderr, "Error authenticating! %s\n", conn.Error)
	} else {
		fmt.Println("Connection successful!")
	}

	conn.ReceivedFileIDCallback = func(fileID string) {
		fmt.Printf("Download file at: http://%s/download/%s\n", *host, fileID)
		waitToUpload(*delay, func() {
			fmt.Printf("Uploading...")
			err = conn.UploadFile(path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error streaming file: %s\n", err)
			}
		})
	}

	stat, _ := os.Stat(path)
	lastProgressPercent := uint64(0)
	conn.ProgressCallback = func(bytes uint64) {
		progressPercent := float64(bytes) / float64(stat.Size()) * 100
		progressPercentInt := uint64(progressPercent)
		if progressPercentInt-lastProgressPercent >= 10 {
			lastProgressPercent = progressPercentInt
			fmt.Printf(" ... Progress %s / %s (%.2f%%)\n", humanize.Bytes(bytes), humanize.Bytes(uint64(stat.Size())), progressPercent)
		}
	}

	conn.SelectFile(path)

	if conn.Error != nil {
		fmt.Fprintf(os.Stderr, "Error sending file metadata! %s\n", conn.Error)
	}

	wsErr := <-doneChan
	if wsErr != nil {
		fmt.Fprintf(os.Stderr, "Websocket closed with error %s\n", wsErr)
	}
}
