// Code generated by protoc-gen-go.
// source: clientmessage.proto
// DO NOT EDIT!

/*
Package protobufs is a generated protocol buffer package.

It is generated from these files:
	clientmessage.proto

It has these top-level messages:
	ClientToTransferNodeMessage
	AuthenticateData
	StartUploadData
	UploadData
	FinishedData
	TransferNodeToClientMessage
	AuthSuccessData
	TransferCreatedData
	RecipientsData
	ProgressData
	ErrorData
*/
package protobufs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ClientToTransferNodeMessage_MessageType int32

const (
	ClientToTransferNodeMessage_AUTHENTICATE ClientToTransferNodeMessage_MessageType = 0
	ClientToTransferNodeMessage_START_UPLOAD ClientToTransferNodeMessage_MessageType = 1
	ClientToTransferNodeMessage_UPLOAD_DATA  ClientToTransferNodeMessage_MessageType = 2
	ClientToTransferNodeMessage_FINISHED     ClientToTransferNodeMessage_MessageType = 3
	ClientToTransferNodeMessage_ERROR        ClientToTransferNodeMessage_MessageType = 4
)

var ClientToTransferNodeMessage_MessageType_name = map[int32]string{
	0: "AUTHENTICATE",
	1: "START_UPLOAD",
	2: "UPLOAD_DATA",
	3: "FINISHED",
	4: "ERROR",
}
var ClientToTransferNodeMessage_MessageType_value = map[string]int32{
	"AUTHENTICATE": 0,
	"START_UPLOAD": 1,
	"UPLOAD_DATA":  2,
	"FINISHED":     3,
	"ERROR":        4,
}

func (x ClientToTransferNodeMessage_MessageType) String() string {
	return proto.EnumName(ClientToTransferNodeMessage_MessageType_name, int32(x))
}
func (ClientToTransferNodeMessage_MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

type TransferNodeToClientMessage_MessageType int32

const (
	TransferNodeToClientMessage_AUTH_SUCCESS     TransferNodeToClientMessage_MessageType = 0
	TransferNodeToClientMessage_TRANSFER_CREATED TransferNodeToClientMessage_MessageType = 1
	TransferNodeToClientMessage_RECIPIENTS       TransferNodeToClientMessage_MessageType = 3
	TransferNodeToClientMessage_PROGRESS         TransferNodeToClientMessage_MessageType = 4
	TransferNodeToClientMessage_FINISHED         TransferNodeToClientMessage_MessageType = 5
	TransferNodeToClientMessage_ERROR            TransferNodeToClientMessage_MessageType = 6
)

var TransferNodeToClientMessage_MessageType_name = map[int32]string{
	0: "AUTH_SUCCESS",
	1: "TRANSFER_CREATED",
	3: "RECIPIENTS",
	4: "PROGRESS",
	5: "FINISHED",
	6: "ERROR",
}
var TransferNodeToClientMessage_MessageType_value = map[string]int32{
	"AUTH_SUCCESS":     0,
	"TRANSFER_CREATED": 1,
	"RECIPIENTS":       3,
	"PROGRESS":         4,
	"FINISHED":         5,
	"ERROR":            6,
}

func (x TransferNodeToClientMessage_MessageType) String() string {
	return proto.EnumName(TransferNodeToClientMessage_MessageType_name, int32(x))
}
func (TransferNodeToClientMessage_MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5, 0}
}

type ClientToTransferNodeMessage struct {
	Type         ClientToTransferNodeMessage_MessageType `protobuf:"varint,1,opt,name=type,enum=protobufs.ClientToTransferNodeMessage_MessageType" json:"type,omitempty"`
	AuthData     *AuthenticateData                       `protobuf:"bytes,2,opt,name=authData" json:"authData,omitempty"`
	StartData    *StartUploadData                        `protobuf:"bytes,3,opt,name=startData" json:"startData,omitempty"`
	UploadData   *UploadData                             `protobuf:"bytes,4,opt,name=uploadData" json:"uploadData,omitempty"`
	FinishedData *FinishedData                           `protobuf:"bytes,5,opt,name=finishedData" json:"finishedData,omitempty"`
	Timestamp    int64                                   `protobuf:"varint,6,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *ClientToTransferNodeMessage) Reset()                    { *m = ClientToTransferNodeMessage{} }
func (m *ClientToTransferNodeMessage) String() string            { return proto.CompactTextString(m) }
func (*ClientToTransferNodeMessage) ProtoMessage()               {}
func (*ClientToTransferNodeMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ClientToTransferNodeMessage) GetType() ClientToTransferNodeMessage_MessageType {
	if m != nil {
		return m.Type
	}
	return ClientToTransferNodeMessage_AUTHENTICATE
}

func (m *ClientToTransferNodeMessage) GetAuthData() *AuthenticateData {
	if m != nil {
		return m.AuthData
	}
	return nil
}

func (m *ClientToTransferNodeMessage) GetStartData() *StartUploadData {
	if m != nil {
		return m.StartData
	}
	return nil
}

func (m *ClientToTransferNodeMessage) GetUploadData() *UploadData {
	if m != nil {
		return m.UploadData
	}
	return nil
}

func (m *ClientToTransferNodeMessage) GetFinishedData() *FinishedData {
	if m != nil {
		return m.FinishedData
	}
	return nil
}

func (m *ClientToTransferNodeMessage) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type AuthenticateData struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *AuthenticateData) Reset()                    { *m = AuthenticateData{} }
func (m *AuthenticateData) String() string            { return proto.CompactTextString(m) }
func (*AuthenticateData) ProtoMessage()               {}
func (*AuthenticateData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AuthenticateData) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type StartUploadData struct {
	Filename string `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
	Mimetype string `protobuf:"bytes,2,opt,name=mimetype" json:"mimetype,omitempty"`
	Size     uint64 `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
}

func (m *StartUploadData) Reset()                    { *m = StartUploadData{} }
func (m *StartUploadData) String() string            { return proto.CompactTextString(m) }
func (*StartUploadData) ProtoMessage()               {}
func (*StartUploadData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *StartUploadData) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *StartUploadData) GetMimetype() string {
	if m != nil {
		return m.Mimetype
	}
	return ""
}

func (m *StartUploadData) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type UploadData struct {
	Order uint64 `protobuf:"varint,1,opt,name=order" json:"order,omitempty"`
	Size  uint64 `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	Data  []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *UploadData) Reset()                    { *m = UploadData{} }
func (m *UploadData) String() string            { return proto.CompactTextString(m) }
func (*UploadData) ProtoMessage()               {}
func (*UploadData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UploadData) GetOrder() uint64 {
	if m != nil {
		return m.Order
	}
	return 0
}

func (m *UploadData) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *UploadData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type FinishedData struct {
	Error   string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	Success bool   `protobuf:"varint,2,opt,name=success" json:"success,omitempty"`
}

func (m *FinishedData) Reset()                    { *m = FinishedData{} }
func (m *FinishedData) String() string            { return proto.CompactTextString(m) }
func (*FinishedData) ProtoMessage()               {}
func (*FinishedData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *FinishedData) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *FinishedData) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type TransferNodeToClientMessage struct {
	Type                TransferNodeToClientMessage_MessageType `protobuf:"varint,1,opt,name=type,enum=protobufs.TransferNodeToClientMessage_MessageType" json:"type,omitempty"`
	AuthSuccessData     *AuthSuccessData                        `protobuf:"bytes,2,opt,name=authSuccessData" json:"authSuccessData,omitempty"`
	TransferCreatedData *TransferCreatedData                    `protobuf:"bytes,3,opt,name=transferCreatedData" json:"transferCreatedData,omitempty"`
	RecipientsData      *RecipientsData                         `protobuf:"bytes,4,opt,name=recipientsData" json:"recipientsData,omitempty"`
	ProgressData        *ProgressData                           `protobuf:"bytes,5,opt,name=progressData" json:"progressData,omitempty"`
	FinishedData        *FinishedData                           `protobuf:"bytes,6,opt,name=finishedData" json:"finishedData,omitempty"`
	ErrorData           *ErrorData                              `protobuf:"bytes,7,opt,name=errorData" json:"errorData,omitempty"`
	Timestamp           int64                                   `protobuf:"varint,8,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *TransferNodeToClientMessage) Reset()                    { *m = TransferNodeToClientMessage{} }
func (m *TransferNodeToClientMessage) String() string            { return proto.CompactTextString(m) }
func (*TransferNodeToClientMessage) ProtoMessage()               {}
func (*TransferNodeToClientMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *TransferNodeToClientMessage) GetType() TransferNodeToClientMessage_MessageType {
	if m != nil {
		return m.Type
	}
	return TransferNodeToClientMessage_AUTH_SUCCESS
}

func (m *TransferNodeToClientMessage) GetAuthSuccessData() *AuthSuccessData {
	if m != nil {
		return m.AuthSuccessData
	}
	return nil
}

func (m *TransferNodeToClientMessage) GetTransferCreatedData() *TransferCreatedData {
	if m != nil {
		return m.TransferCreatedData
	}
	return nil
}

func (m *TransferNodeToClientMessage) GetRecipientsData() *RecipientsData {
	if m != nil {
		return m.RecipientsData
	}
	return nil
}

func (m *TransferNodeToClientMessage) GetProgressData() *ProgressData {
	if m != nil {
		return m.ProgressData
	}
	return nil
}

func (m *TransferNodeToClientMessage) GetFinishedData() *FinishedData {
	if m != nil {
		return m.FinishedData
	}
	return nil
}

func (m *TransferNodeToClientMessage) GetErrorData() *ErrorData {
	if m != nil {
		return m.ErrorData
	}
	return nil
}

func (m *TransferNodeToClientMessage) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type AuthSuccessData struct {
}

func (m *AuthSuccessData) Reset()                    { *m = AuthSuccessData{} }
func (m *AuthSuccessData) String() string            { return proto.CompactTextString(m) }
func (*AuthSuccessData) ProtoMessage()               {}
func (*AuthSuccessData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type TransferCreatedData struct {
	TransferId string `protobuf:"bytes,1,opt,name=transferId" json:"transferId,omitempty"`
}

func (m *TransferCreatedData) Reset()                    { *m = TransferCreatedData{} }
func (m *TransferCreatedData) String() string            { return proto.CompactTextString(m) }
func (*TransferCreatedData) ProtoMessage()               {}
func (*TransferCreatedData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *TransferCreatedData) GetTransferId() string {
	if m != nil {
		return m.TransferId
	}
	return ""
}

type RecipientsData struct {
	Recipients []*RecipientsData_Recipient `protobuf:"bytes,1,rep,name=recipients" json:"recipients,omitempty"`
}

func (m *RecipientsData) Reset()                    { *m = RecipientsData{} }
func (m *RecipientsData) String() string            { return proto.CompactTextString(m) }
func (*RecipientsData) ProtoMessage()               {}
func (*RecipientsData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RecipientsData) GetRecipients() []*RecipientsData_Recipient {
	if m != nil {
		return m.Recipients
	}
	return nil
}

type RecipientsData_Recipient struct {
	Ipv4     string `protobuf:"bytes,1,opt,name=ipv4" json:"ipv4,omitempty"`
	Ipv6     string `protobuf:"bytes,2,opt,name=ipv6" json:"ipv6,omitempty"`
	Identity string `protobuf:"bytes,3,opt,name=identity" json:"identity,omitempty"`
}

func (m *RecipientsData_Recipient) Reset()                    { *m = RecipientsData_Recipient{} }
func (m *RecipientsData_Recipient) String() string            { return proto.CompactTextString(m) }
func (*RecipientsData_Recipient) ProtoMessage()               {}
func (*RecipientsData_Recipient) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8, 0} }

func (m *RecipientsData_Recipient) GetIpv4() string {
	if m != nil {
		return m.Ipv4
	}
	return ""
}

func (m *RecipientsData_Recipient) GetIpv6() string {
	if m != nil {
		return m.Ipv6
	}
	return ""
}

func (m *RecipientsData_Recipient) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

type ProgressData struct {
	BytesUploaded int64 `protobuf:"varint,1,opt,name=bytesUploaded" json:"bytesUploaded,omitempty"`
}

func (m *ProgressData) Reset()                    { *m = ProgressData{} }
func (m *ProgressData) String() string            { return proto.CompactTextString(m) }
func (*ProgressData) ProtoMessage()               {}
func (*ProgressData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ProgressData) GetBytesUploaded() int64 {
	if m != nil {
		return m.BytesUploaded
	}
	return 0
}

type ErrorData struct {
	Title       string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	JsonDetails string `protobuf:"bytes,2,opt,name=jsonDetails" json:"jsonDetails,omitempty"`
	Fatal       bool   `protobuf:"varint,3,opt,name=fatal" json:"fatal,omitempty"`
}

func (m *ErrorData) Reset()                    { *m = ErrorData{} }
func (m *ErrorData) String() string            { return proto.CompactTextString(m) }
func (*ErrorData) ProtoMessage()               {}
func (*ErrorData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ErrorData) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ErrorData) GetJsonDetails() string {
	if m != nil {
		return m.JsonDetails
	}
	return ""
}

func (m *ErrorData) GetFatal() bool {
	if m != nil {
		return m.Fatal
	}
	return false
}

func init() {
	proto.RegisterType((*ClientToTransferNodeMessage)(nil), "protobufs.ClientToTransferNodeMessage")
	proto.RegisterType((*AuthenticateData)(nil), "protobufs.AuthenticateData")
	proto.RegisterType((*StartUploadData)(nil), "protobufs.StartUploadData")
	proto.RegisterType((*UploadData)(nil), "protobufs.UploadData")
	proto.RegisterType((*FinishedData)(nil), "protobufs.FinishedData")
	proto.RegisterType((*TransferNodeToClientMessage)(nil), "protobufs.TransferNodeToClientMessage")
	proto.RegisterType((*AuthSuccessData)(nil), "protobufs.AuthSuccessData")
	proto.RegisterType((*TransferCreatedData)(nil), "protobufs.TransferCreatedData")
	proto.RegisterType((*RecipientsData)(nil), "protobufs.RecipientsData")
	proto.RegisterType((*RecipientsData_Recipient)(nil), "protobufs.RecipientsData.Recipient")
	proto.RegisterType((*ProgressData)(nil), "protobufs.ProgressData")
	proto.RegisterType((*ErrorData)(nil), "protobufs.ErrorData")
	proto.RegisterEnum("protobufs.ClientToTransferNodeMessage_MessageType", ClientToTransferNodeMessage_MessageType_name, ClientToTransferNodeMessage_MessageType_value)
	proto.RegisterEnum("protobufs.TransferNodeToClientMessage_MessageType", TransferNodeToClientMessage_MessageType_name, TransferNodeToClientMessage_MessageType_value)
}

func init() { proto.RegisterFile("clientmessage.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 752 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x94, 0xdd, 0x6e, 0xda, 0x4a,
	0x10, 0xc7, 0x0f, 0x60, 0x08, 0x1e, 0x38, 0xe0, 0xb3, 0xc9, 0xd1, 0xe1, 0x90, 0x2a, 0x8a, 0xdc,
	0x5c, 0xf4, 0x8a, 0x0b, 0x9a, 0xb4, 0x95, 0x2a, 0x55, 0xb2, 0x8c, 0x69, 0xa8, 0x5a, 0x40, 0x63,
	0x73, 0x51, 0x55, 0x15, 0x72, 0x60, 0x49, 0xdc, 0x02, 0x46, 0xb6, 0xa9, 0x44, 0x1f, 0xa8, 0xaf,
	0xd2, 0x77, 0xea, 0x55, 0x77, 0xd7, 0x36, 0x2c, 0x0e, 0x4d, 0xd5, 0x2b, 0x76, 0x3e, 0x7e, 0xe3,
	0x65, 0x66, 0xfe, 0x0b, 0xc7, 0x93, 0xb9, 0x47, 0x97, 0xd1, 0x82, 0x86, 0xa1, 0x7b, 0x4b, 0x5b,
	0xab, 0xc0, 0x8f, 0x7c, 0xa2, 0x8a, 0x9f, 0x9b, 0xf5, 0x2c, 0xd4, 0xbf, 0x17, 0xe0, 0xd4, 0x14,
	0x29, 0x8e, 0xef, 0x04, 0xee, 0x32, 0x9c, 0xd1, 0xa0, 0xef, 0x4f, 0xe9, 0xbb, 0x18, 0x20, 0x5d,
	0x50, 0xa2, 0xcd, 0x8a, 0x36, 0x72, 0xe7, 0xb9, 0x27, 0xb5, 0x76, 0xbb, 0xb5, 0x25, 0x5b, 0x0f,
	0x50, 0xad, 0xe4, 0xd7, 0x61, 0x24, 0x0a, 0x9e, 0x3c, 0x87, 0xb2, 0xbb, 0x8e, 0xee, 0x3a, 0x6e,
	0xe4, 0x36, 0xf2, 0xac, 0x56, 0xa5, 0x7d, 0x2a, 0xd5, 0x32, 0x58, 0x88, 0x15, 0xf3, 0x26, 0x6e,
	0x44, 0x79, 0x0a, 0x6e, 0x93, 0xc9, 0x0b, 0x50, 0xc3, 0xc8, 0x0d, 0x22, 0x41, 0x16, 0x04, 0xd9,
	0x94, 0x48, 0x9b, 0xc7, 0x46, 0xab, 0xb9, 0xef, 0x4e, 0x05, 0xb8, 0x4b, 0x26, 0x57, 0x00, 0xeb,
	0x6d, 0xa0, 0xa1, 0x08, 0xf4, 0x5f, 0x09, 0x95, 0x28, 0x29, 0x91, 0xbc, 0x84, 0xea, 0xcc, 0x5b,
	0x7a, 0xe1, 0x1d, 0x8d, 0xc1, 0xa2, 0x00, 0xff, 0x93, 0xc0, 0xae, 0x14, 0xc6, 0xbd, 0x64, 0xf2,
	0x08, 0xd4, 0xc8, 0x63, 0xcd, 0x8e, 0xdc, 0xc5, 0xaa, 0x51, 0x62, 0x64, 0x01, 0x77, 0x0e, 0xfd,
	0x03, 0x54, 0xa4, 0xce, 0x10, 0x0d, 0xaa, 0xc6, 0xc8, 0xb9, 0xb6, 0xfa, 0x4e, 0xcf, 0x34, 0x1c,
	0x4b, 0xfb, 0x8b, 0x7b, 0x6c, 0xc7, 0x40, 0x67, 0x3c, 0x1a, 0xbe, 0x1d, 0x18, 0x1d, 0x2d, 0x47,
	0xea, 0x50, 0x89, 0xcf, 0xe3, 0x8e, 0xe1, 0x18, 0x5a, 0x9e, 0x54, 0xa1, 0xdc, 0xed, 0xf5, 0x7b,
	0xf6, 0xb5, 0xd5, 0xd1, 0x0a, 0x44, 0x85, 0xa2, 0x85, 0x38, 0x40, 0x4d, 0xd1, 0x2f, 0x40, 0xcb,
	0xb6, 0x91, 0xd5, 0x2b, 0x7c, 0xa6, 0x1b, 0x31, 0x3c, 0x15, 0xf9, 0x51, 0xff, 0x08, 0xf5, 0x4c,
	0xcb, 0x48, 0x13, 0xca, 0x33, 0x6f, 0x4e, 0x97, 0xee, 0x82, 0x26, 0x99, 0x5b, 0x9b, 0xc7, 0x16,
	0xec, 0xfa, 0x62, 0x05, 0xf2, 0x71, 0x2c, 0xb5, 0x09, 0x01, 0x25, 0xf4, 0xbe, 0x52, 0x31, 0x14,
	0x05, 0xc5, 0x59, 0x7f, 0x03, 0x20, 0x55, 0x3e, 0x81, 0xa2, 0x1f, 0x4c, 0x69, 0x20, 0xca, 0x2a,
	0x18, 0x1b, 0x5b, 0x2e, 0xbf, 0xe3, 0xb8, 0x6f, 0x9a, 0x0e, 0xb8, 0x8a, 0xe2, 0xac, 0xbf, 0x82,
	0xaa, 0xdc, 0x69, 0x5e, 0x8d, 0x06, 0x81, 0x1f, 0x24, 0x97, 0x8c, 0x0d, 0xd2, 0x80, 0xa3, 0x70,
	0x3d, 0x99, 0xb0, 0xb6, 0x8a, 0x82, 0x65, 0x4c, 0x4d, 0xfd, 0x87, 0x02, 0xa7, 0xf2, 0x72, 0x3a,
	0x7e, 0xbc, 0xb2, 0xbf, 0x5f, 0xed, 0x07, 0xa8, 0x03, 0xab, 0xdd, 0x81, 0x3a, 0xdf, 0x56, 0x3b,
	0xfe, 0xac, 0xb4, 0xe1, 0xcd, 0xcc, 0x86, 0x4b, 0x19, 0x98, 0x45, 0xc8, 0x10, 0x8e, 0xa3, 0xe4,
	0xb3, 0x66, 0x40, 0xd9, 0x00, 0xa7, 0xd2, 0xc6, 0x9f, 0x1d, 0xb8, 0x9c, 0x94, 0x85, 0x87, 0x50,
	0x62, 0x40, 0x2d, 0xa0, 0x13, 0x6f, 0xc5, 0x6f, 0x1f, 0x4a, 0x1a, 0xf8, 0x5f, 0x2a, 0x86, 0x7b,
	0x09, 0x98, 0x01, 0xb8, 0x16, 0x58, 0xee, 0x6d, 0x90, 0xfe, 0xaf, 0xfb, 0x5a, 0x18, 0x4a, 0x61,
	0xdc, 0x4b, 0xbe, 0x27, 0xa4, 0xd2, 0x9f, 0x08, 0xa9, 0x0d, 0xaa, 0x98, 0xaf, 0x20, 0x8f, 0x04,
	0x79, 0x22, 0x91, 0x56, 0x1a, 0xc3, 0x5d, 0xda, 0xbe, 0xf8, 0xca, 0x59, 0xf1, 0xcd, 0x0f, 0x8a,
	0x6f, 0x6c, 0x8f, 0x4c, 0xd3, 0xb2, 0x6d, 0x26, 0xbe, 0x13, 0xd0, 0x1c, 0x34, 0xfa, 0x76, 0xd7,
	0xc2, 0xb1, 0x89, 0x16, 0x53, 0x24, 0x17, 0x60, 0x0d, 0x00, 0x2d, 0xb3, 0x37, 0xec, 0x31, 0x99,
	0xda, 0x4c, 0x71, 0x4c, 0x7f, 0x43, 0x1c, 0xbc, 0x46, 0xce, 0x28, 0x7b, 0x6a, 0x2c, 0xee, 0xd4,
	0x58, 0xd2, 0xff, 0x81, 0x7a, 0x66, 0xe4, 0xfa, 0x15, 0x1c, 0x1f, 0x98, 0x1d, 0x39, 0x03, 0x48,
	0xa7, 0xd7, 0x9b, 0x26, 0xbb, 0x2d, 0x79, 0xf4, 0x6f, 0x39, 0xa8, 0xed, 0x8f, 0x89, 0x98, 0x00,
	0xbb, 0x41, 0x31, 0xa4, 0xc0, 0xba, 0xf3, 0xf8, 0x97, 0x53, 0xdd, 0x99, 0x28, 0x61, 0xcd, 0x01,
	0xa8, 0xdb, 0x00, 0xd7, 0x9f, 0xb7, 0xfa, 0x72, 0x99, 0x7c, 0x5e, 0x9c, 0x13, 0xdf, 0xb3, 0x44,
	0xf7, 0xe2, 0xcc, 0xdf, 0x03, 0x6f, 0xca, 0x9f, 0x98, 0x68, 0x23, 0x56, 0x93, 0xbd, 0x07, 0xa9,
	0xad, 0x5f, 0x42, 0x55, 0xde, 0x06, 0x72, 0x01, 0x7f, 0xdf, 0x6c, 0x22, 0x1a, 0xc6, 0x0f, 0x02,
	0x8d, 0xff, 0x5b, 0x01, 0xf7, 0x9d, 0xfa, 0x7b, 0x50, 0xb7, 0xc3, 0xe4, 0x12, 0x67, 0xa5, 0xe6,
	0xe9, 0x3b, 0x14, 0x1b, 0xe4, 0x1c, 0x2a, 0x9f, 0x42, 0x7f, 0xd9, 0xa1, 0x91, 0xeb, 0xcd, 0xc3,
	0xe4, 0x3e, 0xb2, 0x8b, 0x73, 0x33, 0xc6, 0xcf, 0xc5, 0x9d, 0xca, 0x18, 0x1b, 0x37, 0x25, 0xd1,
	0x91, 0xa7, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x55, 0x07, 0x5e, 0x04, 0x07, 0x00, 0x00,
}
