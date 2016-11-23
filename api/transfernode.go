package api

import (
	"errors"
	"fmt"
	"io"
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
	Closed                 bool
	AuthKey                string
	FileID                 string
	ReceivedFileIDCallback func(string)
	ProgressCallback       func(uint64)
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
	conn.Conn.SetCloseHandler(func(code int, text string) error {
		conn.Closed = true
		if code != websocket.CloseNormalClosure && code != websocket.CloseNoStatusReceived {
			return fmt.Errorf("Abnormal closure %d %s", code, text)
		}
		return nil
	})
	for !conn.Closed {
		buf, err := conn.readMessage()
		if conn.Closed {
			continue
		}
		if err != nil {
			doneChan <- err
			return
		}

		err = conn.triageMessage(buf)
		if err != nil {
			doneChan <- err
			return
		}
	}
	doneChan <- nil
}

func (conn *TransferNodeConnection) readMessage() (buf []byte, err error) {
	msgType, buf, err := conn.Conn.ReadMessage()

	if err != nil {
		return
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
	case protobufs.TransferNodeToClientMessage_PROGRESS:
		bytesProgress := message.ProgressData.GetBytesUploaded()
		if conn.ProgressCallback != nil {
			conn.ProgressCallback(uint64(bytesProgress))
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

func (conn *TransferNodeConnection) sendChunk(order int, size uint64, chunk []byte) {
	message := &protobufs.ClientToTransferNodeMessage{
		Type: protobufs.ClientToTransferNodeMessage_UPLOAD_DATA,
		UploadData: &protobufs.UploadData{
			Order: 0,
			Size:  size,
			Data:  chunk[:],
		},
		Timestamp: time.Now().Unix(),
	}
	data, _ := proto.Marshal(message)
	conn.Error = conn.Conn.WriteMessage(websocket.BinaryMessage, data)
}

func (conn *TransferNodeConnection) sendComplete() {
	message := &protobufs.ClientToTransferNodeMessage{
		Type:         protobufs.ClientToTransferNodeMessage_FINISHED,
		FinishedData: &protobufs.FinishedData{},
		Timestamp:    time.Now().Unix(),
	}
	data, _ := proto.Marshal(message)
	conn.Error = conn.Conn.WriteMessage(websocket.BinaryMessage, data)
}

func (conn *TransferNodeConnection) UploadFile(path string) error {
	file, _ := os.Open(path)
	buf := make([]byte, 50000) // read 50 KB chunks
	var count int
	var err error
	for i := 0; err == nil; i++ {
		count, err = file.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Printf("got err %s", err)
			return err
		}

		conn.sendChunk(i, uint64(count), buf[:count])
		if conn.Error != nil {
			return conn.Error
		}

	}
	conn.sendComplete()
	return conn.Error
}
