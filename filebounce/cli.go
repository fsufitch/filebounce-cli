package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/cheggaaa/pb"
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

func waitToUpload(conn *api.TransferNodeConnection, delay int, waitRecipients int) {
	switch {
	case delay > 0:
		time.Sleep(time.Duration(delay) * time.Second)
		fallthrough
	case waitRecipients > 0:
		gotRecipients := 0
		fmt.Printf("Waiting for recipients (%d/%d)...\n", gotRecipients, waitRecipients)
		for gotRecipients < waitRecipients {
			newRecipients := <-conn.ReceivedRecipients
			gotRecipients = len(newRecipients)
			fmt.Printf("Got recipients (%d/%d): ", gotRecipients, waitRecipients)
			for _, recipient := range newRecipients {
				switch {
				case recipient.Identity != "":
					fmt.Printf("%s ", recipient.Identity)
				case recipient.Ipv4 != "":
					fmt.Printf("%s ", recipient.Ipv4)
				case recipient.Ipv6 != "":
					fmt.Printf("%s ", recipient.Ipv6)
				default:
					fmt.Print("(unknown) ")
				}
			}
			fmt.Println()
		}
	default:
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Hit Enter to start upload. ")
		_, _ = reader.ReadString('\n')
	}
}

func main() {
	auth := flag.String("auth", "XXX", "authentication key to send to transfer node")
	host := flag.String("host", "localhost:8888", "host of transfer node")
	delay := flag.Int("delay", 0, "delay before upload is triggered (0 = wait for user input)")
	waitRecipients := flag.Int("recipients", 0, "number of recipients to wait for (0 = do not wait)")
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

	conn := api.ConnectTransferNode(url, *auth)

	go func(errChan chan error) {
		err := <-errChan
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}(conn.ReceivedError)

	conn.Authenticate()
	fmt.Printf("Connected and authenticated with transfer node: %s\n", *host)

	conn.SelectFile(path)

	fileID := <-conn.ReceivedFileID
	fmt.Printf("Download file at: http://%s/download/%s\n", *host, fileID)

	waitToUpload(conn, *delay, *waitRecipients)
	fmt.Printf("Uploading...")

	go conn.UploadChunksOnRequest(path)

	stat, _ := os.Stat(path)
	progressBar := pb.New(int(stat.Size()))
	progressBar.SetRefreshRate(500 * time.Millisecond)
	progressBar.ShowPercent = true
	progressBar.ShowBar = true
	progressBar.ShowSpeed = true
	progressBar.SetUnits(pb.U_BYTES)
	progressBar.Start()

	done := false
	for !done {
		select {
		case progressBytes := <-conn.ReceivedProgressBytes:
			progressBar.Set64(int64(progressBytes))
		case <-conn.Done:
			done = true
			progressBar.FinishPrint("Upload complete!")
		case <-conn.ReceivedError:
			progressBar.FinishPrint("Upload failed!")
			//fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
			//os.Exit(1)
		}
	}
}
