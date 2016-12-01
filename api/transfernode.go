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
	Conn                  *websocket.Conn
	Closed                bool
	AuthKey               string
	FileID                string
	ChunksRequests        chan ChunksRequest
	ReceivedFileID        chan string
	ReceivedProgressBytes chan uint64
	ReceivedError         chan error
	Done                  chan bool
	Error                 error
}

// ChunksRequest represents a request from the transfer node for more data
type ChunksRequest struct {
	Count uint64
	Bytes uint64
}

// ConnectTransferNode creates a new connection to a transfer node
func ConnectTransferNode(url string, auth string) *TransferNodeConnection {
	conn, _, err := websocket.DefaultDialer.Dial(url, http.Header{})
	newConn := &TransferNodeConnection{
		Conn:                  conn,
		Error:                 err,
		AuthKey:               auth,
		ChunksRequests:        make(chan ChunksRequest, 10), // buffered so we can write to it before listening to it
		ReceivedFileID:        make(chan string),
		ReceivedProgressBytes: make(chan uint64),
		ReceivedError:         make(chan error),
		Done:                  make(chan bool),
	}
	go newConn.readLoop()
	return newConn
}

func (conn *TransferNodeConnection) readLoop() {
	defer conn.Conn.Close()
	conn.Conn.SetCloseHandler(func(code int, text string) error {
		conn.Closed = true
		if code != websocket.CloseNormalClosure && code != websocket.CloseNoStatusReceived {
			err := fmt.Errorf("Abnormal closure %d %s", code, text)
			conn.ReceivedError <- err
			return err
		}
		return nil
	})

	for !conn.Closed {
		buf, err := conn.readMessage()
		if conn.Closed {
			continue
		}
		if err != nil {
			conn.ReceivedError <- err
			return
		}

		err = conn.triageMessage(buf)
		if err != nil {
			conn.ReceivedError <- err
			return
		}
	}
	conn.Done <- true
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
		conn.ReceivedFileID <- conn.FileID
		if message.TransferCreatedData.GetRequestChunks() > 0 {
			conn.ChunksRequests <- ChunksRequest{
				Count: message.TransferCreatedData.GetRequestChunks(),
				Bytes: message.TransferCreatedData.GetChunkSize(),
			}
		}
		break
	case protobufs.TransferNodeToClientMessage_PROGRESS:
		bytesProgress := message.ProgressData.GetBytesUploaded()
		conn.ReceivedProgressBytes <- uint64(bytesProgress)
		if message.ProgressData.GetRequestChunks() > 0 {
			conn.ChunksRequests <- ChunksRequest{
				Count: message.ProgressData.GetRequestChunks(),
				Bytes: message.ProgressData.GetChunkSize(),
			}
		}
		break
	default:
		err = errors.New("Unexpected message type")
		break
	}
	return
}

// Authenticate runs an authentication request on the connection
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

// SelectFile informs the transfer node of the file/metadata selected
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
			Order: uint64(order),
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

// UploadChunksOnRequest is an asynchronous function for  uploading chunks
// when requests for them arrive through a given channel
func (conn *TransferNodeConnection) UploadChunksOnRequest(path string) {
	file, _ := os.Open(path)
	chunkCount := 0
	for request := range conn.ChunksRequests {
		for i := 0; uint64(i) < request.Count; i++ {
			buf := make([]byte, request.Bytes) // read however many bytes requested
			byteCount, err := file.Read(buf)
			if err != nil && err != io.EOF {
				fmt.Fprintf(os.Stderr, "got err %s", err)
				conn.ReceivedError <- err
				return
			}
			if err == io.EOF {
				conn.sendComplete()
			} else {
				conn.sendChunk(chunkCount, uint64(byteCount), buf[:byteCount])
			}

			if conn.Error != nil {
				conn.ReceivedError <- conn.Error
				return
			}
			chunkCount++
		}
	}
	conn.sendComplete()
}
