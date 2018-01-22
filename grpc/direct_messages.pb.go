// Code generated by protoc-gen-go. DO NOT EDIT.
// source: direct_messages.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	direct_messages.proto

It has these top-level messages:
	CreateMessageRequest
	MessageCreate
	Target
	MessageData
	CreateMessageResponse
	Event
*/
package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc1 "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateMessageRequest struct {
	MessageCreate *MessageCreate `protobuf:"bytes,1,opt,name=messageCreate" json:"messageCreate,omitempty"`
}

func (m *CreateMessageRequest) Reset()                    { *m = CreateMessageRequest{} }
func (m *CreateMessageRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateMessageRequest) ProtoMessage()               {}
func (*CreateMessageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateMessageRequest) GetMessageCreate() *MessageCreate {
	if m != nil {
		return m.MessageCreate
	}
	return nil
}

type MessageCreate struct {
	Target      *Target      `protobuf:"bytes,1,opt,name=target" json:"target,omitempty"`
	MessageData *MessageData `protobuf:"bytes,2,opt,name=messageData" json:"messageData,omitempty"`
}

func (m *MessageCreate) Reset()                    { *m = MessageCreate{} }
func (m *MessageCreate) String() string            { return proto.CompactTextString(m) }
func (*MessageCreate) ProtoMessage()               {}
func (*MessageCreate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MessageCreate) GetTarget() *Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *MessageCreate) GetMessageData() *MessageData {
	if m != nil {
		return m.MessageData
	}
	return nil
}

type Target struct {
	RecipientId string `protobuf:"bytes,1,opt,name=recipient_id,json=recipientId" json:"recipient_id,omitempty"`
}

func (m *Target) Reset()                    { *m = Target{} }
func (m *Target) String() string            { return proto.CompactTextString(m) }
func (*Target) ProtoMessage()               {}
func (*Target) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Target) GetRecipientId() string {
	if m != nil {
		return m.RecipientId
	}
	return ""
}

type MessageData struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *MessageData) Reset()                    { *m = MessageData{} }
func (m *MessageData) String() string            { return proto.CompactTextString(m) }
func (*MessageData) ProtoMessage()               {}
func (*MessageData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MessageData) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type CreateMessageResponse struct {
	Event *Event `protobuf:"bytes,1,opt,name=event" json:"event,omitempty"`
}

func (m *CreateMessageResponse) Reset()                    { *m = CreateMessageResponse{} }
func (m *CreateMessageResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateMessageResponse) ProtoMessage()               {}
func (*CreateMessageResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CreateMessageResponse) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type Event struct {
	Id               string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	CreatedTimestamp string `protobuf:"bytes,2,opt,name=created_timestamp,json=createdTimestamp" json:"created_timestamp,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetCreatedTimestamp() string {
	if m != nil {
		return m.CreatedTimestamp
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateMessageRequest)(nil), "grpc.CreateMessageRequest")
	proto.RegisterType((*MessageCreate)(nil), "grpc.MessageCreate")
	proto.RegisterType((*Target)(nil), "grpc.Target")
	proto.RegisterType((*MessageData)(nil), "grpc.MessageData")
	proto.RegisterType((*CreateMessageResponse)(nil), "grpc.CreateMessageResponse")
	proto.RegisterType((*Event)(nil), "grpc.Event")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for DirectMessagesService service

type DirectMessagesServiceClient interface {
	CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc1.CallOption) (*CreateMessageResponse, error)
}

type directMessagesServiceClient struct {
	cc *grpc1.ClientConn
}

func NewDirectMessagesServiceClient(cc *grpc1.ClientConn) DirectMessagesServiceClient {
	return &directMessagesServiceClient{cc}
}

func (c *directMessagesServiceClient) CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc1.CallOption) (*CreateMessageResponse, error) {
	out := new(CreateMessageResponse)
	err := grpc1.Invoke(ctx, "/grpc.DirectMessagesService/CreateMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DirectMessagesService service

type DirectMessagesServiceServer interface {
	CreateMessage(context.Context, *CreateMessageRequest) (*CreateMessageResponse, error)
}

func RegisterDirectMessagesServiceServer(s *grpc1.Server, srv DirectMessagesServiceServer) {
	s.RegisterService(&_DirectMessagesService_serviceDesc, srv)
}

func _DirectMessagesService_CreateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectMessagesServiceServer).CreateMessage(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.DirectMessagesService/CreateMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectMessagesServiceServer).CreateMessage(ctx, req.(*CreateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DirectMessagesService_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "grpc.DirectMessagesService",
	HandlerType: (*DirectMessagesServiceServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "CreateMessage",
			Handler:    _DirectMessagesService_CreateMessage_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "direct_messages.proto",
}

func init() { proto.RegisterFile("direct_messages.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcf, 0x6b, 0xfa, 0x40,
	0x10, 0xc5, 0x51, 0x54, 0x70, 0xa2, 0x5f, 0xbe, 0x4e, 0x2b, 0x88, 0xbd, 0xd4, 0xa5, 0x87, 0x82,
	0xe0, 0x41, 0x4f, 0xed, 0xb5, 0x29, 0xb4, 0x07, 0x0f, 0xdd, 0x7a, 0x0f, 0xdb, 0x64, 0x08, 0x5b,
	0xc8, 0x8f, 0xee, 0x4e, 0xa5, 0x7f, 0x7e, 0xc9, 0xee, 0x2a, 0x49, 0xf1, 0x96, 0xbc, 0xcf, 0xcb,
	0xdb, 0x37, 0x99, 0x85, 0x79, 0xa6, 0x0d, 0xa5, 0x9c, 0x14, 0x64, 0xad, 0xca, 0xc9, 0x6e, 0x6a,
	0x53, 0x71, 0x85, 0x83, 0xdc, 0xd4, 0xa9, 0x78, 0x83, 0xeb, 0x27, 0x43, 0x8a, 0x69, 0xef, 0xa9,
	0xa4, 0xaf, 0x6f, 0xb2, 0x8c, 0x0f, 0x30, 0x0d, 0x7e, 0x8f, 0x17, 0xbd, 0xdb, 0xde, 0x7d, 0xb4,
	0xbd, 0xda, 0x34, 0x5f, 0x6d, 0xf6, 0x6d, 0x24, 0xbb, 0x4e, 0xf1, 0x09, 0xd3, 0x0e, 0xc7, 0x3b,
	0x18, 0xb1, 0x32, 0x39, 0x71, 0x08, 0x99, 0xf8, 0x90, 0x83, 0xd3, 0x64, 0x60, 0xb8, 0x83, 0x28,
	0xe4, 0xc4, 0x8a, 0xd5, 0xa2, 0xef, 0xac, 0xb3, 0xce, 0x79, 0x0d, 0x90, 0x6d, 0x97, 0x58, 0xc3,
	0xc8, 0xc7, 0xe0, 0x0a, 0x26, 0x86, 0x52, 0x5d, 0x6b, 0x2a, 0x39, 0xd1, 0x99, 0x3b, 0x6a, 0x2c,
	0xa3, 0xb3, 0xf6, 0x9a, 0x89, 0x15, 0x44, 0xad, 0x20, 0x44, 0x18, 0x30, 0xfd, 0x70, 0x70, 0xba,
	0x67, 0xf1, 0x08, 0xf3, 0x3f, 0xbf, 0xc3, 0xd6, 0x55, 0x69, 0x09, 0x57, 0x30, 0xa4, 0x23, 0x95,
	0xa7, 0x11, 0x22, 0xdf, 0xeb, 0xb9, 0x91, 0xa4, 0x27, 0x22, 0x86, 0xa1, 0x7b, 0xc7, 0x7f, 0xd0,
	0x3f, 0x17, 0xe8, 0xeb, 0x0c, 0xd7, 0x30, 0x4b, 0x5d, 0x68, 0x96, 0xb0, 0x2e, 0xc8, 0xb2, 0x2a,
	0x6a, 0x37, 0xdf, 0x58, 0xfe, 0x0f, 0xe0, 0x70, 0xd2, 0xb7, 0x0a, 0xe6, 0xb1, 0xdb, 0x57, 0x68,
	0x60, 0xdf, 0xc9, 0x1c, 0x75, 0x4a, 0xf8, 0x02, 0xd3, 0x4e, 0x35, 0x5c, 0xfa, 0x0e, 0x97, 0xd6,
	0xb7, 0xbc, 0xb9, 0xc8, 0xfc, 0x2c, 0x1f, 0x23, 0x77, 0x01, 0x76, 0xbf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xed, 0x3b, 0x12, 0xf8, 0x19, 0x02, 0x00, 0x00,
}
