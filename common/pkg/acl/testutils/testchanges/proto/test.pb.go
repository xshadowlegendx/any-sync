// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/acl/testutils/testchanges/proto/test.proto

package testchanges

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

type PlainTextChange struct {
}

func (m *PlainTextChange) Reset()         { *m = PlainTextChange{} }
func (m *PlainTextChange) String() string { return proto.CompactTextString(m) }
func (*PlainTextChange) ProtoMessage()    {}
func (*PlainTextChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_37f33c266ada4318, []int{0}
}
func (m *PlainTextChange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PlainTextChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PlainTextChange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PlainTextChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlainTextChange.Merge(m, src)
}
func (m *PlainTextChange) XXX_Size() int {
	return m.Size()
}
func (m *PlainTextChange) XXX_DiscardUnknown() {
	xxx_messageInfo_PlainTextChange.DiscardUnknown(m)
}

var xxx_messageInfo_PlainTextChange proto.InternalMessageInfo

type PlainTextChange_Content struct {
	// Types that are valid to be assigned to Value:
	//	*PlainTextChange_Content_TextAppend
	Value isPlainTextChange_Content_Value `protobuf_oneof:"value"`
}

func (m *PlainTextChange_Content) Reset()         { *m = PlainTextChange_Content{} }
func (m *PlainTextChange_Content) String() string { return proto.CompactTextString(m) }
func (*PlainTextChange_Content) ProtoMessage()    {}
func (*PlainTextChange_Content) Descriptor() ([]byte, []int) {
	return fileDescriptor_37f33c266ada4318, []int{0, 0}
}
func (m *PlainTextChange_Content) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PlainTextChange_Content) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PlainTextChange_Content.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PlainTextChange_Content) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlainTextChange_Content.Merge(m, src)
}
func (m *PlainTextChange_Content) XXX_Size() int {
	return m.Size()
}
func (m *PlainTextChange_Content) XXX_DiscardUnknown() {
	xxx_messageInfo_PlainTextChange_Content.DiscardUnknown(m)
}

var xxx_messageInfo_PlainTextChange_Content proto.InternalMessageInfo

type isPlainTextChange_Content_Value interface {
	isPlainTextChange_Content_Value()
	MarshalTo([]byte) (int, error)
	Size() int
}

type PlainTextChange_Content_TextAppend struct {
	TextAppend *PlainTextChange_TextAppend `protobuf:"bytes,1,opt,name=textAppend,proto3,oneof" json:"textAppend,omitempty"`
}

func (*PlainTextChange_Content_TextAppend) isPlainTextChange_Content_Value() {}

func (m *PlainTextChange_Content) GetValue() isPlainTextChange_Content_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *PlainTextChange_Content) GetTextAppend() *PlainTextChange_TextAppend {
	if x, ok := m.GetValue().(*PlainTextChange_Content_TextAppend); ok {
		return x.TextAppend
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PlainTextChange_Content) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PlainTextChange_Content_TextAppend)(nil),
	}
}

type PlainTextChange_TextAppend struct {
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (m *PlainTextChange_TextAppend) Reset()         { *m = PlainTextChange_TextAppend{} }
func (m *PlainTextChange_TextAppend) String() string { return proto.CompactTextString(m) }
func (*PlainTextChange_TextAppend) ProtoMessage()    {}
func (*PlainTextChange_TextAppend) Descriptor() ([]byte, []int) {
	return fileDescriptor_37f33c266ada4318, []int{0, 1}
}
func (m *PlainTextChange_TextAppend) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PlainTextChange_TextAppend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PlainTextChange_TextAppend.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PlainTextChange_TextAppend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlainTextChange_TextAppend.Merge(m, src)
}
func (m *PlainTextChange_TextAppend) XXX_Size() int {
	return m.Size()
}
func (m *PlainTextChange_TextAppend) XXX_DiscardUnknown() {
	xxx_messageInfo_PlainTextChange_TextAppend.DiscardUnknown(m)
}

var xxx_messageInfo_PlainTextChange_TextAppend proto.InternalMessageInfo

func (m *PlainTextChange_TextAppend) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type PlainTextChange_Snapshot struct {
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (m *PlainTextChange_Snapshot) Reset()         { *m = PlainTextChange_Snapshot{} }
func (m *PlainTextChange_Snapshot) String() string { return proto.CompactTextString(m) }
func (*PlainTextChange_Snapshot) ProtoMessage()    {}
func (*PlainTextChange_Snapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_37f33c266ada4318, []int{0, 2}
}
func (m *PlainTextChange_Snapshot) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PlainTextChange_Snapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PlainTextChange_Snapshot.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PlainTextChange_Snapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlainTextChange_Snapshot.Merge(m, src)
}
func (m *PlainTextChange_Snapshot) XXX_Size() int {
	return m.Size()
}
func (m *PlainTextChange_Snapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_PlainTextChange_Snapshot.DiscardUnknown(m)
}

var xxx_messageInfo_PlainTextChange_Snapshot proto.InternalMessageInfo

func (m *PlainTextChange_Snapshot) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type PlainTextChange_Data struct {
	Content  []*PlainTextChange_Content `protobuf:"bytes,1,rep,name=content,proto3" json:"content,omitempty"`
	Snapshot *PlainTextChange_Snapshot  `protobuf:"bytes,2,opt,name=snapshot,proto3" json:"snapshot,omitempty"`
}

func (m *PlainTextChange_Data) Reset()         { *m = PlainTextChange_Data{} }
func (m *PlainTextChange_Data) String() string { return proto.CompactTextString(m) }
func (*PlainTextChange_Data) ProtoMessage()    {}
func (*PlainTextChange_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_37f33c266ada4318, []int{0, 3}
}
func (m *PlainTextChange_Data) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PlainTextChange_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PlainTextChange_Data.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PlainTextChange_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlainTextChange_Data.Merge(m, src)
}
func (m *PlainTextChange_Data) XXX_Size() int {
	return m.Size()
}
func (m *PlainTextChange_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_PlainTextChange_Data.DiscardUnknown(m)
}

var xxx_messageInfo_PlainTextChange_Data proto.InternalMessageInfo

func (m *PlainTextChange_Data) GetContent() []*PlainTextChange_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *PlainTextChange_Data) GetSnapshot() *PlainTextChange_Snapshot {
	if m != nil {
		return m.Snapshot
	}
	return nil
}

func init() {
	proto.RegisterType((*PlainTextChange)(nil), "anytype.PlainTextChange")
	proto.RegisterType((*PlainTextChange_Content)(nil), "anytype.PlainTextChange.Content")
	proto.RegisterType((*PlainTextChange_TextAppend)(nil), "anytype.PlainTextChange.TextAppend")
	proto.RegisterType((*PlainTextChange_Snapshot)(nil), "anytype.PlainTextChange.Snapshot")
	proto.RegisterType((*PlainTextChange_Data)(nil), "anytype.PlainTextChange.Data")
}

func init() {
	proto.RegisterFile("pkg/acl/testutils/testchanges/proto/test.proto", fileDescriptor_37f33c266ada4318)
}

var fileDescriptor_37f33c266ada4318 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2b, 0xc8, 0x4e, 0xd7,
	0x4f, 0x4c, 0xce, 0xd1, 0x2f, 0x49, 0x2d, 0x2e, 0x29, 0x2d, 0xc9, 0xcc, 0x29, 0x06, 0xb3, 0x92,
	0x33, 0x12, 0xf3, 0xd2, 0x53, 0x8b, 0xf5, 0x0b, 0x8a, 0xf2, 0x4b, 0xf2, 0xc1, 0x22, 0x7a, 0x60,
	0xa6, 0x10, 0x7b, 0x62, 0x5e, 0x65, 0x49, 0x65, 0x41, 0xaa, 0xd2, 0x26, 0x26, 0x2e, 0xfe, 0x80,
	0x9c, 0xc4, 0xcc, 0xbc, 0x90, 0xd4, 0x8a, 0x12, 0x67, 0xb0, 0x72, 0xa9, 0x48, 0x2e, 0x76, 0xe7,
	0xfc, 0xbc, 0x92, 0xd4, 0xbc, 0x12, 0x21, 0x57, 0x2e, 0xae, 0x92, 0xd4, 0x8a, 0x12, 0xc7, 0x82,
	0x82, 0xd4, 0xbc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x65, 0x3d, 0xa8, 0x66, 0x3d,
	0x34, 0x8d, 0x7a, 0x21, 0x70, 0xa5, 0x1e, 0x0c, 0x41, 0x48, 0x1a, 0x9d, 0xd8, 0xb9, 0x58, 0xcb,
	0x12, 0x73, 0x4a, 0x53, 0xa5, 0x14, 0xb8, 0xb8, 0x10, 0x8a, 0x84, 0x84, 0xb8, 0x58, 0x40, 0x8a,
	0xc0, 0xe6, 0x72, 0x06, 0x81, 0xd9, 0x52, 0x72, 0x5c, 0x1c, 0xc1, 0x79, 0x89, 0x05, 0xc5, 0x19,
	0xf9, 0x25, 0x58, 0xe5, 0x1b, 0x19, 0xb9, 0x58, 0x5c, 0x12, 0x4b, 0x12, 0x85, 0xac, 0xb8, 0xd8,
	0x93, 0x21, 0xae, 0x94, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0x52, 0xc0, 0xe9, 0x2e, 0xa8, 0x6f,
	0x82, 0x60, 0x1a, 0x84, 0x6c, 0xb9, 0x38, 0x8a, 0xa1, 0x96, 0x48, 0x30, 0x81, 0x3d, 0xa5, 0x88,
	0x53, 0x33, 0xcc, 0x35, 0x41, 0x70, 0x2d, 0x4e, 0xaa, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24,
	0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78,
	0x2c, 0xc7, 0x10, 0xc5, 0x8d, 0x14, 0xea, 0x49, 0x6c, 0xe0, 0xb0, 0x36, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xf8, 0x8c, 0x6a, 0x1d, 0x9d, 0x01, 0x00, 0x00,
}

func (m *PlainTextChange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PlainTextChange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PlainTextChange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *PlainTextChange_Content) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PlainTextChange_Content) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PlainTextChange_Content) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != nil {
		{
			size := m.Value.Size()
			i -= size
			if _, err := m.Value.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *PlainTextChange_Content_TextAppend) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PlainTextChange_Content_TextAppend) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.TextAppend != nil {
		{
			size, err := m.TextAppend.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTest(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *PlainTextChange_TextAppend) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PlainTextChange_TextAppend) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PlainTextChange_TextAppend) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Text) > 0 {
		i -= len(m.Text)
		copy(dAtA[i:], m.Text)
		i = encodeVarintTest(dAtA, i, uint64(len(m.Text)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PlainTextChange_Snapshot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PlainTextChange_Snapshot) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PlainTextChange_Snapshot) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Text) > 0 {
		i -= len(m.Text)
		copy(dAtA[i:], m.Text)
		i = encodeVarintTest(dAtA, i, uint64(len(m.Text)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PlainTextChange_Data) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PlainTextChange_Data) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PlainTextChange_Data) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Snapshot != nil {
		{
			size, err := m.Snapshot.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTest(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Content) > 0 {
		for iNdEx := len(m.Content) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Content[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTest(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintTest(dAtA []byte, offset int, v uint64) int {
	offset -= sovTest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PlainTextChange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *PlainTextChange_Content) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != nil {
		n += m.Value.Size()
	}
	return n
}

func (m *PlainTextChange_Content_TextAppend) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TextAppend != nil {
		l = m.TextAppend.Size()
		n += 1 + l + sovTest(uint64(l))
	}
	return n
}
func (m *PlainTextChange_TextAppend) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Text)
	if l > 0 {
		n += 1 + l + sovTest(uint64(l))
	}
	return n
}

func (m *PlainTextChange_Snapshot) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Text)
	if l > 0 {
		n += 1 + l + sovTest(uint64(l))
	}
	return n
}

func (m *PlainTextChange_Data) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Content) > 0 {
		for _, e := range m.Content {
			l = e.Size()
			n += 1 + l + sovTest(uint64(l))
		}
	}
	if m.Snapshot != nil {
		l = m.Snapshot.Size()
		n += 1 + l + sovTest(uint64(l))
	}
	return n
}

func sovTest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTest(x uint64) (n int) {
	return sovTest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PlainTextChange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTest
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
			return fmt.Errorf("proto: PlainTextChange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PlainTextChange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTest
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
func (m *PlainTextChange_Content) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTest
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
			return fmt.Errorf("proto: Content: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Content: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TextAppend", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
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
				return ErrInvalidLengthTest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &PlainTextChange_TextAppend{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &PlainTextChange_Content_TextAppend{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTest
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
func (m *PlainTextChange_TextAppend) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTest
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
			return fmt.Errorf("proto: TextAppend: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TextAppend: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Text", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
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
				return ErrInvalidLengthTest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Text = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTest
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
func (m *PlainTextChange_Snapshot) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTest
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
			return fmt.Errorf("proto: Snapshot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Snapshot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Text", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
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
				return ErrInvalidLengthTest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Text = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTest
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
func (m *PlainTextChange_Data) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTest
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
			return fmt.Errorf("proto: Data: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Data: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
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
				return ErrInvalidLengthTest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = append(m.Content, &PlainTextChange_Content{})
			if err := m.Content[len(m.Content)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Snapshot", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
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
				return ErrInvalidLengthTest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Snapshot == nil {
				m.Snapshot = &PlainTextChange_Snapshot{}
			}
			if err := m.Snapshot.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTest
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
func skipTest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTest
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
					return 0, ErrIntOverflowTest
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
					return 0, ErrIntOverflowTest
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
				return 0, ErrInvalidLengthTest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTest = fmt.Errorf("proto: unexpected end of group")
)