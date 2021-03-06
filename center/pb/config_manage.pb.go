// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config_manage.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	config_manage.proto
	forbid_manage.proto
	login_manage.proto
	notice_manage.proto
	player_manage.proto
	server_manage.proto
	trade_manage.proto

It has these top-level messages:
	ClientVersionRefreshRequest
	ClientVersionRefreshResponse
	ServerConfigRefreshRequest
	ServerConfigRefreshResponse
	PlatformConfigRefreshRequest
	PlatformConfigRefreshResponse
	ForbidIpListRequest
	ForbidIpListResponse
	ForbidIpSearchRequest
	ForbidIpSearchResponse
	ForbidIpRequest
	ForbidIpResponse
	ForbidUserRequest
	ForbidUserResponse
	ForbidUserSearchRequest
	ForbidUserSearchResponse
	SelfLoginRequest
	SelfLoginResponse
	LoginRequest
	LoginResponse
	LoginVerifyRequest
	LoginVerifyResponse
	GMLoginRequest
	GMLoginResponse
	NoticeRequest
	NoticeResponse
	RefreshNoticeRequest
	RefreshNoticeResponse
	PlayerServerInfo
	PlayerListRequest
	PlayerListResponse
	PlayerInfoSyncRequest
	PlayerInfoSyncResponse
	PlatformSetting
	PlatformChatSetting
	ServerRegisterRequest
	ServerRegisterResponse
	ServerUnregisterRequest
	ServerUnregisterResponse
	ServerCrossListRequest
	CrossServerInfo
	ServerCrossListResponse
	ServerInfo
	ServerInfoListRequest
	ServerInfoListResponse
	RefreshServerInfoListRequest
	RefreshServerInfoListResponse
	RefreshSDKListRequest
	RefreshSDKListResponse
	ServerInfoRequest
	ServerInfoResponse
	ServerInfoByPlatformRequest
	ServerInfoByPlatformResponse
	RefreshMarryPriceListRequest
	RefreshMarryPriceListResponse
	TradeServerListRequest
	TradeServerInfo
	TradeServerListResponse
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// 客户端配置刷新
type ClientVersionRefreshRequest struct {
}

func (m *ClientVersionRefreshRequest) Reset()                    { *m = ClientVersionRefreshRequest{} }
func (m *ClientVersionRefreshRequest) String() string            { return proto.CompactTextString(m) }
func (*ClientVersionRefreshRequest) ProtoMessage()               {}
func (*ClientVersionRefreshRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// 客户端配置刷新
type ClientVersionRefreshResponse struct {
}

func (m *ClientVersionRefreshResponse) Reset()                    { *m = ClientVersionRefreshResponse{} }
func (m *ClientVersionRefreshResponse) String() string            { return proto.CompactTextString(m) }
func (*ClientVersionRefreshResponse) ProtoMessage()               {}
func (*ClientVersionRefreshResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// 服务器配置刷新
type ServerConfigRefreshRequest struct {
}

func (m *ServerConfigRefreshRequest) Reset()                    { *m = ServerConfigRefreshRequest{} }
func (m *ServerConfigRefreshRequest) String() string            { return proto.CompactTextString(m) }
func (*ServerConfigRefreshRequest) ProtoMessage()               {}
func (*ServerConfigRefreshRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// 服务器配置刷新
type ServerConfigRefreshResponse struct {
}

func (m *ServerConfigRefreshResponse) Reset()                    { *m = ServerConfigRefreshResponse{} }
func (m *ServerConfigRefreshResponse) String() string            { return proto.CompactTextString(m) }
func (*ServerConfigRefreshResponse) ProtoMessage()               {}
func (*ServerConfigRefreshResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// 刷新平台配置
type PlatformConfigRefreshRequest struct {
	Platform int32 `protobuf:"varint,1,opt,name=platform" json:"platform,omitempty"`
}

func (m *PlatformConfigRefreshRequest) Reset()                    { *m = PlatformConfigRefreshRequest{} }
func (m *PlatformConfigRefreshRequest) String() string            { return proto.CompactTextString(m) }
func (*PlatformConfigRefreshRequest) ProtoMessage()               {}
func (*PlatformConfigRefreshRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PlatformConfigRefreshRequest) GetPlatform() int32 {
	if m != nil {
		return m.Platform
	}
	return 0
}

// 刷新平台配置
type PlatformConfigRefreshResponse struct {
	Platform int32 `protobuf:"varint,1,opt,name=platform" json:"platform,omitempty"`
}

func (m *PlatformConfigRefreshResponse) Reset()                    { *m = PlatformConfigRefreshResponse{} }
func (m *PlatformConfigRefreshResponse) String() string            { return proto.CompactTextString(m) }
func (*PlatformConfigRefreshResponse) ProtoMessage()               {}
func (*PlatformConfigRefreshResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PlatformConfigRefreshResponse) GetPlatform() int32 {
	if m != nil {
		return m.Platform
	}
	return 0
}

func init() {
	proto.RegisterType((*ClientVersionRefreshRequest)(nil), "pb.ClientVersionRefreshRequest")
	proto.RegisterType((*ClientVersionRefreshResponse)(nil), "pb.ClientVersionRefreshResponse")
	proto.RegisterType((*ServerConfigRefreshRequest)(nil), "pb.ServerConfigRefreshRequest")
	proto.RegisterType((*ServerConfigRefreshResponse)(nil), "pb.ServerConfigRefreshResponse")
	proto.RegisterType((*PlatformConfigRefreshRequest)(nil), "pb.PlatformConfigRefreshRequest")
	proto.RegisterType((*PlatformConfigRefreshResponse)(nil), "pb.PlatformConfigRefreshResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ConfigManage service

type ConfigManageClient interface {
	// 客户端版本刷新
	RefreshClientVersion(ctx context.Context, in *ClientVersionRefreshRequest, opts ...grpc.CallOption) (*ClientVersionRefreshResponse, error)
	// 服务器配置刷新
	RefreshServerConfig(ctx context.Context, in *ServerConfigRefreshRequest, opts ...grpc.CallOption) (*ServerConfigRefreshResponse, error)
	// 刷新平台配置
	RefreshPlatformConfig(ctx context.Context, in *PlatformConfigRefreshRequest, opts ...grpc.CallOption) (*PlatformConfigRefreshResponse, error)
}

type configManageClient struct {
	cc *grpc.ClientConn
}

func NewConfigManageClient(cc *grpc.ClientConn) ConfigManageClient {
	return &configManageClient{cc}
}

func (c *configManageClient) RefreshClientVersion(ctx context.Context, in *ClientVersionRefreshRequest, opts ...grpc.CallOption) (*ClientVersionRefreshResponse, error) {
	out := new(ClientVersionRefreshResponse)
	err := grpc.Invoke(ctx, "/pb.ConfigManage/RefreshClientVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configManageClient) RefreshServerConfig(ctx context.Context, in *ServerConfigRefreshRequest, opts ...grpc.CallOption) (*ServerConfigRefreshResponse, error) {
	out := new(ServerConfigRefreshResponse)
	err := grpc.Invoke(ctx, "/pb.ConfigManage/RefreshServerConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configManageClient) RefreshPlatformConfig(ctx context.Context, in *PlatformConfigRefreshRequest, opts ...grpc.CallOption) (*PlatformConfigRefreshResponse, error) {
	out := new(PlatformConfigRefreshResponse)
	err := grpc.Invoke(ctx, "/pb.ConfigManage/RefreshPlatformConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ConfigManage service

type ConfigManageServer interface {
	// 客户端版本刷新
	RefreshClientVersion(context.Context, *ClientVersionRefreshRequest) (*ClientVersionRefreshResponse, error)
	// 服务器配置刷新
	RefreshServerConfig(context.Context, *ServerConfigRefreshRequest) (*ServerConfigRefreshResponse, error)
	// 刷新平台配置
	RefreshPlatformConfig(context.Context, *PlatformConfigRefreshRequest) (*PlatformConfigRefreshResponse, error)
}

func RegisterConfigManageServer(s *grpc.Server, srv ConfigManageServer) {
	s.RegisterService(&_ConfigManage_serviceDesc, srv)
}

func _ConfigManage_RefreshClientVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientVersionRefreshRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigManageServer).RefreshClientVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ConfigManage/RefreshClientVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigManageServer).RefreshClientVersion(ctx, req.(*ClientVersionRefreshRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigManage_RefreshServerConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerConfigRefreshRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigManageServer).RefreshServerConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ConfigManage/RefreshServerConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigManageServer).RefreshServerConfig(ctx, req.(*ServerConfigRefreshRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigManage_RefreshPlatformConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlatformConfigRefreshRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigManageServer).RefreshPlatformConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ConfigManage/RefreshPlatformConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigManageServer).RefreshPlatformConfig(ctx, req.(*PlatformConfigRefreshRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConfigManage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ConfigManage",
	HandlerType: (*ConfigManageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RefreshClientVersion",
			Handler:    _ConfigManage_RefreshClientVersion_Handler,
		},
		{
			MethodName: "RefreshServerConfig",
			Handler:    _ConfigManage_RefreshServerConfig_Handler,
		},
		{
			MethodName: "RefreshPlatformConfig",
			Handler:    _ConfigManage_RefreshPlatformConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "config_manage.proto",
}

func init() { proto.RegisterFile("config_manage.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xce, 0xcf, 0x4b,
	0xcb, 0x4c, 0x8f, 0xcf, 0x4d, 0xcc, 0x4b, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x2a, 0x48, 0x52, 0x92, 0xe5, 0x92, 0x76, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0x09, 0x4b, 0x2d,
	0x2a, 0xce, 0xcc, 0xcf, 0x0b, 0x4a, 0x4d, 0x2b, 0x4a, 0x2d, 0xce, 0x08, 0x4a, 0x2d, 0x2c, 0x4d,
	0x2d, 0x2e, 0x51, 0x92, 0xe3, 0x92, 0xc1, 0x2e, 0x5d, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0xaa, 0x24,
	0xc3, 0x25, 0x15, 0x9c, 0x5a, 0x54, 0x96, 0x5a, 0xe4, 0x0c, 0x36, 0x1f, 0x4d, 0xb7, 0x2c, 0x97,
	0x34, 0x56, 0x59, 0xa8, 0x66, 0x2b, 0x2e, 0x99, 0x80, 0x9c, 0xc4, 0x92, 0xb4, 0xfc, 0xa2, 0x5c,
	0x6c, 0xda, 0x85, 0xa4, 0xb8, 0x38, 0x0a, 0xa0, 0xf2, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41,
	0x70, 0xbe, 0x92, 0x35, 0x97, 0x2c, 0x0e, 0xbd, 0x10, 0xc3, 0xf1, 0x69, 0x36, 0x5a, 0xca, 0xc4,
	0xc5, 0x03, 0xd1, 0xe5, 0x0b, 0x0e, 0x0f, 0xa1, 0x68, 0x2e, 0x11, 0xa8, 0x7e, 0x14, 0xdf, 0x0a,
	0xc9, 0xeb, 0x15, 0x24, 0xe9, 0xe1, 0x09, 0x1f, 0x29, 0x05, 0xdc, 0x0a, 0xa0, 0x9e, 0x64, 0x10,
	0x8a, 0xe0, 0x12, 0x86, 0x0a, 0x22, 0x07, 0x86, 0x90, 0x1c, 0x48, 0x2b, 0xee, 0xc0, 0x93, 0x92,
	0xc7, 0x29, 0x0f, 0x37, 0x39, 0x8e, 0x4b, 0x14, 0x2a, 0x88, 0x1a, 0x16, 0x42, 0x60, 0x67, 0xe1,
	0x0b, 0x5b, 0x29, 0x45, 0x3c, 0x2a, 0x60, 0xe6, 0x27, 0xb1, 0x81, 0xd3, 0x89, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0x63, 0x4d, 0x58, 0xdc, 0x3e, 0x02, 0x00, 0x00,
}
