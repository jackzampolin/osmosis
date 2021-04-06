// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/pool-yield/v1beta1/yield.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/duration"
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

type Params struct {
	// minted_denom is the denomination of the coin expected to be minted by the minting module.
	// Pool-yield module doesn’t actually mint the coin itself, but rather manages the distribution of coins that matches the defined minted_denom.
	MintedDenom string `protobuf:"bytes,1,opt,name=minted_denom,json=mintedDenom,proto3" json:"minted_denom,omitempty" yaml:"minted_denom"`
	// allocation_ratio defines the proportion of the minted minted_denom that is to be allocated as pool incentives.
	AllocationRatio github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=allocation_ratio,json=allocationRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"allocation_ratio" yaml:"allocation_ratio"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4b628395712b609, []int{0}
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

func (m *Params) GetMintedDenom() string {
	if m != nil {
		return m.MintedDenom
	}
	return ""
}

type LockableDurationsInfo struct {
	LockableDurations []time.Duration `protobuf:"bytes,1,rep,name=lockable_durations,json=lockableDurations,proto3,stdduration" json:"lockable_durations" yaml:"lockable_durations"`
}

func (m *LockableDurationsInfo) Reset()         { *m = LockableDurationsInfo{} }
func (m *LockableDurationsInfo) String() string { return proto.CompactTextString(m) }
func (*LockableDurationsInfo) ProtoMessage()    {}
func (*LockableDurationsInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4b628395712b609, []int{1}
}
func (m *LockableDurationsInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LockableDurationsInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LockableDurationsInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LockableDurationsInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockableDurationsInfo.Merge(m, src)
}
func (m *LockableDurationsInfo) XXX_Size() int {
	return m.Size()
}
func (m *LockableDurationsInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LockableDurationsInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LockableDurationsInfo proto.InternalMessageInfo

func (m *LockableDurationsInfo) GetLockableDurations() []time.Duration {
	if m != nil {
		return m.LockableDurations
	}
	return nil
}

type DistrInfo struct {
	TotalWeight github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=total_weight,json=totalWeight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_weight" yaml:"total_weight"`
	Records     []DistrRecord                          `protobuf:"bytes,2,rep,name=records,proto3" json:"records"`
}

func (m *DistrInfo) Reset()         { *m = DistrInfo{} }
func (m *DistrInfo) String() string { return proto.CompactTextString(m) }
func (*DistrInfo) ProtoMessage()    {}
func (*DistrInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4b628395712b609, []int{2}
}
func (m *DistrInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DistrInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DistrInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DistrInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistrInfo.Merge(m, src)
}
func (m *DistrInfo) XXX_Size() int {
	return m.Size()
}
func (m *DistrInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DistrInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DistrInfo proto.InternalMessageInfo

func (m *DistrInfo) GetRecords() []DistrRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

type DistrRecord struct {
	FarmId uint64                                 `protobuf:"varint,1,opt,name=farm_id,json=farmId,proto3" json:"farm_id,omitempty" yaml:"farm_id"`
	Weight github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=weight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"weight"`
}

func (m *DistrRecord) Reset()         { *m = DistrRecord{} }
func (m *DistrRecord) String() string { return proto.CompactTextString(m) }
func (*DistrRecord) ProtoMessage()    {}
func (*DistrRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4b628395712b609, []int{3}
}
func (m *DistrRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DistrRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DistrRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DistrRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistrRecord.Merge(m, src)
}
func (m *DistrRecord) XXX_Size() int {
	return m.Size()
}
func (m *DistrRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_DistrRecord.DiscardUnknown(m)
}

var xxx_messageInfo_DistrRecord proto.InternalMessageInfo

func (m *DistrRecord) GetFarmId() uint64 {
	if m != nil {
		return m.FarmId
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "osmosis.poolyield.v1beta1.Params")
	proto.RegisterType((*LockableDurationsInfo)(nil), "osmosis.poolyield.v1beta1.LockableDurationsInfo")
	proto.RegisterType((*DistrInfo)(nil), "osmosis.poolyield.v1beta1.DistrInfo")
	proto.RegisterType((*DistrRecord)(nil), "osmosis.poolyield.v1beta1.DistrRecord")
}

func init() {
	proto.RegisterFile("osmosis/pool-yield/v1beta1/yield.proto", fileDescriptor_f4b628395712b609)
}

var fileDescriptor_f4b628395712b609 = []byte{
	// 489 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xbf, 0x6b, 0xdb, 0x40,
	0x14, 0xd6, 0xa5, 0xc6, 0x21, 0xe7, 0xd0, 0x1f, 0x6a, 0x4b, 0xec, 0x0c, 0x52, 0x10, 0xd4, 0x04,
	0x8a, 0x4f, 0xa4, 0xdd, 0x3c, 0x0a, 0x37, 0x20, 0xe8, 0x50, 0xb4, 0x14, 0xba, 0x08, 0xfd, 0x38,
	0xcb, 0x22, 0x27, 0x3d, 0xa3, 0x3b, 0xb7, 0xf5, 0x7f, 0xd0, 0xb1, 0x63, 0xc6, 0xfc, 0x23, 0x85,
	0x8e, 0x19, 0x33, 0x96, 0x0e, 0x6a, 0xb1, 0x97, 0xce, 0xfe, 0x0b, 0x8a, 0xee, 0x4e, 0x44, 0xa4,
	0x74, 0xc8, 0xa2, 0xbb, 0xa7, 0xf7, 0xde, 0xf7, 0xbe, 0xef, 0xbb, 0x3b, 0x3c, 0x06, 0x5e, 0x00,
	0xcf, 0xb9, 0xbb, 0x04, 0x60, 0x93, 0x75, 0x4e, 0x59, 0xea, 0x7e, 0x3c, 0x8b, 0xa9, 0x88, 0xce,
	0x5c, 0x19, 0x91, 0x65, 0x05, 0x02, 0xcc, 0x91, 0xae, 0x23, 0x4d, 0x9d, 0x4a, 0xe8, 0xb2, 0xe3,
	0x67, 0x19, 0x64, 0x20, 0xab, 0xdc, 0x66, 0xa7, 0x1a, 0x8e, 0xad, 0x0c, 0x20, 0x63, 0xd4, 0x95,
	0x51, 0xbc, 0x9a, 0xbb, 0xe9, 0xaa, 0x8a, 0x44, 0x0e, 0xa5, 0xca, 0x3b, 0xdf, 0x11, 0xee, 0xbf,
	0x8b, 0xaa, 0xa8, 0xe0, 0xe6, 0x14, 0x1f, 0x16, 0x79, 0x29, 0x68, 0x1a, 0xa6, 0xb4, 0x84, 0x62,
	0x88, 0x4e, 0xd0, 0xe9, 0x81, 0x77, 0xb4, 0xab, 0xed, 0xa7, 0xeb, 0xa8, 0x60, 0x53, 0xa7, 0x9b,
	0x75, 0x82, 0x81, 0x0a, 0x67, 0x4d, 0x64, 0x0a, 0xfc, 0x38, 0x62, 0x0c, 0x12, 0x09, 0x1d, 0xca,
	0x09, 0xc3, 0x3d, 0xd9, 0xef, 0x5f, 0xd7, 0xb6, 0xf1, 0xb3, 0xb6, 0xc7, 0x59, 0x2e, 0x16, 0xab,
	0x98, 0x24, 0x50, 0xb8, 0x89, 0x54, 0xa1, 0x97, 0x09, 0x4f, 0x2f, 0x5c, 0xb1, 0x5e, 0x52, 0x4e,
	0x66, 0x34, 0xd9, 0xd5, 0xf6, 0x91, 0x9a, 0x76, 0x17, 0xcf, 0x09, 0x1e, 0xdd, 0xfe, 0x0a, 0x9a,
	0xef, 0xb4, 0x77, 0x79, 0x65, 0x1b, 0xce, 0x17, 0x84, 0x9f, 0xbf, 0x85, 0xe4, 0x22, 0x8a, 0x19,
	0x9d, 0x69, 0x75, 0xdc, 0x2f, 0xe7, 0x60, 0x02, 0x36, 0x99, 0x4e, 0x84, 0xad, 0x6e, 0x3e, 0x44,
	0x27, 0x0f, 0x4e, 0x07, 0xaf, 0x46, 0x44, 0x39, 0x43, 0x5a, 0x67, 0x48, 0xdb, 0xeb, 0xbd, 0x68,
	0x28, 0xef, 0x6a, 0x7b, 0xa4, 0x88, 0xfc, 0x0b, 0xe1, 0x5c, 0xfe, 0xb2, 0x51, 0xf0, 0x84, 0xdd,
	0x1d, 0xea, 0x7c, 0x43, 0xf8, 0x60, 0x96, 0x73, 0x51, 0xc9, 0xf1, 0x0b, 0x7c, 0x28, 0x40, 0x44,
	0x2c, 0xfc, 0x44, 0xf3, 0x6c, 0x21, 0xb4, 0xa1, 0x6f, 0xee, 0x61, 0x88, 0x5f, 0x8a, 0x5b, 0xfb,
	0xbb, 0x58, 0x4e, 0x30, 0x90, 0xe1, 0x7b, 0x19, 0x99, 0xe7, 0x78, 0xbf, 0xa2, 0x09, 0x54, 0x29,
	0x1f, 0xee, 0x49, 0x75, 0x63, 0xf2, 0xdf, 0x8b, 0x42, 0x24, 0xc1, 0x40, 0x96, 0x7b, 0xbd, 0x86,
	0x4c, 0xd0, 0x36, 0x37, 0x56, 0x0e, 0x3a, 0x69, 0xf3, 0x25, 0xde, 0x9f, 0x47, 0x55, 0x11, 0xe6,
	0xa9, 0x24, 0xdf, 0xf3, 0xcc, 0x5d, 0x6d, 0x3f, 0x54, 0x74, 0x74, 0xc2, 0x09, 0xfa, 0xcd, 0xce,
	0x4f, 0xcd, 0x73, 0xdc, 0xd7, 0x42, 0xd5, 0xc9, 0x93, 0xfb, 0x09, 0x0d, 0x74, 0xf7, 0xb4, 0xf7,
	0xe7, 0xca, 0x46, 0x9e, 0x7f, 0xbd, 0xb1, 0xd0, 0xcd, 0xc6, 0x42, 0xbf, 0x37, 0x16, 0xfa, 0xba,
	0xb5, 0x8c, 0x9b, 0xad, 0x65, 0xfc, 0xd8, 0x5a, 0xc6, 0x07, 0xb7, 0x8b, 0x37, 0x69, 0x1f, 0x4e,
	0xbb, 0x7e, 0xee, 0x3e, 0x21, 0x09, 0x1e, 0xf7, 0xe5, 0x11, 0xbf, 0xfe, 0x1b, 0x00, 0x00, 0xff,
	0xff, 0xb9, 0x7c, 0x8c, 0xf5, 0x65, 0x03, 0x00, 0x00,
}

func (this *DistrRecord) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DistrRecord)
	if !ok {
		that2, ok := that.(DistrRecord)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.FarmId != that1.FarmId {
		return false
	}
	if !this.Weight.Equal(that1.Weight) {
		return false
	}
	return true
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
	{
		size := m.AllocationRatio.Size()
		i -= size
		if _, err := m.AllocationRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintYield(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.MintedDenom) > 0 {
		i -= len(m.MintedDenom)
		copy(dAtA[i:], m.MintedDenom)
		i = encodeVarintYield(dAtA, i, uint64(len(m.MintedDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LockableDurationsInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LockableDurationsInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LockableDurationsInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LockableDurations) > 0 {
		for iNdEx := len(m.LockableDurations) - 1; iNdEx >= 0; iNdEx-- {
			n, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.LockableDurations[iNdEx], dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.LockableDurations[iNdEx]):])
			if err != nil {
				return 0, err
			}
			i -= n
			i = encodeVarintYield(dAtA, i, uint64(n))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *DistrInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DistrInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DistrInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Records) > 0 {
		for iNdEx := len(m.Records) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Records[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintYield(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size := m.TotalWeight.Size()
		i -= size
		if _, err := m.TotalWeight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintYield(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *DistrRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DistrRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DistrRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Weight.Size()
		i -= size
		if _, err := m.Weight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintYield(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.FarmId != 0 {
		i = encodeVarintYield(dAtA, i, uint64(m.FarmId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintYield(dAtA []byte, offset int, v uint64) int {
	offset -= sovYield(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MintedDenom)
	if l > 0 {
		n += 1 + l + sovYield(uint64(l))
	}
	l = m.AllocationRatio.Size()
	n += 1 + l + sovYield(uint64(l))
	return n
}

func (m *LockableDurationsInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.LockableDurations) > 0 {
		for _, e := range m.LockableDurations {
			l = github_com_gogo_protobuf_types.SizeOfStdDuration(e)
			n += 1 + l + sovYield(uint64(l))
		}
	}
	return n
}

func (m *DistrInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TotalWeight.Size()
	n += 1 + l + sovYield(uint64(l))
	if len(m.Records) > 0 {
		for _, e := range m.Records {
			l = e.Size()
			n += 1 + l + sovYield(uint64(l))
		}
	}
	return n
}

func (m *DistrRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FarmId != 0 {
		n += 1 + sovYield(uint64(m.FarmId))
	}
	l = m.Weight.Size()
	n += 1 + l + sovYield(uint64(l))
	return n
}

func sovYield(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozYield(x uint64) (n int) {
	return sovYield(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowYield
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
				return fmt.Errorf("proto: wrong wireType = %d for field MintedDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowYield
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
				return ErrInvalidLengthYield
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthYield
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintedDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllocationRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowYield
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
				return ErrInvalidLengthYield
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthYield
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AllocationRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipYield(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthYield
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthYield
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
func (m *LockableDurationsInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowYield
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
			return fmt.Errorf("proto: LockableDurationsInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LockableDurationsInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockableDurations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowYield
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
				return ErrInvalidLengthYield
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthYield
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LockableDurations = append(m.LockableDurations, time.Duration(0))
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&(m.LockableDurations[len(m.LockableDurations)-1]), dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipYield(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthYield
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthYield
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
func (m *DistrInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowYield
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
			return fmt.Errorf("proto: DistrInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DistrInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalWeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowYield
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
				return ErrInvalidLengthYield
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthYield
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalWeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Records", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowYield
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
				return ErrInvalidLengthYield
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthYield
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Records = append(m.Records, DistrRecord{})
			if err := m.Records[len(m.Records)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipYield(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthYield
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthYield
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
func (m *DistrRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowYield
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
			return fmt.Errorf("proto: DistrRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DistrRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FarmId", wireType)
			}
			m.FarmId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowYield
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FarmId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowYield
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
				return ErrInvalidLengthYield
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthYield
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Weight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipYield(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthYield
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthYield
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
func skipYield(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowYield
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
					return 0, ErrIntOverflowYield
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
					return 0, ErrIntOverflowYield
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
				return 0, ErrInvalidLengthYield
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupYield
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthYield
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthYield        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowYield          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupYield = fmt.Errorf("proto: unexpected end of group")
)
