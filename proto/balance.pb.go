// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/balance.proto

package account

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Balance struct {
	Balance              string   `protobuf:"bytes,1,opt,name=balance,proto3" json:"balance,omitempty"`
	TotalBalance         string   `protobuf:"bytes,2,opt,name=total_balance,json=totalBalance,proto3" json:"total_balance,omitempty"`
	Currency             string   `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	SpendToday           string   `protobuf:"bytes,4,opt,name=spend_today,json=spendToday,proto3" json:"spend_today,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Balance) Reset()         { *m = Balance{} }
func (m *Balance) String() string { return proto.CompactTextString(m) }
func (*Balance) ProtoMessage()    {}
func (*Balance) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab57c80074bdde37, []int{0}
}

func (m *Balance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Balance.Unmarshal(m, b)
}
func (m *Balance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Balance.Marshal(b, m, deterministic)
}
func (m *Balance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Balance.Merge(m, src)
}
func (m *Balance) XXX_Size() int {
	return xxx_messageInfo_Balance.Size(m)
}
func (m *Balance) XXX_DiscardUnknown() {
	xxx_messageInfo_Balance.DiscardUnknown(m)
}

var xxx_messageInfo_Balance proto.InternalMessageInfo

func (m *Balance) GetBalance() string {
	if m != nil {
		return m.Balance
	}
	return ""
}

func (m *Balance) GetTotalBalance() string {
	if m != nil {
		return m.TotalBalance
	}
	return ""
}

func (m *Balance) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *Balance) GetSpendToday() string {
	if m != nil {
		return m.SpendToday
	}
	return ""
}

type BalanceRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BalanceRequest) Reset()         { *m = BalanceRequest{} }
func (m *BalanceRequest) String() string { return proto.CompactTextString(m) }
func (*BalanceRequest) ProtoMessage()    {}
func (*BalanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab57c80074bdde37, []int{1}
}

func (m *BalanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BalanceRequest.Unmarshal(m, b)
}
func (m *BalanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BalanceRequest.Marshal(b, m, deterministic)
}
func (m *BalanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceRequest.Merge(m, src)
}
func (m *BalanceRequest) XXX_Size() int {
	return xxx_messageInfo_BalanceRequest.Size(m)
}
func (m *BalanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceRequest proto.InternalMessageInfo

type BalanceResponse struct {
	Balance              *Balance `protobuf:"bytes,1,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BalanceResponse) Reset()         { *m = BalanceResponse{} }
func (m *BalanceResponse) String() string { return proto.CompactTextString(m) }
func (*BalanceResponse) ProtoMessage()    {}
func (*BalanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab57c80074bdde37, []int{2}
}

func (m *BalanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BalanceResponse.Unmarshal(m, b)
}
func (m *BalanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BalanceResponse.Marshal(b, m, deterministic)
}
func (m *BalanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceResponse.Merge(m, src)
}
func (m *BalanceResponse) XXX_Size() int {
	return xxx_messageInfo_BalanceResponse.Size(m)
}
func (m *BalanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceResponse proto.InternalMessageInfo

func (m *BalanceResponse) GetBalance() *Balance {
	if m != nil {
		return m.Balance
	}
	return nil
}

func init() {
	proto.RegisterType((*Balance)(nil), "account.Balance")
	proto.RegisterType((*BalanceRequest)(nil), "account.BalanceRequest")
	proto.RegisterType((*BalanceResponse)(nil), "account.BalanceResponse")
}

func init() { proto.RegisterFile("proto/balance.proto", fileDescriptor_ab57c80074bdde37) }

var fileDescriptor_ab57c80074bdde37 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x50, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0x26, 0x55, 0x8c, 0x4e, 0xd5, 0xd6, 0xf5, 0xd0, 0x25, 0x08, 0x4a, 0xbc, 0x88, 0x87, 0x04,
	0xeb, 0xd9, 0x8b, 0x8f, 0x50, 0xf5, 0x5c, 0xb6, 0xdb, 0x21, 0x04, 0xc2, 0x4e, 0xcc, 0x4e, 0x0a,
	0xbd, 0x0a, 0x3e, 0x81, 0x8f, 0xe6, 0x2b, 0xf8, 0x20, 0xe2, 0x24, 0x1b, 0x29, 0x3d, 0x7e, 0xbf,
	0xbb, 0xf3, 0xc1, 0x65, 0xdd, 0x10, 0x53, 0xbe, 0x32, 0x95, 0x71, 0x16, 0x33, 0x41, 0x2a, 0x36,
	0xd6, 0x52, 0xeb, 0x38, 0xb9, 0x2a, 0x88, 0x8a, 0x0a, 0x73, 0x53, 0x97, 0xb9, 0x71, 0x8e, 0xd8,
	0x70, 0x49, 0xce, 0x77, 0xb6, 0xf4, 0x33, 0x82, 0xf8, 0xb9, 0x0b, 0x2a, 0x0d, 0x71, 0xdf, 0xa1,
	0xa3, 0x9b, 0xe8, 0xee, 0x64, 0x11, 0xa0, 0xba, 0x85, 0x33, 0x26, 0x36, 0xd5, 0x32, 0xe8, 0x23,
	0xd1, 0x4f, 0x85, 0x0c, 0xf1, 0x04, 0x8e, 0x6d, 0xdb, 0x34, 0xe8, 0xec, 0x56, 0x1f, 0x88, 0x3e,
	0x60, 0x75, 0x0d, 0x63, 0x5f, 0xa3, 0x5b, 0x2f, 0x99, 0xd6, 0x66, 0xab, 0x0f, 0x45, 0x06, 0xa1,
	0x5e, 0xff, 0x98, 0x74, 0x0a, 0xe7, 0x7d, 0xcf, 0x02, 0xdf, 0x5b, 0xf4, 0x9c, 0x3e, 0xc1, 0x64,
	0x60, 0x7c, 0x4d, 0xce, 0xa3, 0xba, 0xdf, 0xfd, 0xe0, 0x78, 0x3e, 0xcd, 0xfa, 0x2b, 0xb3, 0x60,
	0x0d, 0x86, 0x79, 0x31, 0x14, 0xbe, 0x60, 0xb3, 0x29, 0x2d, 0xaa, 0xb7, 0xff, 0x4b, 0x67, 0x7b,
	0xb9, 0xee, 0xd1, 0x44, 0xef, 0x0b, 0xdd, 0xdb, 0xe9, 0xec, 0xe3, 0xfb, 0xe7, 0x6b, 0x74, 0xa1,
	0x26, 0x32, 0xe4, 0xe6, 0x21, 0xcc, 0xbd, 0x3a, 0x92, 0x21, 0x1f, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x6e, 0xb1, 0x83, 0x18, 0x86, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BalanceServiceClient is the client API for BalanceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BalanceServiceClient interface {
	Balance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error)
}

type balanceServiceClient struct {
	cc *grpc.ClientConn
}

func NewBalanceServiceClient(cc *grpc.ClientConn) BalanceServiceClient {
	return &balanceServiceClient{cc}
}

func (c *balanceServiceClient) Balance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error) {
	out := new(BalanceResponse)
	err := c.cc.Invoke(ctx, "/account.BalanceService/Balance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BalanceServiceServer is the server API for BalanceService service.
type BalanceServiceServer interface {
	Balance(context.Context, *BalanceRequest) (*BalanceResponse, error)
}

// UnimplementedBalanceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBalanceServiceServer struct {
}

func (*UnimplementedBalanceServiceServer) Balance(ctx context.Context, req *BalanceRequest) (*BalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Balance not implemented")
}

func RegisterBalanceServiceServer(s *grpc.Server, srv BalanceServiceServer) {
	s.RegisterService(&_BalanceService_serviceDesc, srv)
}

func _BalanceService_Balance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalanceServiceServer).Balance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.BalanceService/Balance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalanceServiceServer).Balance(ctx, req.(*BalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BalanceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "account.BalanceService",
	HandlerType: (*BalanceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Balance",
			Handler:    _BalanceService_Balance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/balance.proto",
}
