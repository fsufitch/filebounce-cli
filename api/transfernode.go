package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/fsufitch/filebounce-cli/protobufs"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

// TransferNodeConnection holds details about a connection to a transfer node
type TransferNodeConnection struct {
	Conn    *websocket.Conn
	AuthKey string
	FileID  string
	Error   error
}

func ConnectTransferNode(url string) *TransferNodeConnection {
	conn, _, err := websocket.DefaultDialer.Dial(url, http.Header{})
	return &TransferNodeConnection{Conn: conn, Error: err}
}

func (conn *TransferNodeConnection) readMessage() (buf []byte, err error) {
	msgType, buf, err := conn.Conn.ReadMessage()
	if err != nil {
		return
	}
	if msgType != websocket.BinaryMessage {
		err = errors.New("Unexpected non-binary message type")
	}
	return
}

func (conn *TransferNodeConnection) Authenticate() {
	message := &protobufs.ClientToTransferNodeMessage{
		Type: protobufs.ClientToTransferNodeMessage_AUTHENTICATE,
		AuthData: &protobufs.AuthenticateData{
			Key: conn.AuthKey,
		},
		Timestamp: time.Now().Unix(),
	}

	data, _ := proto.Marshal(message)
	conn.Error = conn.Conn.WriteMessage(websocket.BinaryMessage, data)
	if conn.Error != nil {
		return
	}

	buf, err := conn.readMessage()
	if err != nil {
		conn.Error = err
		return
	}

	response := &protobufs.TransferNodeToClientMessage{}
	err = proto.Unmarshal(buf, response)

	if err != nil {
		conn.Error = err
		return
	}

	switch response.GetType() {
	case protobufs.TransferNodeToClientMessage_ERROR:
		conn.Error = errors.New(response.GetErrorData().String())
		break
	case protobufs.TransferNodeToClientMessage_AUTH_SUCCESS:
		break
	default:
		conn.Error = errors.New("Unexpected message type")
		break
	}
}
