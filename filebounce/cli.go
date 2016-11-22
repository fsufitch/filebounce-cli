package main

import (
	"fmt"

	"github.com/fsufitch/filebounce-cli/api"
)

func main() {
	conn := api.ConnectTransferNode("ws://localhost:8888/client_ws")
	conn.AuthKey = "fake auth key"

	if conn.Error != nil {
		fmt.Printf("Error connecting! %s\n", conn.Error)
		return
	}

	conn.Authenticate()

	if conn.Error != nil {
		fmt.Printf("Error authenticating! %s\n", conn.Error)
	} else {
		fmt.Println("Connection successful!")
	}
}
