// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rook/game/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgMove struct {
	Creator    string    `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	GameId     uint64    `protobuf:"varint,2,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	Position   *Position `protobuf:"bytes,3,opt,name=position,proto3" json:"position,omitempty"`
	Direction  Direction `protobuf:"varint,4,opt,name=direction,proto3,enum=rook.game.Direction" json:"direction,omitempty"`
	Population uint32    `protobuf:"varint,5,opt,name=population,proto3" json:"population,omitempty"`
}

func (m *MsgMove) Reset()         { *m = MsgMove{} }
func (m *MsgMove) String() string { return proto.CompactTextString(m) }
func (*MsgMove) ProtoMessage()    {}
func (*MsgMove) Descriptor() ([]byte, []int) {
	return fileDescriptor_18171b8d3f78c24d, []int{0}
}
func (m *MsgMove) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMove) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMove.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMove) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMove.Merge(m, src)
}
func (m *MsgMove) XXX_Size() int {
	return m.Size()
}
func (m *MsgMove) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMove.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMove proto.InternalMessageInfo

func (m *MsgMove) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgMove) GetGameId() uint64 {
	if m != nil {
		return m.GameId
	}
	return 0
}

func (m *MsgMove) GetPosition() *Position {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *MsgMove) GetDirection() Direction {
	if m != nil {
		return m.Direction
	}
	return Direction_LEFT
}

func (m *MsgMove) GetPopulation() uint32 {
	if m != nil {
		return m.Population
	}
	return 0
}

type MsgMoveResponse struct {
}

func (m *MsgMoveResponse) Reset()         { *m = MsgMoveResponse{} }
func (m *MsgMoveResponse) String() string { return proto.CompactTextString(m) }
func (*MsgMoveResponse) ProtoMessage()    {}
func (*MsgMoveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_18171b8d3f78c24d, []int{1}
}
func (m *MsgMoveResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMoveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMoveResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMoveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMoveResponse.Merge(m, src)
}
func (m *MsgMoveResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgMoveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMoveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMoveResponse proto.InternalMessageInfo

type MsgBuild struct {
	Creator    string     `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	GameId     uint64     `protobuf:"varint,2,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	Settlement Settlement `protobuf:"varint,3,opt,name=settlement,proto3,enum=rook.game.Settlement" json:"settlement,omitempty"`
	Position   *Position  `protobuf:"bytes,4,opt,name=position,proto3" json:"position,omitempty"`
}

func (m *MsgBuild) Reset()         { *m = MsgBuild{} }
func (m *MsgBuild) String() string { return proto.CompactTextString(m) }
func (*MsgBuild) ProtoMessage()    {}
func (*MsgBuild) Descriptor() ([]byte, []int) {
	return fileDescriptor_18171b8d3f78c24d, []int{2}
}
func (m *MsgBuild) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBuild) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBuild.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBuild) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBuild.Merge(m, src)
}
func (m *MsgBuild) XXX_Size() int {
	return m.Size()
}
func (m *MsgBuild) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBuild.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBuild proto.InternalMessageInfo

func (m *MsgBuild) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgBuild) GetGameId() uint64 {
	if m != nil {
		return m.GameId
	}
	return 0
}

func (m *MsgBuild) GetSettlement() Settlement {
	if m != nil {
		return m.Settlement
	}
	return Settlement_NONE
}

func (m *MsgBuild) GetPosition() *Position {
	if m != nil {
		return m.Position
	}
	return nil
}

type MsgBuildResponse struct {
}

func (m *MsgBuildResponse) Reset()         { *m = MsgBuildResponse{} }
func (m *MsgBuildResponse) String() string { return proto.CompactTextString(m) }
func (*MsgBuildResponse) ProtoMessage()    {}
func (*MsgBuildResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_18171b8d3f78c24d, []int{3}
}
func (m *MsgBuildResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBuildResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBuildResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBuildResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBuildResponse.Merge(m, src)
}
func (m *MsgBuildResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgBuildResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBuildResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBuildResponse proto.InternalMessageInfo

type MsgCreate struct {
	Players []string    `protobuf:"bytes,2,rep,name=players,proto3" json:"players,omitempty"`
	Config  *GameConfig `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
}

func (m *MsgCreate) Reset()         { *m = MsgCreate{} }
func (m *MsgCreate) String() string { return proto.CompactTextString(m) }
func (*MsgCreate) ProtoMessage()    {}
func (*MsgCreate) Descriptor() ([]byte, []int) {
	return fileDescriptor_18171b8d3f78c24d, []int{4}
}
func (m *MsgCreate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreate.Merge(m, src)
}
func (m *MsgCreate) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreate) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreate.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreate proto.InternalMessageInfo

func (m *MsgCreate) GetPlayers() []string {
	if m != nil {
		return m.Players
	}
	return nil
}

func (m *MsgCreate) GetConfig() *GameConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

type MsgCreateResponse struct {
	GameId uint64 `protobuf:"varint,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
}

func (m *MsgCreateResponse) Reset()         { *m = MsgCreateResponse{} }
func (m *MsgCreateResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateResponse) ProtoMessage()    {}
func (*MsgCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_18171b8d3f78c24d, []int{5}
}
func (m *MsgCreateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateResponse.Merge(m, src)
}
func (m *MsgCreateResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateResponse proto.InternalMessageInfo

func (m *MsgCreateResponse) GetGameId() uint64 {
	if m != nil {
		return m.GameId
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgMove)(nil), "rook.game.MsgMove")
	proto.RegisterType((*MsgMoveResponse)(nil), "rook.game.MsgMoveResponse")
	proto.RegisterType((*MsgBuild)(nil), "rook.game.MsgBuild")
	proto.RegisterType((*MsgBuildResponse)(nil), "rook.game.MsgBuildResponse")
	proto.RegisterType((*MsgCreate)(nil), "rook.game.MsgCreate")
	proto.RegisterType((*MsgCreateResponse)(nil), "rook.game.MsgCreateResponse")
}

func init() { proto.RegisterFile("rook/game/tx.proto", fileDescriptor_18171b8d3f78c24d) }

var fileDescriptor_18171b8d3f78c24d = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4d, 0x8b, 0xd4, 0x40,
	0x10, 0x9d, 0xde, 0xc9, 0x66, 0x36, 0x25, 0xae, 0x6e, 0xef, 0xae, 0x86, 0x28, 0x21, 0x04, 0x84,
	0x1c, 0x34, 0x81, 0x88, 0x82, 0x20, 0x08, 0xbb, 0x82, 0x78, 0x08, 0x48, 0xeb, 0xc9, 0x8b, 0x64,
	0x93, 0x36, 0x06, 0x93, 0x74, 0x48, 0xf7, 0xe8, 0xcc, 0xbf, 0xf0, 0x4f, 0xf8, 0x17, 0xbc, 0x7b,
	0xf3, 0x38, 0x47, 0x8f, 0x32, 0xf3, 0x47, 0x24, 0x9d, 0x8f, 0xe9, 0xc1, 0x39, 0xc8, 0x9e, 0x86,
	0x7a, 0xf5, 0x5e, 0xcd, 0x7b, 0xd5, 0x29, 0xc0, 0x0d, 0x63, 0x9f, 0x83, 0x2c, 0x2e, 0x69, 0x20,
	0x16, 0x7e, 0xdd, 0x30, 0xc1, 0xb0, 0xd1, 0x62, 0x7e, 0x8b, 0x59, 0xe7, 0x4a, 0x7b, 0x59, 0x53,
	0xde, 0x31, 0xac, 0x3b, 0x5b, 0x38, 0x61, 0xd5, 0xc7, 0x3c, 0xeb, 0x70, 0xf7, 0x27, 0x82, 0x59,
	0xc4, 0xb3, 0x88, 0x7d, 0xa1, 0xd8, 0x84, 0x59, 0xd2, 0xd0, 0x58, 0xb0, 0xc6, 0x44, 0x0e, 0xf2,
	0x0c, 0x32, 0x94, 0xf8, 0x2e, 0xcc, 0x5a, 0xe9, 0x87, 0x3c, 0x35, 0x0f, 0x1c, 0xe4, 0x69, 0x44,
	0x6f, 0xcb, 0xd7, 0x29, 0x0e, 0xe0, 0xa8, 0x66, 0x3c, 0x17, 0x39, 0xab, 0xcc, 0xa9, 0x83, 0xbc,
	0x1b, 0xe1, 0xa9, 0x3f, 0x7a, 0xf1, 0xdf, 0xf4, 0x2d, 0x32, 0x92, 0x70, 0x08, 0x46, 0x9a, 0x37,
	0x34, 0x91, 0x0a, 0xcd, 0x41, 0xde, 0x71, 0x78, 0xa6, 0x28, 0x5e, 0x0e, 0x3d, 0xb2, 0xa5, 0x61,
	0x1b, 0xa0, 0x66, 0xf5, 0xbc, 0x88, 0xa5, 0xe8, 0xd0, 0x41, 0xde, 0x4d, 0xa2, 0x20, 0xee, 0x09,
	0xdc, 0xea, 0x23, 0x10, 0xca, 0x6b, 0x56, 0x71, 0xea, 0x7e, 0x47, 0x70, 0x14, 0xf1, 0xec, 0x62,
	0x9e, 0x17, 0xe9, 0x75, 0x72, 0x3d, 0x01, 0xe0, 0x54, 0x88, 0x82, 0x96, 0xb4, 0x12, 0x32, 0xd9,
	0x71, 0x78, 0xae, 0xf8, 0x7c, 0x3b, 0x36, 0x89, 0x42, 0xdc, 0x59, 0x87, 0xf6, 0x1f, 0xeb, 0x70,
	0x31, 0xdc, 0x1e, 0x6c, 0x8e, 0xde, 0xdf, 0x81, 0x11, 0xf1, 0xec, 0xb2, 0xb5, 0x28, 0xdf, 0xa4,
	0x2e, 0xe2, 0x25, 0x6d, 0xb8, 0x79, 0xe0, 0x4c, 0x5b, 0xef, 0x7d, 0x89, 0x1f, 0x81, 0xde, 0xbd,
	0x64, 0xbf, 0x78, 0xd5, 0xde, 0xab, 0xb8, 0xa4, 0x97, 0xb2, 0x49, 0x7a, 0x92, 0xfb, 0x10, 0x4e,
	0xc6, 0xa9, 0xc3, 0x5f, 0xa9, 0xf9, 0x91, 0x9a, 0x3f, 0xfc, 0x81, 0x60, 0x1a, 0xf1, 0x0c, 0x3f,
	0x05, 0x4d, 0x7e, 0x1a, 0x58, 0x19, 0xde, 0xef, 0xda, 0xb2, 0xfe, 0xc5, 0xc6, 0xc1, 0xcf, 0xe0,
	0xb0, 0xdb, 0xfd, 0xe9, 0x2e, 0x49, 0x82, 0xd6, 0xbd, 0x3d, 0xe0, 0x28, 0x7d, 0x0e, 0x7a, 0x9f,
	0xfd, 0x6c, 0x97, 0xd6, 0xa1, 0xd6, 0xfd, 0x7d, 0xe8, 0xa0, 0xbe, 0x78, 0xf1, 0x6b, 0x6d, 0xa3,
	0xd5, 0xda, 0x46, 0x7f, 0xd6, 0x36, 0xfa, 0xb6, 0xb1, 0x27, 0xab, 0x8d, 0x3d, 0xf9, 0xbd, 0xb1,
	0x27, 0xef, 0x1f, 0x64, 0xb9, 0xf8, 0x34, 0xbf, 0xf2, 0x13, 0x56, 0x06, 0x49, 0xf9, 0x35, 0x16,
	0xb4, 0xe1, 0x81, 0xbc, 0x8a, 0x45, 0xf7, 0x23, 0xcf, 0xe5, 0x4a, 0x97, 0x77, 0xf1, 0xf8, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x5a, 0x37, 0xe5, 0x67, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// this line is used by starport scaffolding # proto/tx/rpc
	Move(ctx context.Context, in *MsgMove, opts ...grpc.CallOption) (*MsgMoveResponse, error)
	Build(ctx context.Context, in *MsgBuild, opts ...grpc.CallOption) (*MsgBuildResponse, error)
	Create(ctx context.Context, in *MsgCreate, opts ...grpc.CallOption) (*MsgCreateResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Move(ctx context.Context, in *MsgMove, opts ...grpc.CallOption) (*MsgMoveResponse, error) {
	out := new(MsgMoveResponse)
	err := c.cc.Invoke(ctx, "/rook.game.Msg/Move", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Build(ctx context.Context, in *MsgBuild, opts ...grpc.CallOption) (*MsgBuildResponse, error) {
	out := new(MsgBuildResponse)
	err := c.cc.Invoke(ctx, "/rook.game.Msg/Build", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Create(ctx context.Context, in *MsgCreate, opts ...grpc.CallOption) (*MsgCreateResponse, error) {
	out := new(MsgCreateResponse)
	err := c.cc.Invoke(ctx, "/rook.game.Msg/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// this line is used by starport scaffolding # proto/tx/rpc
	Move(context.Context, *MsgMove) (*MsgMoveResponse, error)
	Build(context.Context, *MsgBuild) (*MsgBuildResponse, error)
	Create(context.Context, *MsgCreate) (*MsgCreateResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Move(ctx context.Context, req *MsgMove) (*MsgMoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Move not implemented")
}
func (*UnimplementedMsgServer) Build(ctx context.Context, req *MsgBuild) (*MsgBuildResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Build not implemented")
}
func (*UnimplementedMsgServer) Create(ctx context.Context, req *MsgCreate) (*MsgCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Move_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMove)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Move(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rook.game.Msg/Move",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Move(ctx, req.(*MsgMove))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Build_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBuild)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Build(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rook.game.Msg/Build",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Build(ctx, req.(*MsgBuild))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rook.game.Msg/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Create(ctx, req.(*MsgCreate))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rook.game.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Move",
			Handler:    _Msg_Move_Handler,
		},
		{
			MethodName: "Build",
			Handler:    _Msg_Build_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Msg_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rook/game/tx.proto",
}

func (m *MsgMove) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMove) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMove) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Population != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Population))
		i--
		dAtA[i] = 0x28
	}
	if m.Direction != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Direction))
		i--
		dAtA[i] = 0x20
	}
	if m.Position != nil {
		{
			size, err := m.Position.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.GameId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.GameId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgMoveResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMoveResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMoveResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgBuild) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBuild) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBuild) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Position != nil {
		{
			size, err := m.Position.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Settlement != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Settlement))
		i--
		dAtA[i] = 0x18
	}
	if m.GameId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.GameId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgBuildResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBuildResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBuildResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgCreate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Config != nil {
		{
			size, err := m.Config.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Players) > 0 {
		for iNdEx := len(m.Players) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Players[iNdEx])
			copy(dAtA[i:], m.Players[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Players[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GameId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.GameId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgMove) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.GameId != 0 {
		n += 1 + sovTx(uint64(m.GameId))
	}
	if m.Position != nil {
		l = m.Position.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Direction != 0 {
		n += 1 + sovTx(uint64(m.Direction))
	}
	if m.Population != 0 {
		n += 1 + sovTx(uint64(m.Population))
	}
	return n
}

func (m *MsgMoveResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgBuild) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.GameId != 0 {
		n += 1 + sovTx(uint64(m.GameId))
	}
	if m.Settlement != 0 {
		n += 1 + sovTx(uint64(m.Settlement))
	}
	if m.Position != nil {
		l = m.Position.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgBuildResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgCreate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Players) > 0 {
		for _, s := range m.Players {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if m.Config != nil {
		l = m.Config.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GameId != 0 {
		n += 1 + sovTx(uint64(m.GameId))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgMove) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgMove: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMove: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GameId", wireType)
			}
			m.GameId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GameId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Position", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Position == nil {
				m.Position = &Position{}
			}
			if err := m.Position.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Direction", wireType)
			}
			m.Direction = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Direction |= Direction(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Population", wireType)
			}
			m.Population = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Population |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgMoveResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgMoveResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMoveResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgBuild) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgBuild: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBuild: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GameId", wireType)
			}
			m.GameId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GameId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Settlement", wireType)
			}
			m.Settlement = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Settlement |= Settlement(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Position", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Position == nil {
				m.Position = &Position{}
			}
			if err := m.Position.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgBuildResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgBuildResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBuildResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCreate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Players", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Players = append(m.Players, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Config == nil {
				m.Config = &GameConfig{}
			}
			if err := m.Config.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCreateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GameId", wireType)
			}
			m.GameId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GameId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
