// Code generated by protoc-gen-go. DO NOT EDIT.
// source: spire/server/notifier/notifier.proto

package notifier

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common "github.com/spiffe/spire/proto/spire/common"
	plugin "github.com/spiffe/spire/proto/spire/common/plugin"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type BundleLoaded struct {
	Bundle               *common.Bundle `protobuf:"bytes,1,opt,name=bundle,proto3" json:"bundle,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BundleLoaded) Reset()         { *m = BundleLoaded{} }
func (m *BundleLoaded) String() string { return proto.CompactTextString(m) }
func (*BundleLoaded) ProtoMessage()    {}
func (*BundleLoaded) Descriptor() ([]byte, []int) {
	return fileDescriptor_c27428e9e6d193e9, []int{0}
}

func (m *BundleLoaded) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BundleLoaded.Unmarshal(m, b)
}
func (m *BundleLoaded) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BundleLoaded.Marshal(b, m, deterministic)
}
func (m *BundleLoaded) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BundleLoaded.Merge(m, src)
}
func (m *BundleLoaded) XXX_Size() int {
	return xxx_messageInfo_BundleLoaded.Size(m)
}
func (m *BundleLoaded) XXX_DiscardUnknown() {
	xxx_messageInfo_BundleLoaded.DiscardUnknown(m)
}

var xxx_messageInfo_BundleLoaded proto.InternalMessageInfo

func (m *BundleLoaded) GetBundle() *common.Bundle {
	if m != nil {
		return m.Bundle
	}
	return nil
}

type BundleUpdated struct {
	Bundle               *common.Bundle `protobuf:"bytes,1,opt,name=bundle,proto3" json:"bundle,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BundleUpdated) Reset()         { *m = BundleUpdated{} }
func (m *BundleUpdated) String() string { return proto.CompactTextString(m) }
func (*BundleUpdated) ProtoMessage()    {}
func (*BundleUpdated) Descriptor() ([]byte, []int) {
	return fileDescriptor_c27428e9e6d193e9, []int{1}
}

func (m *BundleUpdated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BundleUpdated.Unmarshal(m, b)
}
func (m *BundleUpdated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BundleUpdated.Marshal(b, m, deterministic)
}
func (m *BundleUpdated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BundleUpdated.Merge(m, src)
}
func (m *BundleUpdated) XXX_Size() int {
	return xxx_messageInfo_BundleUpdated.Size(m)
}
func (m *BundleUpdated) XXX_DiscardUnknown() {
	xxx_messageInfo_BundleUpdated.DiscardUnknown(m)
}

var xxx_messageInfo_BundleUpdated proto.InternalMessageInfo

func (m *BundleUpdated) GetBundle() *common.Bundle {
	if m != nil {
		return m.Bundle
	}
	return nil
}

type NotifyRequest struct {
	// Types that are valid to be assigned to Event:
	//	*NotifyRequest_BundleUpdated
	Event                isNotifyRequest_Event `protobuf_oneof:"event"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *NotifyRequest) Reset()         { *m = NotifyRequest{} }
func (m *NotifyRequest) String() string { return proto.CompactTextString(m) }
func (*NotifyRequest) ProtoMessage()    {}
func (*NotifyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c27428e9e6d193e9, []int{2}
}

func (m *NotifyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyRequest.Unmarshal(m, b)
}
func (m *NotifyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyRequest.Marshal(b, m, deterministic)
}
func (m *NotifyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyRequest.Merge(m, src)
}
func (m *NotifyRequest) XXX_Size() int {
	return xxx_messageInfo_NotifyRequest.Size(m)
}
func (m *NotifyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyRequest proto.InternalMessageInfo

type isNotifyRequest_Event interface {
	isNotifyRequest_Event()
}

type NotifyRequest_BundleUpdated struct {
	BundleUpdated *BundleUpdated `protobuf:"bytes,1,opt,name=bundle_updated,json=bundleUpdated,proto3,oneof"`
}

func (*NotifyRequest_BundleUpdated) isNotifyRequest_Event() {}

func (m *NotifyRequest) GetEvent() isNotifyRequest_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *NotifyRequest) GetBundleUpdated() *BundleUpdated {
	if x, ok := m.GetEvent().(*NotifyRequest_BundleUpdated); ok {
		return x.BundleUpdated
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*NotifyRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*NotifyRequest_BundleUpdated)(nil),
	}
}

type NotifyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyResponse) Reset()         { *m = NotifyResponse{} }
func (m *NotifyResponse) String() string { return proto.CompactTextString(m) }
func (*NotifyResponse) ProtoMessage()    {}
func (*NotifyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c27428e9e6d193e9, []int{3}
}

func (m *NotifyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyResponse.Unmarshal(m, b)
}
func (m *NotifyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyResponse.Marshal(b, m, deterministic)
}
func (m *NotifyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyResponse.Merge(m, src)
}
func (m *NotifyResponse) XXX_Size() int {
	return xxx_messageInfo_NotifyResponse.Size(m)
}
func (m *NotifyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyResponse proto.InternalMessageInfo

type NotifyAndAdviseRequest struct {
	// Types that are valid to be assigned to Event:
	//	*NotifyAndAdviseRequest_BundleLoaded
	Event                isNotifyAndAdviseRequest_Event `protobuf_oneof:"event"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *NotifyAndAdviseRequest) Reset()         { *m = NotifyAndAdviseRequest{} }
func (m *NotifyAndAdviseRequest) String() string { return proto.CompactTextString(m) }
func (*NotifyAndAdviseRequest) ProtoMessage()    {}
func (*NotifyAndAdviseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c27428e9e6d193e9, []int{4}
}

func (m *NotifyAndAdviseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyAndAdviseRequest.Unmarshal(m, b)
}
func (m *NotifyAndAdviseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyAndAdviseRequest.Marshal(b, m, deterministic)
}
func (m *NotifyAndAdviseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyAndAdviseRequest.Merge(m, src)
}
func (m *NotifyAndAdviseRequest) XXX_Size() int {
	return xxx_messageInfo_NotifyAndAdviseRequest.Size(m)
}
func (m *NotifyAndAdviseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyAndAdviseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyAndAdviseRequest proto.InternalMessageInfo

type isNotifyAndAdviseRequest_Event interface {
	isNotifyAndAdviseRequest_Event()
}

type NotifyAndAdviseRequest_BundleLoaded struct {
	BundleLoaded *BundleLoaded `protobuf:"bytes,1,opt,name=bundle_loaded,json=bundleLoaded,proto3,oneof"`
}

func (*NotifyAndAdviseRequest_BundleLoaded) isNotifyAndAdviseRequest_Event() {}

func (m *NotifyAndAdviseRequest) GetEvent() isNotifyAndAdviseRequest_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *NotifyAndAdviseRequest) GetBundleLoaded() *BundleLoaded {
	if x, ok := m.GetEvent().(*NotifyAndAdviseRequest_BundleLoaded); ok {
		return x.BundleLoaded
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*NotifyAndAdviseRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*NotifyAndAdviseRequest_BundleLoaded)(nil),
	}
}

type NotifyAndAdviseResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyAndAdviseResponse) Reset()         { *m = NotifyAndAdviseResponse{} }
func (m *NotifyAndAdviseResponse) String() string { return proto.CompactTextString(m) }
func (*NotifyAndAdviseResponse) ProtoMessage()    {}
func (*NotifyAndAdviseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c27428e9e6d193e9, []int{5}
}

func (m *NotifyAndAdviseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyAndAdviseResponse.Unmarshal(m, b)
}
func (m *NotifyAndAdviseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyAndAdviseResponse.Marshal(b, m, deterministic)
}
func (m *NotifyAndAdviseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyAndAdviseResponse.Merge(m, src)
}
func (m *NotifyAndAdviseResponse) XXX_Size() int {
	return xxx_messageInfo_NotifyAndAdviseResponse.Size(m)
}
func (m *NotifyAndAdviseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyAndAdviseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyAndAdviseResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BundleLoaded)(nil), "spire.server.notifier.BundleLoaded")
	proto.RegisterType((*BundleUpdated)(nil), "spire.server.notifier.BundleUpdated")
	proto.RegisterType((*NotifyRequest)(nil), "spire.server.notifier.NotifyRequest")
	proto.RegisterType((*NotifyResponse)(nil), "spire.server.notifier.NotifyResponse")
	proto.RegisterType((*NotifyAndAdviseRequest)(nil), "spire.server.notifier.NotifyAndAdviseRequest")
	proto.RegisterType((*NotifyAndAdviseResponse)(nil), "spire.server.notifier.NotifyAndAdviseResponse")
}

func init() {
	proto.RegisterFile("spire/server/notifier/notifier.proto", fileDescriptor_c27428e9e6d193e9)
}

var fileDescriptor_c27428e9e6d193e9 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4f, 0x4f, 0xea, 0x40,
	0x14, 0xc5, 0x21, 0x2f, 0x8f, 0xf7, 0xde, 0x7d, 0x14, 0xcd, 0xc4, 0x7f, 0x74, 0x45, 0x2a, 0x18,
	0x35, 0xda, 0x26, 0x10, 0x77, 0xba, 0x00, 0x17, 0xa2, 0x51, 0x62, 0x9a, 0xb0, 0x61, 0x43, 0xa8,
	0xbd, 0xad, 0x93, 0xc0, 0x4c, 0xed, 0x1f, 0x12, 0x3f, 0x89, 0x5f, 0xd7, 0xd0, 0x3b, 0x55, 0x8a,
	0x08, 0xb8, 0x6a, 0x67, 0xce, 0xef, 0x9e, 0x73, 0xe7, 0x66, 0x06, 0xea, 0x51, 0xc0, 0x43, 0xb4,
	0x22, 0x0c, 0xa7, 0x18, 0x5a, 0x42, 0xc6, 0xdc, 0xe3, 0x73, 0x3f, 0x66, 0x10, 0xca, 0x58, 0xb2,
	0xdd, 0x94, 0x32, 0x89, 0x32, 0x33, 0x51, 0xaf, 0x52, 0xf1, 0x93, 0x9c, 0x4c, 0xa4, 0x50, 0x1f,
	0xaa, 0xd0, 0x6b, 0x39, 0x29, 0x18, 0x27, 0x3e, 0xcf, 0x3e, 0x44, 0x18, 0x97, 0x50, 0xee, 0x24,
	0xc2, 0x1d, 0xe3, 0xbd, 0x1c, 0xb9, 0xe8, 0xb2, 0x33, 0x28, 0x39, 0xe9, 0xfa, 0xa0, 0x58, 0x2b,
	0x1e, 0xff, 0x6f, 0xee, 0x98, 0x14, 0xaa, 0x6c, 0x89, 0xb5, 0x15, 0x63, 0x5c, 0x81, 0x46, 0x3b,
	0xfd, 0xc0, 0x1d, 0xc5, 0x3f, 0x2e, 0xf7, 0x41, 0xeb, 0xcd, 0x4e, 0xf1, 0x6a, 0xe3, 0x4b, 0x82,
	0x51, 0xcc, 0x1e, 0xa0, 0x42, 0xd2, 0x30, 0x21, 0x43, 0x65, 0x53, 0x37, 0x97, 0x1e, 0xdd, 0xcc,
	0x85, 0x77, 0x0b, 0xb6, 0xe6, 0xcc, 0x6f, 0x74, 0xfe, 0xc0, 0x6f, 0x9c, 0xa2, 0x88, 0x8d, 0x6d,
	0xa8, 0x64, 0x41, 0x51, 0x20, 0x45, 0x84, 0xc6, 0x04, 0xf6, 0x68, 0xa7, 0x2d, 0xdc, 0xb6, 0x3b,
	0xe5, 0x11, 0x66, 0x3d, 0xdc, 0x81, 0x72, 0x19, 0x8e, 0xd3, 0x91, 0xa8, 0x16, 0x0e, 0x57, 0xb6,
	0x40, 0xd3, 0xeb, 0x16, 0xec, 0xb2, 0x33, 0xb7, 0xfe, 0x6c, 0xa0, 0x0a, 0xfb, 0x5f, 0xe2, 0xa8,
	0x93, 0xe6, 0xdb, 0x2f, 0xf8, 0xdb, 0x53, 0x6e, 0xac, 0x0f, 0x25, 0xe2, 0xd8, 0x77, 0x47, 0xce,
	0x0d, 0x4c, 0x6f, 0xac, 0xa1, 0x28, 0x83, 0x05, 0xb0, 0xb5, 0x10, 0xcf, 0xce, 0x57, 0x56, 0x2e,
	0x4e, 0x45, 0x37, 0x37, 0xc5, 0x55, 0xe2, 0x00, 0xfe, 0x5d, 0x4b, 0xe1, 0x71, 0x3f, 0x09, 0x91,
	0x35, 0xf2, 0xb7, 0x40, 0x5d, 0xc0, 0x0f, 0x3d, 0xcb, 0x38, 0x5a, 0x87, 0x29, 0x6f, 0x0f, 0xb4,
	0x1b, 0x8c, 0x1f, 0x53, 0xf9, 0x56, 0x78, 0x92, 0x9d, 0x2c, 0x2d, 0xcc, 0x31, 0x59, 0xc6, 0xe9,
	0x26, 0x28, 0xe5, 0x74, 0x2e, 0x06, 0x2d, 0x9f, 0xc7, 0xcf, 0x89, 0x33, 0xa3, 0xad, 0x28, 0xe0,
	0x9e, 0x87, 0x16, 0xbd, 0xa8, 0xf4, 0xf1, 0x58, 0x4b, 0x5f, 0xad, 0x53, 0x4a, 0xc5, 0xd6, 0x7b,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x06, 0x0b, 0xf9, 0xd5, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NotifierClient is the client API for Notifier service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotifierClient interface {
	// Notify notifies the plugin that an event occurred. Errors returned by
	// the plugin are logged but otherwise ignored.
	Notify(ctx context.Context, in *NotifyRequest, opts ...grpc.CallOption) (*NotifyResponse, error)
	// NotifyAndAdvise notifies the plugin that an event occurred and waits
	// for a response. Errors returned by the plugin control SPIRE server
	// behavior. See NotifyAndAdviseRequest for per-event details.
	NotifyAndAdvise(ctx context.Context, in *NotifyAndAdviseRequest, opts ...grpc.CallOption) (*NotifyAndAdviseResponse, error)
	Configure(ctx context.Context, in *plugin.ConfigureRequest, opts ...grpc.CallOption) (*plugin.ConfigureResponse, error)
	GetPluginInfo(ctx context.Context, in *plugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*plugin.GetPluginInfoResponse, error)
}

type notifierClient struct {
	cc *grpc.ClientConn
}

func NewNotifierClient(cc *grpc.ClientConn) NotifierClient {
	return &notifierClient{cc}
}

func (c *notifierClient) Notify(ctx context.Context, in *NotifyRequest, opts ...grpc.CallOption) (*NotifyResponse, error) {
	out := new(NotifyResponse)
	err := c.cc.Invoke(ctx, "/spire.server.notifier.Notifier/Notify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifierClient) NotifyAndAdvise(ctx context.Context, in *NotifyAndAdviseRequest, opts ...grpc.CallOption) (*NotifyAndAdviseResponse, error) {
	out := new(NotifyAndAdviseResponse)
	err := c.cc.Invoke(ctx, "/spire.server.notifier.Notifier/NotifyAndAdvise", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifierClient) Configure(ctx context.Context, in *plugin.ConfigureRequest, opts ...grpc.CallOption) (*plugin.ConfigureResponse, error) {
	out := new(plugin.ConfigureResponse)
	err := c.cc.Invoke(ctx, "/spire.server.notifier.Notifier/Configure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifierClient) GetPluginInfo(ctx context.Context, in *plugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*plugin.GetPluginInfoResponse, error) {
	out := new(plugin.GetPluginInfoResponse)
	err := c.cc.Invoke(ctx, "/spire.server.notifier.Notifier/GetPluginInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotifierServer is the server API for Notifier service.
type NotifierServer interface {
	// Notify notifies the plugin that an event occurred. Errors returned by
	// the plugin are logged but otherwise ignored.
	Notify(context.Context, *NotifyRequest) (*NotifyResponse, error)
	// NotifyAndAdvise notifies the plugin that an event occurred and waits
	// for a response. Errors returned by the plugin control SPIRE server
	// behavior. See NotifyAndAdviseRequest for per-event details.
	NotifyAndAdvise(context.Context, *NotifyAndAdviseRequest) (*NotifyAndAdviseResponse, error)
	Configure(context.Context, *plugin.ConfigureRequest) (*plugin.ConfigureResponse, error)
	GetPluginInfo(context.Context, *plugin.GetPluginInfoRequest) (*plugin.GetPluginInfoResponse, error)
}

// UnimplementedNotifierServer can be embedded to have forward compatible implementations.
type UnimplementedNotifierServer struct {
}

func (*UnimplementedNotifierServer) Notify(ctx context.Context, req *NotifyRequest) (*NotifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Notify not implemented")
}
func (*UnimplementedNotifierServer) NotifyAndAdvise(ctx context.Context, req *NotifyAndAdviseRequest) (*NotifyAndAdviseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyAndAdvise not implemented")
}
func (*UnimplementedNotifierServer) Configure(ctx context.Context, req *plugin.ConfigureRequest) (*plugin.ConfigureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Configure not implemented")
}
func (*UnimplementedNotifierServer) GetPluginInfo(ctx context.Context, req *plugin.GetPluginInfoRequest) (*plugin.GetPluginInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPluginInfo not implemented")
}

func RegisterNotifierServer(s *grpc.Server, srv NotifierServer) {
	s.RegisterService(&_Notifier_serviceDesc, srv)
}

func _Notifier_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.server.notifier.Notifier/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).Notify(ctx, req.(*NotifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifier_NotifyAndAdvise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyAndAdviseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).NotifyAndAdvise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.server.notifier.Notifier/NotifyAndAdvise",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).NotifyAndAdvise(ctx, req.(*NotifyAndAdviseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifier_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(plugin.ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.server.notifier.Notifier/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).Configure(ctx, req.(*plugin.ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifier_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(plugin.GetPluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.server.notifier.Notifier/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).GetPluginInfo(ctx, req.(*plugin.GetPluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Notifier_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spire.server.notifier.Notifier",
	HandlerType: (*NotifierServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Notify",
			Handler:    _Notifier_Notify_Handler,
		},
		{
			MethodName: "NotifyAndAdvise",
			Handler:    _Notifier_NotifyAndAdvise_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _Notifier_Configure_Handler,
		},
		{
			MethodName: "GetPluginInfo",
			Handler:    _Notifier_GetPluginInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spire/server/notifier/notifier.proto",
}
