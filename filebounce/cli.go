package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

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

func main() {
	auth := flag.String("auth", "XXX", "authentication key to send to transfer node")
	host := flag.String("host", "localhost:8888", "host of transfer node")
	flag.Parse()

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
		conn.Conn.Close()
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
