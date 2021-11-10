// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rook/matchmaker/event.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type EventRoomUpdated struct {
	RoomId         uint64   `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	AddedPlayers   []string `protobuf:"bytes,2,rep,name=added_players,json=addedPlayers,proto3" json:"added_players,omitempty"`
	RemovedPlayers []string `protobuf:"bytes,3,rep,name=removed_players,json=removedPlayers,proto3" json:"removed_players,omitempty"`
}

func (m *EventRoomUpdated) Reset()         { *m = EventRoomUpdated{} }
func (m *EventRoomUpdated) String() string { return proto.CompactTextString(m) }
func (*EventRoomUpdated) ProtoMessage()    {}
func (*EventRoomUpdated) Descriptor() ([]byte, []int) {
	return fileDescriptor_94273982a1588e6b, []int{0}
}
func (m *EventRoomUpdated) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRoomUpdated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventRoomUpdated.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventRoomUpdated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRoomUpdated.Merge(m, src)
}
func (m *EventRoomUpdated) XXX_Size() int {
	return m.Size()
}
func (m *EventRoomUpdated) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRoomUpdated.DiscardUnknown(m)
}

var xxx_messageInfo_EventRoomUpdated proto.InternalMessageInfo

func (m *EventRoomUpdated) GetRoomId() uint64 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *EventRoomUpdated) GetAddedPlayers() []string {
	if m != nil {
		return m.AddedPlayers
	}
	return nil
}

func (m *EventRoomUpdated) GetRemovedPlayers() []string {
	if m != nil {
		return m.RemovedPlayers
	}
	return nil
}

type EventRoomError struct {
	RoomId uint64 `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *EventRoomError) Reset()         { *m = EventRoomError{} }
func (m *EventRoomError) String() string { return proto.CompactTextString(m) }
func (*EventRoomError) ProtoMessage()    {}
func (*EventRoomError) Descriptor() ([]byte, []int) {
	return fileDescriptor_94273982a1588e6b, []int{1}
}
func (m *EventRoomError) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRoomError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventRoomError.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventRoomError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRoomError.Merge(m, src)
}
func (m *EventRoomError) XXX_Size() int {
	return m.Size()
}
func (m *EventRoomError) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRoomError.DiscardUnknown(m)
}

var xxx_messageInfo_EventRoomError proto.InternalMessageInfo

func (m *EventRoomError) GetRoomId() uint64 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *EventRoomError) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type EventGameStarted struct {
	RoomId uint64 `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	GameId uint64 `protobuf:"varint,2,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
}

func (m *EventGameStarted) Reset()         { *m = EventGameStarted{} }
func (m *EventGameStarted) String() string { return proto.CompactTextString(m) }
func (*EventGameStarted) ProtoMessage()    {}
func (*EventGameStarted) Descriptor() ([]byte, []int) {
	return fileDescriptor_94273982a1588e6b, []int{2}
}
func (m *EventGameStarted) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventGameStarted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventGameStarted.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventGameStarted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventGameStarted.Merge(m, src)
}
func (m *EventGameStarted) XXX_Size() int {
	return m.Size()
}
func (m *EventGameStarted) XXX_DiscardUnknown() {
	xxx_messageInfo_EventGameStarted.DiscardUnknown(m)
}

var xxx_messageInfo_EventGameStarted proto.InternalMessageInfo

func (m *EventGameStarted) GetRoomId() uint64 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *EventGameStarted) GetGameId() uint64 {
	if m != nil {
		return m.GameId
	}
	return 0
}

func init() {
	proto.RegisterType((*EventRoomUpdated)(nil), "rook.matchmaker.EventRoomUpdated")
	proto.RegisterType((*EventRoomError)(nil), "rook.matchmaker.EventRoomError")
	proto.RegisterType((*EventGameStarted)(nil), "rook.matchmaker.EventGameStarted")
}

func init() { proto.RegisterFile("rook/matchmaker/event.proto", fileDescriptor_94273982a1588e6b) }

var fileDescriptor_94273982a1588e6b = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0x3b, 0xe9, 0xff, 0xb7, 0x74, 0xd0, 0x56, 0x82, 0xd0, 0x80, 0x30, 0x84, 0xb8, 0x30,
	0x1b, 0x13, 0xc4, 0x07, 0x10, 0xc4, 0x22, 0xc5, 0x8d, 0x44, 0xdc, 0xb8, 0x29, 0xd3, 0xcc, 0xa5,
	0x2d, 0x75, 0x7a, 0xc3, 0xcd, 0x58, 0xcc, 0x5b, 0xf8, 0x58, 0x2e, 0xbb, 0x74, 0x29, 0xc9, 0x8b,
	0xc8, 0x24, 0x41, 0xbb, 0xd1, 0xe5, 0x39, 0xdf, 0x07, 0xf7, 0x72, 0xf8, 0x09, 0x21, 0xae, 0x63,
	0x2d, 0x4d, 0xba, 0xd4, 0x72, 0x0d, 0x14, 0xc3, 0x16, 0x36, 0x26, 0xca, 0x08, 0x0d, 0xba, 0x23,
	0x0b, 0xa3, 0x1f, 0x18, 0x14, 0xfc, 0x68, 0x62, 0x79, 0x82, 0xa8, 0x1f, 0x33, 0x25, 0x0d, 0x28,
	0x77, 0xcc, 0xfb, 0x84, 0xa8, 0x67, 0x2b, 0xe5, 0x31, 0x9f, 0x85, 0xff, 0x92, 0x9e, 0x8d, 0x53,
	0xe5, 0x9e, 0xf2, 0x43, 0xa9, 0x14, 0xa8, 0x59, 0xf6, 0x2c, 0x0b, 0xa0, 0xdc, 0x73, 0xfc, 0x6e,
	0x38, 0x48, 0x0e, 0xea, 0xf2, 0xbe, 0xe9, 0xdc, 0x33, 0x3e, 0x22, 0xd0, 0xb8, 0xdd, 0xd3, 0xba,
	0xb5, 0x36, 0x6c, 0xeb, 0x56, 0x0c, 0xae, 0xf8, 0xf0, 0xfb, 0xf4, 0x84, 0x08, 0xe9, 0xf7, 0xc3,
	0xc7, 0xfc, 0x3f, 0x58, 0xc3, 0x73, 0x7c, 0x16, 0x0e, 0x92, 0x26, 0x04, 0x37, 0xed, 0xef, 0xb7,
	0x52, 0xc3, 0x83, 0x91, 0xf4, 0xe7, 0xef, 0x63, 0xde, 0x5f, 0x48, 0x0d, 0x16, 0x38, 0x0d, 0xb0,
	0x71, 0xaa, 0xae, 0xef, 0xde, 0x4b, 0xc1, 0x76, 0xa5, 0x60, 0x9f, 0xa5, 0x60, 0x6f, 0x95, 0xe8,
	0xec, 0x2a, 0xd1, 0xf9, 0xa8, 0x44, 0xe7, 0xe9, 0x62, 0xb1, 0x32, 0xcb, 0x97, 0x79, 0x94, 0xa2,
	0x8e, 0x25, 0xa5, 0x72, 0x03, 0xe7, 0x79, 0x91, 0x1b, 0xd0, 0x79, 0x5c, 0x6f, 0xfc, 0xba, 0xbf,
	0xb2, 0x29, 0x32, 0xc8, 0xe7, 0xbd, 0x7a, 0xe6, 0xcb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4b,
	0x80, 0x1e, 0x5e, 0x85, 0x01, 0x00, 0x00,
}

func (m *EventRoomUpdated) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRoomUpdated) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRoomUpdated) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RemovedPlayers) > 0 {
		for iNdEx := len(m.RemovedPlayers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.RemovedPlayers[iNdEx])
			copy(dAtA[i:], m.RemovedPlayers[iNdEx])
			i = encodeVarintEvent(dAtA, i, uint64(len(m.RemovedPlayers[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.AddedPlayers) > 0 {
		for iNdEx := len(m.AddedPlayers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AddedPlayers[iNdEx])
			copy(dAtA[i:], m.AddedPlayers[iNdEx])
			i = encodeVarintEvent(dAtA, i, uint64(len(m.AddedPlayers[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.RoomId != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.RoomId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventRoomError) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRoomError) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRoomError) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0x12
	}
	if m.RoomId != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.RoomId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventGameStarted) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventGameStarted) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventGameStarted) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GameId != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.GameId))
		i--
		dAtA[i] = 0x10
	}
	if m.RoomId != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.RoomId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventRoomUpdated) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RoomId != 0 {
		n += 1 + sovEvent(uint64(m.RoomId))
	}
	if len(m.AddedPlayers) > 0 {
		for _, s := range m.AddedPlayers {
			l = len(s)
			n += 1 + l + sovEvent(uint64(l))
		}
	}
	if len(m.RemovedPlayers) > 0 {
		for _, s := range m.RemovedPlayers {
			l = len(s)
			n += 1 + l + sovEvent(uint64(l))
		}
	}
	return n
}

func (m *EventRoomError) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RoomId != 0 {
		n += 1 + sovEvent(uint64(m.RoomId))
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *EventGameStarted) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RoomId != 0 {
		n += 1 + sovEvent(uint64(m.RoomId))
	}
	if m.GameId != 0 {
		n += 1 + sovEvent(uint64(m.GameId))
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventRoomUpdated) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventRoomUpdated: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRoomUpdated: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoomId", wireType)
			}
			m.RoomId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RoomId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddedPlayers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AddedPlayers = append(m.AddedPlayers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemovedPlayers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RemovedPlayers = append(m.RemovedPlayers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventRoomError) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventRoomError: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRoomError: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoomId", wireType)
			}
			m.RoomId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RoomId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventGameStarted) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventGameStarted: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventGameStarted: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoomId", wireType)
			}
			m.RoomId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RoomId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GameId", wireType)
			}
			m.GameId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)
