// Code generated by protoc-gen-go. DO NOT EDIT.
// source: jwtrpcdeclare.proto

package jwtrpcdeclare

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Algo int32

const (
	Algo_HS256 Algo = 0
	Algo_RS256 Algo = 1
)

var Algo_name = map[int32]string{
	0: "HS256",
	1: "RS256",
}

var Algo_value = map[string]int32{
	"HS256": 0,
	"RS256": 1,
}

func (x Algo) String() string {
	return proto.EnumName(Algo_name, int32(x))
}

func (Algo) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_de98355277429ad0, []int{0}
}

type StatusData_Status int32

const (
	StatusData_SUCCEED StatusData_Status = 0
	StatusData_ERROR   StatusData_Status = 1
)

var StatusData_Status_name = map[int32]string{
	0: "SUCCEED",
	1: "ERROR",
}

var StatusData_Status_value = map[string]int32{
	"SUCCEED": 0,
	"ERROR":   1,
}

func (x StatusData_Status) String() string {
	return proto.EnumName(StatusData_Status_name, int32(x))
}

func (StatusData_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_de98355277429ad0, []int{2, 0}
}

type SignJSONRequest struct {
	Algo                 Algo     `protobuf:"varint,1,opt,name=algo,proto3,enum=jwtrpcdeclare.Algo" json:"algo,omitempty"`
	Payload              []byte   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Aud                  string   `protobuf:"bytes,3,opt,name=aud,proto3" json:"aud,omitempty"`
	Iss                  string   `protobuf:"bytes,4,opt,name=iss,proto3" json:"iss,omitempty"`
	Exp                  int64    `protobuf:"varint,5,opt,name=exp,proto3" json:"exp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignJSONRequest) Reset()         { *m = SignJSONRequest{} }
func (m *SignJSONRequest) String() string { return proto.CompactTextString(m) }
func (*SignJSONRequest) ProtoMessage()    {}
func (*SignJSONRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_de98355277429ad0, []int{0}
}

func (m *SignJSONRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignJSONRequest.Unmarshal(m, b)
}
func (m *SignJSONRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignJSONRequest.Marshal(b, m, deterministic)
}
func (m *SignJSONRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignJSONRequest.Merge(m, src)
}
func (m *SignJSONRequest) XXX_Size() int {
	return xxx_messageInfo_SignJSONRequest.Size(m)
}
func (m *SignJSONRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignJSONRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignJSONRequest proto.InternalMessageInfo

func (m *SignJSONRequest) GetAlgo() Algo {
	if m != nil {
		return m.Algo
	}
	return Algo_HS256
}

func (m *SignJSONRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *SignJSONRequest) GetAud() string {
	if m != nil {
		return m.Aud
	}
	return ""
}

func (m *SignJSONRequest) GetIss() string {
	if m != nil {
		return m.Iss
	}
	return ""
}

func (m *SignJSONRequest) GetExp() int64 {
	if m != nil {
		return m.Exp
	}
	return 0
}

type VerifyRequest struct {
	Algo                 Algo     `protobuf:"varint,1,opt,name=algo,proto3,enum=jwtrpcdeclare.Algo" json:"algo,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyRequest) Reset()         { *m = VerifyRequest{} }
func (m *VerifyRequest) String() string { return proto.CompactTextString(m) }
func (*VerifyRequest) ProtoMessage()    {}
func (*VerifyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_de98355277429ad0, []int{1}
}

func (m *VerifyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyRequest.Unmarshal(m, b)
}
func (m *VerifyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyRequest.Marshal(b, m, deterministic)
}
func (m *VerifyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyRequest.Merge(m, src)
}
func (m *VerifyRequest) XXX_Size() int {
	return xxx_messageInfo_VerifyRequest.Size(m)
}
func (m *VerifyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyRequest proto.InternalMessageInfo

func (m *VerifyRequest) GetAlgo() Algo {
	if m != nil {
		return m.Algo
	}
	return Algo_HS256
}

func (m *VerifyRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type StatusData struct {
	Status               StatusData_Status `protobuf:"varint,1,opt,name=status,proto3,enum=jwtrpcdeclare.StatusData_Status" json:"status,omitempty"`
	Msg                  string            `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *StatusData) Reset()         { *m = StatusData{} }
func (m *StatusData) String() string { return proto.CompactTextString(m) }
func (*StatusData) ProtoMessage()    {}
func (*StatusData) Descriptor() ([]byte, []int) {
	return fileDescriptor_de98355277429ad0, []int{2}
}

func (m *StatusData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusData.Unmarshal(m, b)
}
func (m *StatusData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusData.Marshal(b, m, deterministic)
}
func (m *StatusData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusData.Merge(m, src)
}
func (m *StatusData) XXX_Size() int {
	return xxx_messageInfo_StatusData.Size(m)
}
func (m *StatusData) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusData.DiscardUnknown(m)
}

var xxx_messageInfo_StatusData proto.InternalMessageInfo

func (m *StatusData) GetStatus() StatusData_Status {
	if m != nil {
		return m.Status
	}
	return StatusData_SUCCEED
}

func (m *StatusData) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type SignResponse struct {
	Status               *StatusData `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Token                string      `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SignResponse) Reset()         { *m = SignResponse{} }
func (m *SignResponse) String() string { return proto.CompactTextString(m) }
func (*SignResponse) ProtoMessage()    {}
func (*SignResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_de98355277429ad0, []int{3}
}

func (m *SignResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignResponse.Unmarshal(m, b)
}
func (m *SignResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignResponse.Marshal(b, m, deterministic)
}
func (m *SignResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignResponse.Merge(m, src)
}
func (m *SignResponse) XXX_Size() int {
	return xxx_messageInfo_SignResponse.Size(m)
}
func (m *SignResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignResponse proto.InternalMessageInfo

func (m *SignResponse) GetStatus() *StatusData {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *SignResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type VerifyJSONResponse struct {
	Status               *StatusData `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Payload              []byte      `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *VerifyJSONResponse) Reset()         { *m = VerifyJSONResponse{} }
func (m *VerifyJSONResponse) String() string { return proto.CompactTextString(m) }
func (*VerifyJSONResponse) ProtoMessage()    {}
func (*VerifyJSONResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_de98355277429ad0, []int{4}
}

func (m *VerifyJSONResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyJSONResponse.Unmarshal(m, b)
}
func (m *VerifyJSONResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyJSONResponse.Marshal(b, m, deterministic)
}
func (m *VerifyJSONResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyJSONResponse.Merge(m, src)
}
func (m *VerifyJSONResponse) XXX_Size() int {
	return xxx_messageInfo_VerifyJSONResponse.Size(m)
}
func (m *VerifyJSONResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyJSONResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyJSONResponse proto.InternalMessageInfo

func (m *VerifyJSONResponse) GetStatus() *StatusData {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *VerifyJSONResponse) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterEnum("jwtrpcdeclare.Algo", Algo_name, Algo_value)
	proto.RegisterEnum("jwtrpcdeclare.StatusData_Status", StatusData_Status_name, StatusData_Status_value)
	proto.RegisterType((*SignJSONRequest)(nil), "jwtrpcdeclare.SignJSONRequest")
	proto.RegisterType((*VerifyRequest)(nil), "jwtrpcdeclare.VerifyRequest")
	proto.RegisterType((*StatusData)(nil), "jwtrpcdeclare.StatusData")
	proto.RegisterType((*SignResponse)(nil), "jwtrpcdeclare.SignResponse")
	proto.RegisterType((*VerifyJSONResponse)(nil), "jwtrpcdeclare.VerifyJSONResponse")
}

func init() { proto.RegisterFile("jwtrpcdeclare.proto", fileDescriptor_de98355277429ad0) }

var fileDescriptor_de98355277429ad0 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4f, 0x4f, 0xfa, 0x40,
	0x10, 0x65, 0x7f, 0xfc, 0xfb, 0x31, 0x80, 0x90, 0xc5, 0x43, 0x45, 0x62, 0xea, 0x5e, 0x6c, 0x3c,
	0x90, 0x58, 0xa3, 0xf1, 0xaa, 0x40, 0xa2, 0x1c, 0x20, 0xd9, 0x46, 0x4d, 0xbc, 0xad, 0xb0, 0x36,
	0xd5, 0xca, 0xd6, 0xee, 0x22, 0xe2, 0x47, 0xf0, 0x93, 0xf8, 0x31, 0xcd, 0xb6, 0x45, 0x2c, 0x81,
	0x83, 0xde, 0xde, 0xcc, 0xbe, 0xbc, 0x99, 0x37, 0x33, 0x0b, 0x8d, 0xc7, 0x99, 0x0a, 0x83, 0xd1,
	0x98, 0x8f, 0x7c, 0x16, 0xf2, 0x76, 0x10, 0x0a, 0x25, 0x70, 0x35, 0x95, 0x24, 0x1f, 0x08, 0x6a,
	0x8e, 0xe7, 0x4e, 0xfa, 0xce, 0x70, 0x40, 0xf9, 0xcb, 0x94, 0x4b, 0x85, 0x0f, 0x20, 0xc7, 0x7c,
	0x57, 0x18, 0xc8, 0x44, 0xd6, 0x96, 0xdd, 0x68, 0xa7, 0x65, 0xce, 0x7d, 0x57, 0xd0, 0x88, 0x80,
	0x0d, 0x28, 0x06, 0x6c, 0xee, 0x0b, 0x36, 0x36, 0xfe, 0x99, 0xc8, 0xaa, 0xd0, 0x45, 0x88, 0xeb,
	0x90, 0x65, 0xd3, 0xb1, 0x91, 0x35, 0x91, 0x55, 0xa2, 0x1a, 0xea, 0x8c, 0x27, 0xa5, 0x91, 0x8b,
	0x33, 0x9e, 0x94, 0x3a, 0xc3, 0xdf, 0x02, 0x23, 0x6f, 0x22, 0x2b, 0x4b, 0x35, 0x24, 0x03, 0xa8,
	0xde, 0xf0, 0xd0, 0x7b, 0x98, 0xff, 0xba, 0x93, 0x6d, 0xc8, 0x2b, 0xf1, 0xc4, 0x27, 0x51, 0x1f,
	0x25, 0x1a, 0x07, 0xe4, 0x1d, 0xc0, 0x51, 0x4c, 0x4d, 0x65, 0x97, 0x29, 0x86, 0xcf, 0xa0, 0x20,
	0xa3, 0x28, 0x91, 0x33, 0x57, 0xe4, 0x96, 0xd4, 0x04, 0xd2, 0x84, 0xaf, 0x3b, 0x7d, 0x96, 0x6e,
	0xa2, 0xad, 0x21, 0x31, 0xa1, 0x10, 0x73, 0x70, 0x19, 0x8a, 0xce, 0x75, 0xa7, 0xd3, 0xeb, 0x75,
	0xeb, 0x19, 0x5c, 0x82, 0x7c, 0x8f, 0xd2, 0x21, 0xad, 0x23, 0x72, 0x0b, 0x15, 0x3d, 0x57, 0xca,
	0x65, 0x20, 0x26, 0x92, 0xe3, 0xa3, 0x54, 0xf5, 0xb2, 0xbd, 0xb3, 0xb1, 0xfa, 0x77, 0xd9, 0xf5,
	0xa6, 0x18, 0xe0, 0x78, 0x48, 0xf1, 0xca, 0xfe, 0x2e, 0xbf, 0x71, 0x7b, 0x87, 0x2d, 0xc8, 0xe9,
	0xd9, 0x6a, 0x3b, 0x97, 0x8e, 0x7d, 0x72, 0x1a, 0x3b, 0xa3, 0x11, 0x44, 0xf6, 0x27, 0x02, 0xe8,
	0xcf, 0x94, 0xc3, 0xc3, 0x57, 0x6f, 0xc4, 0xf1, 0x15, 0xfc, 0x5f, 0x1c, 0x10, 0xde, 0x5b, 0xad,
	0x9a, 0xbe, 0xac, 0xe6, 0xee, 0x9a, 0xf7, 0x85, 0x05, 0x92, 0xc1, 0x43, 0x80, 0xa5, 0x35, 0xdc,
	0x5a, 0x21, 0xa7, 0x4e, 0xa3, 0xb9, 0xbf, 0xf6, 0xf5, 0xe7, 0x4c, 0x48, 0xe6, 0xa2, 0x76, 0x97,
	0x3e, 0xf7, 0xfb, 0x42, 0xf4, 0x09, 0x8e, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x19, 0xf8, 0xfd,
	0x2b, 0x1b, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// JwtServiceClient is the client API for JwtService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JwtServiceClient interface {
	// SignJSON 为JSON字符串签名
	SignJSON(ctx context.Context, in *SignJSONRequest, opts ...grpc.CallOption) (*SignResponse, error)
	// VerifyJSON 将token解析到的负载转码为JSON字符串返回
	VerifyJSON(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyJSONResponse, error)
}

type jwtServiceClient struct {
	cc *grpc.ClientConn
}

func NewJwtServiceClient(cc *grpc.ClientConn) JwtServiceClient {
	return &jwtServiceClient{cc}
}

func (c *jwtServiceClient) SignJSON(ctx context.Context, in *SignJSONRequest, opts ...grpc.CallOption) (*SignResponse, error) {
	out := new(SignResponse)
	err := c.cc.Invoke(ctx, "/jwtrpcdeclare.JwtService/SignJSON", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jwtServiceClient) VerifyJSON(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyJSONResponse, error) {
	out := new(VerifyJSONResponse)
	err := c.cc.Invoke(ctx, "/jwtrpcdeclare.JwtService/VerifyJSON", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JwtServiceServer is the server API for JwtService service.
type JwtServiceServer interface {
	// SignJSON 为JSON字符串签名
	SignJSON(context.Context, *SignJSONRequest) (*SignResponse, error)
	// VerifyJSON 将token解析到的负载转码为JSON字符串返回
	VerifyJSON(context.Context, *VerifyRequest) (*VerifyJSONResponse, error)
}

// UnimplementedJwtServiceServer can be embedded to have forward compatible implementations.
type UnimplementedJwtServiceServer struct {
}

func (*UnimplementedJwtServiceServer) SignJSON(ctx context.Context, req *SignJSONRequest) (*SignResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignJSON not implemented")
}
func (*UnimplementedJwtServiceServer) VerifyJSON(ctx context.Context, req *VerifyRequest) (*VerifyJSONResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyJSON not implemented")
}

func RegisterJwtServiceServer(s *grpc.Server, srv JwtServiceServer) {
	s.RegisterService(&_JwtService_serviceDesc, srv)
}

func _JwtService_SignJSON_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignJSONRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JwtServiceServer).SignJSON(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jwtrpcdeclare.JwtService/SignJSON",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JwtServiceServer).SignJSON(ctx, req.(*SignJSONRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JwtService_VerifyJSON_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JwtServiceServer).VerifyJSON(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jwtrpcdeclare.JwtService/VerifyJSON",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JwtServiceServer).VerifyJSON(ctx, req.(*VerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _JwtService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "jwtrpcdeclare.JwtService",
	HandlerType: (*JwtServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignJSON",
			Handler:    _JwtService_SignJSON_Handler,
		},
		{
			MethodName: "VerifyJSON",
			Handler:    _JwtService_VerifyJSON_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jwtrpcdeclare.proto",
}
