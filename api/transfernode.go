package api

import (
	"errors"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/fsufitch/filebounce-cli/protobufs"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

// TransferNodeConnection holds details about a connection to a transfer node
type TransferNodeConnection struct {
	Conn                   *websocket.Conn
	AuthKey                string
	FileID                 string
	ReceivedFileIDCallback func(string)
	Error                  error
}

func ConnectTransferNode(url string, auth string) (*TransferNodeConnection, chan error) {
	conn, _, err := websocket.DefaultDialer.Dial(url, http.Header{})
	newConn := &TransferNodeConnection{Conn: conn, Error: err, AuthKey: auth}
	doneChan := make(chan error)
	go newConn.readLoop(doneChan)
	return newConn, doneChan
}

func (conn *TransferNodeConnection) readLoop(doneChan chan error) {
	defer conn.Conn.Close()
	for {
		buf, closed, err := conn.readMessage()
		if err != nil {
			doneChan <- err
			break
		}

		if closed {
			doneChan <- nil
			break
		}

		err = conn.triageMessage(buf)
		if err != nil {
			doneChan <- err
			break
		}
	}
}

func (conn *TransferNodeConnection) readMessage() (buf []byte, closed bool, err error) {
	closed = false
	msgType, buf, err := conn.Conn.ReadMessage()
	if err != nil {
		return
	}
	if msgType == websocket.CloseMessage {
		closed = true
	} else if msgType != websocket.BinaryMessage {
		err = errors.New("Unexpected non-binary message type")
	}
	return
}

func (conn *TransferNodeConnection) triageMessage(buf []byte) (err error) {
	message := &protobufs.TransferNodeToClientMessage{}
	err = proto.Unmarshal(buf, message)
	if err != nil {
		return
	}

	switch message.GetType() {
	case protobufs.TransferNodeToClientMessage_ERROR:
		err = errors.New(message.GetErrorData().String())
		break
	case protobufs.TransferNodeToClientMessage_AUTH_SUCCESS:
		// Nothing to do here
		break
	case protobufs.TransferNodeToClientMessage_TRANSFER_CREATED:
		conn.FileID = message.TransferCreatedData.GetTransferId()
		if conn.ReceivedFileIDCallback != nil {
			conn.ReceivedFileIDCallback(conn.FileID)
		}
		break
	default:
		err = errors.New("Unexpected message type")
		break
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
}

func (conn *TransferNodeConnection) SelectFile(path string) {
	fi, _ := os.Stat(path)

	fname := fi.Name()
	fsize := fi.Size()

	mimetype := mime.TypeByExtension(filepath.Ext(path))
	if mimetype == "" {
		mimetype = "application/octet-stream"
	}

	message := &protobufs.ClientToTransferNodeMessage{
		Type: protobufs.ClientToTransferNodeMessage_START_UPLOAD,
		StartData: &protobufs.StartUploadData{
			Filename: fname,
			Size:     uint64(fsize),
			Mimetype: mimetype,
		},
		Timestamp: time.Now().Unix(),
	}

	data, _ := proto.Marshal(message)
	conn.Error = conn.Conn.WriteMessage(websocket.BinaryMessage, data)
}
