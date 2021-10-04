// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rook/claim/claim.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/x/bank/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Action int32

const (
	// Activate action must be completed before all others
	ActionActivate Action = 0
	ActionPlay     Action = 1
	ActionWin      Action = 2
	ActionDelegate Action = 3
	ActionTrade    Action = 4
)

var Action_name = map[int32]string{
	0: "ActionActivate",
	1: "ActionPlay",
	2: "ActionWin",
	3: "ActionDelegate",
	4: "ActionTrade",
}

var Action_value = map[string]int32{
	"ActionActivate": 0,
	"ActionPlay":     1,
	"ActionWin":      2,
	"ActionDelegate": 3,
	"ActionTrade":    4,
}

func (x Action) String() string {
	return proto.EnumName(Action_name, int32(x))
}

func (Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_95c997e8a9ebb362, []int{0}
}

// A Claim Records is the metadata of claim data per address
type ClaimRecord struct {
	// address of claim user
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	// total initial claimable amount for the user
	InitialClaimableAmount types.Coin `protobuf:"bytes,2,opt,name=initial_claimable_amount,json=initialClaimableAmount,proto3" json:"initial_claimable_amount" yaml:"initial_claimable_amount"`
	// true if action is completed
	// index of bool in array refers to action enum #
	ActionCompleted []bool `protobuf:"varint,3,rep,packed,name=action_completed,json=actionCompleted,proto3" json:"action_completed,omitempty" yaml:"action_completed"`
}

func (m *ClaimRecord) Reset()         { *m = ClaimRecord{} }
func (m *ClaimRecord) String() string { return proto.CompactTextString(m) }
func (*ClaimRecord) ProtoMessage()    {}
func (*ClaimRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_95c997e8a9ebb362, []int{0}
}
func (m *ClaimRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimRecord.Merge(m, src)
}
func (m *ClaimRecord) XXX_Size() int {
	return m.Size()
}
func (m *ClaimRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimRecord proto.InternalMessageInfo

func (m *ClaimRecord) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ClaimRecord) GetInitialClaimableAmount() types.Coin {
	if m != nil {
		return m.InitialClaimableAmount
	}
	return types.Coin{}
}

func (m *ClaimRecord) GetActionCompleted() []bool {
	if m != nil {
		return m.ActionCompleted
	}
	return nil
}

// Params defines the claim module's parameters.
type Params struct {
	AirdropStartTime   time.Time     `protobuf:"bytes,1,opt,name=airdrop_start_time,json=airdropStartTime,proto3,stdtime" json:"airdrop_start_time" yaml:"airdrop_start_time"`
	DurationUntilDecay time.Duration `protobuf:"bytes,2,opt,name=duration_until_decay,json=durationUntilDecay,proto3,stdduration" json:"duration_until_decay,omitempty" yaml:"duration_until_decay"`
	DurationOfDecay    time.Duration `protobuf:"bytes,3,opt,name=duration_of_decay,json=durationOfDecay,proto3,stdduration" json:"duration_of_decay,omitempty" yaml:"duration_of_decay"`
	// denom of claimable asset
	ClaimDenom string `protobuf:"bytes,4,opt,name=claim_denom,json=claimDenom,proto3" json:"claim_denom,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_95c997e8a9ebb362, []int{1}
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

func (m *Params) GetAirdropStartTime() time.Time {
	if m != nil {
		return m.AirdropStartTime
	}
	return time.Time{}
}

func (m *Params) GetDurationUntilDecay() time.Duration {
	if m != nil {
		return m.DurationUntilDecay
	}
	return 0
}

func (m *Params) GetDurationOfDecay() time.Duration {
	if m != nil {
		return m.DurationOfDecay
	}
	return 0
}

func (m *Params) GetClaimDenom() string {
	if m != nil {
		return m.ClaimDenom
	}
	return ""
}

func init() {
	proto.RegisterEnum("rook.claim.Action", Action_name, Action_value)
	proto.RegisterType((*ClaimRecord)(nil), "rook.claim.ClaimRecord")
	proto.RegisterType((*Params)(nil), "rook.claim.Params")
}

func init() { proto.RegisterFile("rook/claim/claim.proto", fileDescriptor_95c997e8a9ebb362) }

var fileDescriptor_95c997e8a9ebb362 = []byte{
	// 606 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xc1, 0x6b, 0xd4, 0x4e,
	0x14, 0x4e, 0xba, 0xa5, 0xbf, 0x5f, 0x67, 0xb1, 0x8d, 0x43, 0xa9, 0xe9, 0x16, 0x92, 0x1a, 0x10,
	0x8b, 0xd4, 0x84, 0xd6, 0x9b, 0xb7, 0x66, 0x17, 0x41, 0x2f, 0x96, 0x58, 0x11, 0xbc, 0x84, 0x49,
	0x32, 0x1b, 0x87, 0x66, 0x32, 0x21, 0x33, 0x5b, 0x0c, 0x78, 0x15, 0xc4, 0x53, 0x4f, 0xe2, 0xdd,
	0x7f, 0xc4, 0x63, 0x8f, 0x3d, 0x7a, 0x5a, 0xa5, 0xbd, 0x79, 0xdc, 0xbf, 0x40, 0x26, 0x33, 0x59,
	0x71, 0x5b, 0xf1, 0x32, 0xcc, 0x7c, 0xdf, 0xf7, 0xde, 0xfb, 0xde, 0x9b, 0x61, 0xc0, 0x66, 0xcd,
	0xd8, 0x49, 0x90, 0x16, 0x88, 0x50, 0xb5, 0xfa, 0x55, 0xcd, 0x04, 0x83, 0x40, 0xe2, 0x7e, 0x8b,
	0x0c, 0x36, 0x72, 0x96, 0xb3, 0x16, 0x0e, 0xe4, 0x4e, 0x29, 0x06, 0x4e, 0xca, 0x38, 0x65, 0x3c,
	0x48, 0x10, 0xc7, 0xc1, 0xe9, 0x7e, 0x82, 0x05, 0xda, 0x0f, 0x52, 0x46, 0x4a, 0xcd, 0xdf, 0x9d,
	0xf3, 0xe5, 0xc9, 0x9c, 0xcf, 0x71, 0x89, 0x39, 0xe1, 0x5a, 0xe2, 0xe6, 0x8c, 0xe5, 0x05, 0x0e,
	0xda, 0x53, 0x32, 0x19, 0x07, 0x82, 0x50, 0xcc, 0x05, 0xa2, 0x55, 0x57, 0x63, 0x51, 0x90, 0x4d,
	0x6a, 0x24, 0x08, 0xd3, 0x35, 0xbc, 0xf7, 0x4b, 0xa0, 0x3f, 0x94, 0x1e, 0x23, 0x9c, 0xb2, 0x3a,
	0x83, 0x7b, 0xe0, 0x3f, 0x94, 0x65, 0x35, 0xe6, 0xdc, 0x36, 0x77, 0xcc, 0xdd, 0xd5, 0x10, 0xce,
	0xa6, 0xee, 0x5a, 0x83, 0x68, 0xf1, 0xd8, 0xd3, 0x84, 0x17, 0x75, 0x12, 0xf8, 0x0e, 0xd8, 0xa4,
	0x24, 0x82, 0xa0, 0x22, 0x6e, 0x1b, 0x45, 0x49, 0x81, 0x63, 0x44, 0xd9, 0xa4, 0x14, 0xf6, 0xd2,
	0x8e, 0xb9, 0xdb, 0x3f, 0xd8, 0xf2, 0x55, 0x13, 0xbe, 0x6c, 0xd2, 0xd7, 0x4d, 0xf8, 0x43, 0x46,
	0xca, 0xf0, 0xfe, 0xf9, 0xd4, 0x35, 0x66, 0x53, 0xd7, 0x55, 0xd9, 0xff, 0x96, 0xc8, 0x8b, 0x36,
	0x35, 0x35, 0xec, 0x98, 0xc3, 0x96, 0x80, 0xcf, 0x80, 0x85, 0x52, 0xd9, 0x4b, 0x9c, 0x32, 0x5a,
	0x15, 0x58, 0xe0, 0xcc, 0xee, 0xed, 0xf4, 0x76, 0xff, 0x0f, 0x5d, 0x9d, 0xfa, 0x8e, 0x36, 0xbe,
	0xa0, 0xf2, 0xa2, 0x75, 0x05, 0x0d, 0xe7, 0xc8, 0xd7, 0x1e, 0x58, 0x39, 0x42, 0x35, 0xa2, 0x1c,
	0x32, 0x00, 0x11, 0xa9, 0xb3, 0x9a, 0x55, 0x31, 0x17, 0xa8, 0x16, 0xb1, 0x9c, 0x69, 0x3b, 0x8d,
	0xfe, 0xc1, 0xc0, 0x57, 0xf3, 0xf4, 0xbb, 0x79, 0xfa, 0xc7, 0xdd, 0xc0, 0xc3, 0x7b, 0xba, 0xe8,
	0x96, 0x2e, 0x7a, 0x2d, 0x87, 0x77, 0xf6, 0xdd, 0x35, 0x23, 0x4b, 0x13, 0x2f, 0x24, 0x2e, 0xa3,
	0xe1, 0x27, 0x13, 0x6c, 0x74, 0xd7, 0x12, 0x4f, 0x4a, 0x41, 0x8a, 0x38, 0xc3, 0x29, 0x6a, 0xe6,
	0x23, 0x5c, 0xac, 0x39, 0xd2, 0xe2, 0xf0, 0xa9, 0x2c, 0xf9, 0x73, 0xea, 0x3a, 0x37, 0x85, 0xef,
	0x31, 0x4a, 0x04, 0xa6, 0x95, 0x68, 0x66, 0x53, 0x77, 0x5b, 0x99, 0xba, 0x49, 0xe7, 0x7d, 0x96,
	0xb6, 0x60, 0x47, 0xbd, 0x94, 0xcc, 0x48, 0x12, 0xf0, 0xa3, 0x09, 0x6e, 0xcf, 0x23, 0xd8, 0x58,
	0xbb, 0xea, 0xfd, 0xcb, 0xd5, 0x50, 0xbb, 0xda, 0xbe, 0x16, 0xfb, 0x87, 0x25, 0x7b, 0xc1, 0x52,
	0x27, 0x52, 0x7e, 0xd6, 0x3b, 0xfc, 0xf9, 0x58, 0x99, 0x71, 0x41, 0xbf, 0x7d, 0x1a, 0x71, 0x86,
	0x4b, 0x46, 0xed, 0x65, 0xf9, 0x3a, 0x23, 0xd0, 0x42, 0x23, 0x89, 0x3c, 0x18, 0x83, 0x95, 0xc3,
	0xf6, 0x56, 0x21, 0x04, 0x6b, 0x6a, 0x27, 0xd7, 0x53, 0x24, 0xb0, 0x65, 0xc0, 0x35, 0x00, 0x14,
	0x76, 0x54, 0xa0, 0xc6, 0x32, 0xe1, 0x2d, 0xb0, 0xaa, 0xce, 0xaf, 0x48, 0x69, 0x2d, 0xfd, 0x0e,
	0x19, 0xe1, 0x02, 0xe7, 0x32, 0xa4, 0x07, 0xd7, 0x41, 0x5f, 0x61, 0xc7, 0x35, 0xca, 0xb0, 0xb5,
	0x3c, 0x58, 0xfe, 0xf0, 0xc5, 0x31, 0xc2, 0x27, 0xe7, 0x97, 0x8e, 0x79, 0x71, 0xe9, 0x98, 0x3f,
	0x2e, 0x1d, 0xf3, 0xec, 0xca, 0x31, 0x2e, 0xae, 0x1c, 0xe3, 0xdb, 0x95, 0x63, 0xbc, 0xde, 0xcb,
	0x89, 0x78, 0x33, 0x49, 0xfc, 0x94, 0xd1, 0x00, 0xd5, 0x29, 0x2a, 0xf1, 0x43, 0xde, 0x70, 0x81,
	0x29, 0x0f, 0xda, 0x4f, 0xe2, 0xad, 0xfe, 0x26, 0x44, 0x53, 0x61, 0x9e, 0xac, 0xb4, 0x93, 0x7b,
	0xf4, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xad, 0xab, 0xa9, 0xa5, 0x41, 0x04, 0x00, 0x00,
}

func (m *ClaimRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ActionCompleted) > 0 {
		for iNdEx := len(m.ActionCompleted) - 1; iNdEx >= 0; iNdEx-- {
			i--
			if m.ActionCompleted[iNdEx] {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
		}
		i = encodeVarintClaim(dAtA, i, uint64(len(m.ActionCompleted)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.InitialClaimableAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintClaim(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintClaim(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
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
	if len(m.ClaimDenom) > 0 {
		i -= len(m.ClaimDenom)
		copy(dAtA[i:], m.ClaimDenom)
		i = encodeVarintClaim(dAtA, i, uint64(len(m.ClaimDenom)))
		i--
		dAtA[i] = 0x22
	}
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.DurationOfDecay, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.DurationOfDecay):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintClaim(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	n3, err3 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.DurationUntilDecay, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.DurationUntilDecay):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintClaim(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x12
	n4, err4 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.AirdropStartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.AirdropStartTime):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintClaim(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintClaim(dAtA []byte, offset int, v uint64) int {
	offset -= sovClaim(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClaimRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovClaim(uint64(l))
	}
	l = m.InitialClaimableAmount.Size()
	n += 1 + l + sovClaim(uint64(l))
	if len(m.ActionCompleted) > 0 {
		n += 1 + sovClaim(uint64(len(m.ActionCompleted))) + len(m.ActionCompleted)*1
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.AirdropStartTime)
	n += 1 + l + sovClaim(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.DurationUntilDecay)
	n += 1 + l + sovClaim(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.DurationOfDecay)
	n += 1 + l + sovClaim(uint64(l))
	l = len(m.ClaimDenom)
	if l > 0 {
		n += 1 + l + sovClaim(uint64(l))
	}
	return n
}

func sovClaim(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClaim(x uint64) (n int) {
	return sovClaim(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClaimRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaim
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
			return fmt.Errorf("proto: ClaimRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
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
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialClaimableAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
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
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InitialClaimableAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType == 0 {
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaim
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
				m.ActionCompleted = append(m.ActionCompleted, bool(v != 0))
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaim
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthClaim
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthClaim
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				elementCount = packedLen
				if elementCount != 0 && len(m.ActionCompleted) == 0 {
					m.ActionCompleted = make([]bool, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClaim
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
					m.ActionCompleted = append(m.ActionCompleted, bool(v != 0))
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionCompleted", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClaim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaim
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
				return ErrIntOverflowClaim
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
				return fmt.Errorf("proto: wrong wireType = %d for field AirdropStartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
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
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.AirdropStartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DurationUntilDecay", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
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
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.DurationUntilDecay, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DurationOfDecay", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
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
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.DurationOfDecay, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
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
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClaim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaim
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
func skipClaim(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClaim
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
					return 0, ErrIntOverflowClaim
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
					return 0, ErrIntOverflowClaim
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
				return 0, ErrInvalidLengthClaim
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClaim
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClaim
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClaim        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClaim          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClaim = fmt.Errorf("proto: unexpected end of group")
)
