// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: matchmaker/types.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type Room struct {
	// the config to be used for the game. Either this or the mode need to be set
	// Using this should only be done for custom games
	Config []byte `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	// the game mode to be played
	Mode *Mode `protobuf:"bytes,2,opt,name=mode,proto3" json:"mode,omitempty"`
	// the current players in the room
	Players []string `protobuf:"bytes,3,rep,name=players,proto3" json:"players,omitempty"`
	// pending invitations for players that can join (like a whitelist)
	Pending []string `protobuf:"bytes,4,rep,name=pending,proto3" json:"pending,omitempty"`
	// anyone can join
	Public bool `protobuf:"varint,5,opt,name=public,proto3" json:"public,omitempty"`
	// the minimum amount of players needed to start a game
	Quorum uint32 `protobuf:"varint,6,opt,name=quorum,proto3" json:"quorum,omitempty"`
	// the max amount of players that can join the room
	Capacity uint32 `protobuf:"varint,7,opt,name=capacity,proto3" json:"capacity,omitempty"`
	// when the room was created. Rooms get garbage collected after a while
	Created *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created,proto3" json:"created,omitempty"`
}

func (m *Room) Reset()         { *m = Room{} }
func (m *Room) String() string { return proto.CompactTextString(m) }
func (*Room) ProtoMessage()    {}
func (*Room) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0d5a66ac304af69, []int{0}
}
func (m *Room) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Room) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Room.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Room) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Room.Merge(m, src)
}
func (m *Room) XXX_Size() int {
	return m.Size()
}
func (m *Room) XXX_DiscardUnknown() {
	xxx_messageInfo_Room.DiscardUnknown(m)
}

var xxx_messageInfo_Room proto.InternalMessageInfo

func (m *Room) GetConfig() []byte {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *Room) GetMode() *Mode {
	if m != nil {
		return m.Mode
	}
	return nil
}

func (m *Room) GetPlayers() []string {
	if m != nil {
		return m.Players
	}
	return nil
}

func (m *Room) GetPending() []string {
	if m != nil {
		return m.Pending
	}
	return nil
}

func (m *Room) GetPublic() bool {
	if m != nil {
		return m.Public
	}
	return false
}

func (m *Room) GetQuorum() uint32 {
	if m != nil {
		return m.Quorum
	}
	return 0
}

func (m *Room) GetCapacity() uint32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Room) GetCreated() *timestamppb.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

// Modes are a way of accumulating a small set of possible games that people can choose between
type Mode struct {
	Id     uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Config []byte `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (m *Mode) Reset()         { *m = Mode{} }
func (m *Mode) String() string { return proto.CompactTextString(m) }
func (*Mode) ProtoMessage()    {}
func (*Mode) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0d5a66ac304af69, []int{1}
}
func (m *Mode) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Mode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Mode.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Mode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mode.Merge(m, src)
}
func (m *Mode) XXX_Size() int {
	return m.Size()
}
func (m *Mode) XXX_DiscardUnknown() {
	xxx_messageInfo_Mode.DiscardUnknown(m)
}

var xxx_messageInfo_Mode proto.InternalMessageInfo

func (m *Mode) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Mode) GetConfig() []byte {
	if m != nil {
		return m.Config
	}
	return nil
}

type Params struct {
	// the maximum duration a room can last for before it is closed and all
	// players are kicked
	RoomLifespan *durationpb.Duration `protobuf:"bytes,1,opt,name=room_lifespan,json=roomLifespan,proto3" json:"room_lifespan,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0d5a66ac304af69, []int{2}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetRoomLifespan() *durationpb.Duration {
	if m != nil {
		return m.RoomLifespan
	}
	return nil
}

func init() {
	proto.RegisterType((*Room)(nil), "cmwaters.rook.matchmaker.Room")
	proto.RegisterType((*Mode)(nil), "cmwaters.rook.matchmaker.Mode")
	proto.RegisterType((*Params)(nil), "cmwaters.rook.matchmaker.Params")
}

func init() { proto.RegisterFile("matchmaker/types.proto", fileDescriptor_f0d5a66ac304af69) }

var fileDescriptor_f0d5a66ac304af69 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x31, 0x8f, 0xd3, 0x30,
	0x14, 0xc7, 0xeb, 0x5c, 0x49, 0x8b, 0xef, 0xca, 0xe0, 0xe1, 0x64, 0x32, 0x98, 0xa8, 0x53, 0x24,
	0x24, 0x47, 0x2a, 0xcc, 0x0c, 0x08, 0x24, 0x06, 0x90, 0x90, 0xc5, 0xc4, 0x82, 0x1c, 0xc7, 0xcd,
	0x59, 0x17, 0xe7, 0x05, 0xc7, 0x11, 0xf4, 0x1b, 0x30, 0xf2, 0xb1, 0x18, 0x6f, 0x64, 0x44, 0xed,
	0x17, 0x41, 0x71, 0x12, 0x38, 0x15, 0x31, 0xfe, 0xf3, 0xff, 0x3d, 0xbd, 0x97, 0x5f, 0x82, 0xaf,
	0xad, 0xf4, 0xea, 0xc6, 0xca, 0x5b, 0xed, 0x72, 0x7f, 0x68, 0x75, 0xc7, 0x5b, 0x07, 0x1e, 0x08,
	0x55, 0xf6, 0x8b, 0xf4, 0xda, 0x75, 0xdc, 0x01, 0xdc, 0xf2, 0xbf, 0x54, 0xc2, 0x2a, 0x80, 0xaa,
	0xd6, 0x79, 0xe0, 0x8a, 0x7e, 0x9f, 0x97, 0xbd, 0x93, 0xde, 0x40, 0x33, 0x4e, 0x26, 0x4f, 0xce,
	0x7b, 0x6f, 0xac, 0xee, 0xbc, 0xb4, 0xed, 0x08, 0x6c, 0xbf, 0x45, 0x78, 0x29, 0x00, 0x2c, 0xb9,
	0xc6, 0xb1, 0x82, 0x66, 0x6f, 0x2a, 0x8a, 0x52, 0x94, 0x5d, 0x89, 0x29, 0x91, 0x1d, 0x5e, 0x5a,
	0x28, 0x35, 0x8d, 0x52, 0x94, 0x5d, 0xee, 0x18, 0xff, 0xdf, 0x29, 0xfc, 0x1d, 0x94, 0x5a, 0x04,
	0x96, 0x50, 0xbc, 0x6a, 0x6b, 0x79, 0xd0, 0xae, 0xa3, 0x17, 0xe9, 0x45, 0xf6, 0x50, 0xcc, 0x31,
	0x34, 0xba, 0x29, 0x4d, 0x53, 0xd1, 0xe5, 0xd4, 0x8c, 0x71, 0xd8, 0xdf, 0xf6, 0x45, 0x6d, 0x14,
	0x7d, 0x90, 0xa2, 0x6c, 0x2d, 0xa6, 0x34, 0x3c, 0xff, 0xdc, 0x83, 0xeb, 0x2d, 0x8d, 0x53, 0x94,
	0x6d, 0xc4, 0x94, 0x48, 0x82, 0xd7, 0x4a, 0xb6, 0x52, 0x19, 0x7f, 0xa0, 0xab, 0xd0, 0xfc, 0xc9,
	0xe4, 0x39, 0x5e, 0x29, 0xa7, 0xa5, 0xd7, 0x25, 0x5d, 0x87, 0xb3, 0x13, 0x3e, 0x7a, 0xe0, 0xb3,
	0x07, 0xfe, 0x61, 0xf6, 0x20, 0x66, 0x74, 0xcb, 0xf1, 0x72, 0x78, 0x07, 0xf2, 0x08, 0x47, 0xa6,
	0x0c, 0x16, 0x36, 0x22, 0x32, 0xe5, 0x3d, 0x33, 0xd1, 0x7d, 0x33, 0xdb, 0x37, 0x38, 0x7e, 0x2f,
	0x9d, 0xb4, 0x1d, 0x79, 0x81, 0x37, 0x0e, 0xc0, 0x7e, 0xaa, 0xcd, 0x5e, 0x77, 0xad, 0x6c, 0xc2,
	0xf0, 0xe5, 0xee, 0xf1, 0x3f, 0x5b, 0x5f, 0x4d, 0x5f, 0x47, 0x5c, 0x0d, 0xfc, 0xdb, 0x09, 0x7f,
	0xf9, 0xfa, 0xc7, 0x91, 0xa1, 0xbb, 0x23, 0x43, 0xbf, 0x8e, 0x0c, 0x7d, 0x3f, 0xb1, 0xc5, 0xdd,
	0x89, 0x2d, 0x7e, 0x9e, 0xd8, 0xe2, 0xe3, 0xd3, 0xca, 0xf8, 0x9b, 0xbe, 0xe0, 0x0a, 0x6c, 0x3e,
	0x9b, 0xcf, 0x07, 0xf3, 0xf9, 0xd7, 0xfc, 0xfc, 0x67, 0x29, 0xe2, 0xb0, 0xe7, 0xd9, 0xef, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x7d, 0x0b, 0x46, 0xc5, 0x47, 0x02, 0x00, 0x00,
}

func (m *Room) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Room) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Room) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Created != nil {
		{
			size, err := m.Created.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.Capacity != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Capacity))
		i--
		dAtA[i] = 0x38
	}
	if m.Quorum != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Quorum))
		i--
		dAtA[i] = 0x30
	}
	if m.Public {
		i--
		if m.Public {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.Pending) > 0 {
		for iNdEx := len(m.Pending) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Pending[iNdEx])
			copy(dAtA[i:], m.Pending[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.Pending[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Players) > 0 {
		for iNdEx := len(m.Players) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Players[iNdEx])
			copy(dAtA[i:], m.Players[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.Players[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Mode != nil {
		{
			size, err := m.Mode.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Config) > 0 {
		i -= len(m.Config)
		copy(dAtA[i:], m.Config)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Config)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Mode) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Mode) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Mode) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Config) > 0 {
		i -= len(m.Config)
		copy(dAtA[i:], m.Config)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Config)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RoomLifespan != nil {
		{
			size, err := m.RoomLifespan.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Room) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Config)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Mode != nil {
		l = m.Mode.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Players) > 0 {
		for _, s := range m.Players {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if len(m.Pending) > 0 {
		for _, s := range m.Pending {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.Public {
		n += 2
	}
	if m.Quorum != 0 {
		n += 1 + sovTypes(uint64(m.Quorum))
	}
	if m.Capacity != 0 {
		n += 1 + sovTypes(uint64(m.Capacity))
	}
	if m.Created != nil {
		l = m.Created.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *Mode) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTypes(uint64(m.Id))
	}
	l = len(m.Config)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RoomLifespan != nil {
		l = m.RoomLifespan.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Room) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Room: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Room: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Config = append(m.Config[:0], dAtA[iNdEx:postIndex]...)
			if m.Config == nil {
				m.Config = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mode", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Mode == nil {
				m.Mode = &Mode{}
			}
			if err := m.Mode.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Players", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Players = append(m.Players, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pending", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pending = append(m.Pending, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Public", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Public = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quorum", wireType)
			}
			m.Quorum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Quorum |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Capacity", wireType)
			}
			m.Capacity = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Capacity |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Created", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Created == nil {
				m.Created = &timestamppb.Timestamp{}
			}
			if err := m.Created.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Mode) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Mode: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Mode: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Config = append(m.Config[:0], dAtA[iNdEx:postIndex]...)
			if m.Config == nil {
				m.Config = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoomLifespan", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RoomLifespan == nil {
				m.RoomLifespan = &durationpb.Duration{}
			}
			if err := m.RoomLifespan.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
