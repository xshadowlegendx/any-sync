// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: coordinator/coordinatorproto/protos/coordinator.proto

package coordinatorproto

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

type SpaceSignRequest struct {
	SpaceId string `protobuf:"bytes,1,opt,name=spaceId,proto3" json:"spaceId,omitempty"`
}

func (m *SpaceSignRequest) Reset()         { *m = SpaceSignRequest{} }
func (m *SpaceSignRequest) String() string { return proto.CompactTextString(m) }
func (*SpaceSignRequest) ProtoMessage()    {}
func (*SpaceSignRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d94f6f99586adae2, []int{0}
}
func (m *SpaceSignRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SpaceSignRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SpaceSignRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SpaceSignRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpaceSignRequest.Merge(m, src)
}
func (m *SpaceSignRequest) XXX_Size() int {
	return m.Size()
}
func (m *SpaceSignRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SpaceSignRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SpaceSignRequest proto.InternalMessageInfo

func (m *SpaceSignRequest) GetSpaceId() string {
	if m != nil {
		return m.SpaceId
	}
	return ""
}

type SpaceSignResponse struct {
	Receipt *SpaceReceiptWithSignature `protobuf:"bytes,1,opt,name=receipt,proto3" json:"receipt,omitempty"`
}

func (m *SpaceSignResponse) Reset()         { *m = SpaceSignResponse{} }
func (m *SpaceSignResponse) String() string { return proto.CompactTextString(m) }
func (*SpaceSignResponse) ProtoMessage()    {}
func (*SpaceSignResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d94f6f99586adae2, []int{1}
}
func (m *SpaceSignResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SpaceSignResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SpaceSignResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SpaceSignResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpaceSignResponse.Merge(m, src)
}
func (m *SpaceSignResponse) XXX_Size() int {
	return m.Size()
}
func (m *SpaceSignResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SpaceSignResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SpaceSignResponse proto.InternalMessageInfo

func (m *SpaceSignResponse) GetReceipt() *SpaceReceiptWithSignature {
	if m != nil {
		return m.Receipt
	}
	return nil
}

// SpaceReceiptWithSignature contains protobuf encoded receipt and its signature
type SpaceReceiptWithSignature struct {
	SpaceReceiptPayload []byte `protobuf:"bytes,1,opt,name=spaceReceiptPayload,proto3" json:"spaceReceiptPayload,omitempty"`
	Signature           []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *SpaceReceiptWithSignature) Reset()         { *m = SpaceReceiptWithSignature{} }
func (m *SpaceReceiptWithSignature) String() string { return proto.CompactTextString(m) }
func (*SpaceReceiptWithSignature) ProtoMessage()    {}
func (*SpaceReceiptWithSignature) Descriptor() ([]byte, []int) {
	return fileDescriptor_d94f6f99586adae2, []int{2}
}
func (m *SpaceReceiptWithSignature) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SpaceReceiptWithSignature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SpaceReceiptWithSignature.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SpaceReceiptWithSignature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpaceReceiptWithSignature.Merge(m, src)
}
func (m *SpaceReceiptWithSignature) XXX_Size() int {
	return m.Size()
}
func (m *SpaceReceiptWithSignature) XXX_DiscardUnknown() {
	xxx_messageInfo_SpaceReceiptWithSignature.DiscardUnknown(m)
}

var xxx_messageInfo_SpaceReceiptWithSignature proto.InternalMessageInfo

func (m *SpaceReceiptWithSignature) GetSpaceReceiptPayload() []byte {
	if m != nil {
		return m.SpaceReceiptPayload
	}
	return nil
}

func (m *SpaceReceiptWithSignature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// SpaceReceipt contains permission to SpacePush operation
type SpaceReceipt struct {
	SpaceId string `protobuf:"bytes,1,opt,name=spaceId,proto3" json:"spaceId,omitempty"`
	// identity of space owner
	AccountIdentity []byte `protobuf:"bytes,2,opt,name=accountIdentity,proto3" json:"accountIdentity,omitempty"`
	// identity of control node
	ControlNodeIdentity []byte `protobuf:"bytes,3,opt,name=controlNodeIdentity,proto3" json:"controlNodeIdentity,omitempty"`
	// unix-timestamp with a deadline time of receipt validity
	ValidUntil uint64 `protobuf:"varint,4,opt,name=validUntil,proto3" json:"validUntil,omitempty"`
}

func (m *SpaceReceipt) Reset()         { *m = SpaceReceipt{} }
func (m *SpaceReceipt) String() string { return proto.CompactTextString(m) }
func (*SpaceReceipt) ProtoMessage()    {}
func (*SpaceReceipt) Descriptor() ([]byte, []int) {
	return fileDescriptor_d94f6f99586adae2, []int{3}
}
func (m *SpaceReceipt) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SpaceReceipt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SpaceReceipt.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SpaceReceipt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpaceReceipt.Merge(m, src)
}
func (m *SpaceReceipt) XXX_Size() int {
	return m.Size()
}
func (m *SpaceReceipt) XXX_DiscardUnknown() {
	xxx_messageInfo_SpaceReceipt.DiscardUnknown(m)
}

var xxx_messageInfo_SpaceReceipt proto.InternalMessageInfo

func (m *SpaceReceipt) GetSpaceId() string {
	if m != nil {
		return m.SpaceId
	}
	return ""
}

func (m *SpaceReceipt) GetAccountIdentity() []byte {
	if m != nil {
		return m.AccountIdentity
	}
	return nil
}

func (m *SpaceReceipt) GetControlNodeIdentity() []byte {
	if m != nil {
		return m.ControlNodeIdentity
	}
	return nil
}

func (m *SpaceReceipt) GetValidUntil() uint64 {
	if m != nil {
		return m.ValidUntil
	}
	return 0
}

// FileLimitCheckRequest contains an account identity and spaceId
// control node checks that identity owns a given space
type FileLimitCheckRequest struct {
	AccountIdentity []byte `protobuf:"bytes,1,opt,name=accountIdentity,proto3" json:"accountIdentity,omitempty"`
	SpaceId         string `protobuf:"bytes,2,opt,name=spaceId,proto3" json:"spaceId,omitempty"`
}

func (m *FileLimitCheckRequest) Reset()         { *m = FileLimitCheckRequest{} }
func (m *FileLimitCheckRequest) String() string { return proto.CompactTextString(m) }
func (*FileLimitCheckRequest) ProtoMessage()    {}
func (*FileLimitCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d94f6f99586adae2, []int{4}
}
func (m *FileLimitCheckRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FileLimitCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FileLimitCheckRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FileLimitCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileLimitCheckRequest.Merge(m, src)
}
func (m *FileLimitCheckRequest) XXX_Size() int {
	return m.Size()
}
func (m *FileLimitCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FileLimitCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FileLimitCheckRequest proto.InternalMessageInfo

func (m *FileLimitCheckRequest) GetAccountIdentity() []byte {
	if m != nil {
		return m.AccountIdentity
	}
	return nil
}

func (m *FileLimitCheckRequest) GetSpaceId() string {
	if m != nil {
		return m.SpaceId
	}
	return ""
}

// FileLimitCheckResponse returns a current space limit in bytes
type FileLimitCheckResponse struct {
	Limit uint64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (m *FileLimitCheckResponse) Reset()         { *m = FileLimitCheckResponse{} }
func (m *FileLimitCheckResponse) String() string { return proto.CompactTextString(m) }
func (*FileLimitCheckResponse) ProtoMessage()    {}
func (*FileLimitCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d94f6f99586adae2, []int{5}
}
func (m *FileLimitCheckResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FileLimitCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FileLimitCheckResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FileLimitCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileLimitCheckResponse.Merge(m, src)
}
func (m *FileLimitCheckResponse) XXX_Size() int {
	return m.Size()
}
func (m *FileLimitCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FileLimitCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FileLimitCheckResponse proto.InternalMessageInfo

func (m *FileLimitCheckResponse) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func init() {
	proto.RegisterType((*SpaceSignRequest)(nil), "coordinator.SpaceSignRequest")
	proto.RegisterType((*SpaceSignResponse)(nil), "coordinator.SpaceSignResponse")
	proto.RegisterType((*SpaceReceiptWithSignature)(nil), "coordinator.SpaceReceiptWithSignature")
	proto.RegisterType((*SpaceReceipt)(nil), "coordinator.SpaceReceipt")
	proto.RegisterType((*FileLimitCheckRequest)(nil), "coordinator.FileLimitCheckRequest")
	proto.RegisterType((*FileLimitCheckResponse)(nil), "coordinator.FileLimitCheckResponse")
}

func init() {
	proto.RegisterFile("coordinator/coordinatorproto/protos/coordinator.proto", fileDescriptor_d94f6f99586adae2)
}

var fileDescriptor_d94f6f99586adae2 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xdd, 0x4e, 0xc2, 0x30,
	0x18, 0xa5, 0x88, 0x12, 0x3e, 0x88, 0x3f, 0xf5, 0x27, 0x93, 0x60, 0x43, 0x66, 0x62, 0xb8, 0x30,
	0x60, 0x30, 0x7a, 0x6d, 0x24, 0x31, 0xc1, 0x18, 0x63, 0x46, 0x88, 0x51, 0xaf, 0xe6, 0xd6, 0x48,
	0xc3, 0x5c, 0xe7, 0x56, 0x4c, 0x78, 0x0b, 0x1f, 0xc2, 0x47, 0xf0, 0x21, 0xbc, 0xe4, 0xd2, 0x4b,
	0x03, 0x2f, 0x62, 0xd6, 0x31, 0x28, 0x38, 0xb8, 0xd9, 0xd6, 0xf3, 0x9d, 0xf3, 0x7d, 0xa7, 0x3b,
	0x2d, 0x9c, 0x59, 0x9c, 0xfb, 0x36, 0x73, 0x4d, 0xc1, 0xfd, 0x9a, 0xf2, 0xed, 0xf9, 0x5c, 0xf0,
	0x9a, 0x7c, 0x06, 0x2a, 0x5e, 0x95, 0x10, 0xce, 0x2b, 0x90, 0x7e, 0x0c, 0x9b, 0x2d, 0xcf, 0xb4,
	0x68, 0x8b, 0xbd, 0xb8, 0x06, 0x7d, 0xeb, 0xd1, 0x40, 0x60, 0x0d, 0xb2, 0x41, 0x88, 0x35, 0x6d,
	0x0d, 0x95, 0x51, 0x25, 0x67, 0xc4, 0x4b, 0xbd, 0x0d, 0x5b, 0x0a, 0x3b, 0xf0, 0xb8, 0x1b, 0x50,
	0x7c, 0x01, 0x59, 0x9f, 0x5a, 0x94, 0x79, 0x42, 0xd2, 0xf3, 0xf5, 0xa3, 0xaa, 0x3a, 0x54, 0x0a,
	0x8c, 0x88, 0x70, 0xcf, 0x44, 0x27, 0xd4, 0x9a, 0xa2, 0xe7, 0x53, 0x23, 0x96, 0xe9, 0x5d, 0xd8,
	0x5f, 0xc8, 0xc2, 0x27, 0xb0, 0x1d, 0x28, 0xc5, 0x3b, 0xb3, 0xef, 0x70, 0x33, 0x72, 0x56, 0x30,
	0x92, 0x4a, 0xb8, 0x04, 0xb9, 0x20, 0x96, 0x6b, 0x69, 0xc9, 0x9b, 0x02, 0xfa, 0x27, 0x82, 0x82,
	0x3a, 0x6d, 0xf1, 0x76, 0x71, 0x05, 0x36, 0x4c, 0xcb, 0xe2, 0x3d, 0x57, 0x34, 0x6d, 0xea, 0x0a,
	0x26, 0xfa, 0xe3, 0x76, 0xf3, 0x70, 0x68, 0xd2, 0xe2, 0xae, 0xf0, 0xb9, 0x73, 0xcb, 0x6d, 0x3a,
	0x61, 0xaf, 0x44, 0x26, 0x13, 0x4a, 0x98, 0x00, 0xbc, 0x9b, 0x0e, 0xb3, 0xdb, 0xae, 0x60, 0x8e,
	0x96, 0x29, 0xa3, 0x4a, 0xc6, 0x50, 0x10, 0xfd, 0x09, 0x76, 0xaf, 0x98, 0x43, 0x6f, 0xd8, 0x2b,
	0x13, 0x8d, 0x0e, 0xb5, 0xba, 0x71, 0x3a, 0x09, 0xa6, 0x50, 0xb2, 0x29, 0x65, 0x63, 0xe9, 0xd9,
	0x1c, 0xab, 0xb0, 0x37, 0xdf, 0x7c, 0x1c, 0xe6, 0x0e, 0xac, 0x3a, 0x21, 0x2a, 0x7b, 0x66, 0x8c,
	0x68, 0x51, 0xff, 0x42, 0x90, 0x6f, 0x4c, 0x33, 0xc5, 0xd7, 0x90, 0x9b, 0x9c, 0x03, 0x7c, 0xf0,
	0x3f, 0x6e, 0xe5, 0x34, 0x15, 0xc9, 0xa2, 0xf2, 0x78, 0xe2, 0x03, 0xac, 0xcf, 0x7a, 0xc1, 0xfa,
	0x8c, 0x22, 0xf1, 0x2f, 0x14, 0x0f, 0x97, 0x72, 0xa2, 0xd6, 0x97, 0xe7, 0xdf, 0x43, 0x82, 0x06,
	0x43, 0x82, 0x7e, 0x87, 0x04, 0x7d, 0x8c, 0x48, 0x6a, 0x30, 0x22, 0xa9, 0x9f, 0x11, 0x49, 0x3d,
	0x96, 0x96, 0x5d, 0x9d, 0xe7, 0x35, 0xf9, 0x3a, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x99, 0x06,
	0x09, 0x83, 0x61, 0x03, 0x00, 0x00,
}

func (m *SpaceSignRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SpaceSignRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SpaceSignRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SpaceId) > 0 {
		i -= len(m.SpaceId)
		copy(dAtA[i:], m.SpaceId)
		i = encodeVarintCoordinator(dAtA, i, uint64(len(m.SpaceId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SpaceSignResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SpaceSignResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SpaceSignResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Receipt != nil {
		{
			size, err := m.Receipt.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCoordinator(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SpaceReceiptWithSignature) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SpaceReceiptWithSignature) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SpaceReceiptWithSignature) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintCoordinator(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SpaceReceiptPayload) > 0 {
		i -= len(m.SpaceReceiptPayload)
		copy(dAtA[i:], m.SpaceReceiptPayload)
		i = encodeVarintCoordinator(dAtA, i, uint64(len(m.SpaceReceiptPayload)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SpaceReceipt) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SpaceReceipt) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SpaceReceipt) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ValidUntil != 0 {
		i = encodeVarintCoordinator(dAtA, i, uint64(m.ValidUntil))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ControlNodeIdentity) > 0 {
		i -= len(m.ControlNodeIdentity)
		copy(dAtA[i:], m.ControlNodeIdentity)
		i = encodeVarintCoordinator(dAtA, i, uint64(len(m.ControlNodeIdentity)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AccountIdentity) > 0 {
		i -= len(m.AccountIdentity)
		copy(dAtA[i:], m.AccountIdentity)
		i = encodeVarintCoordinator(dAtA, i, uint64(len(m.AccountIdentity)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SpaceId) > 0 {
		i -= len(m.SpaceId)
		copy(dAtA[i:], m.SpaceId)
		i = encodeVarintCoordinator(dAtA, i, uint64(len(m.SpaceId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FileLimitCheckRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FileLimitCheckRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FileLimitCheckRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SpaceId) > 0 {
		i -= len(m.SpaceId)
		copy(dAtA[i:], m.SpaceId)
		i = encodeVarintCoordinator(dAtA, i, uint64(len(m.SpaceId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.AccountIdentity) > 0 {
		i -= len(m.AccountIdentity)
		copy(dAtA[i:], m.AccountIdentity)
		i = encodeVarintCoordinator(dAtA, i, uint64(len(m.AccountIdentity)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FileLimitCheckResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FileLimitCheckResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FileLimitCheckResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Limit != 0 {
		i = encodeVarintCoordinator(dAtA, i, uint64(m.Limit))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCoordinator(dAtA []byte, offset int, v uint64) int {
	offset -= sovCoordinator(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SpaceSignRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SpaceId)
	if l > 0 {
		n += 1 + l + sovCoordinator(uint64(l))
	}
	return n
}

func (m *SpaceSignResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Receipt != nil {
		l = m.Receipt.Size()
		n += 1 + l + sovCoordinator(uint64(l))
	}
	return n
}

func (m *SpaceReceiptWithSignature) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SpaceReceiptPayload)
	if l > 0 {
		n += 1 + l + sovCoordinator(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovCoordinator(uint64(l))
	}
	return n
}

func (m *SpaceReceipt) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SpaceId)
	if l > 0 {
		n += 1 + l + sovCoordinator(uint64(l))
	}
	l = len(m.AccountIdentity)
	if l > 0 {
		n += 1 + l + sovCoordinator(uint64(l))
	}
	l = len(m.ControlNodeIdentity)
	if l > 0 {
		n += 1 + l + sovCoordinator(uint64(l))
	}
	if m.ValidUntil != 0 {
		n += 1 + sovCoordinator(uint64(m.ValidUntil))
	}
	return n
}

func (m *FileLimitCheckRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AccountIdentity)
	if l > 0 {
		n += 1 + l + sovCoordinator(uint64(l))
	}
	l = len(m.SpaceId)
	if l > 0 {
		n += 1 + l + sovCoordinator(uint64(l))
	}
	return n
}

func (m *FileLimitCheckResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Limit != 0 {
		n += 1 + sovCoordinator(uint64(m.Limit))
	}
	return n
}

func sovCoordinator(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCoordinator(x uint64) (n int) {
	return sovCoordinator(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SpaceSignRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoordinator
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
			return fmt.Errorf("proto: SpaceSignRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SpaceSignRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpaceId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoordinator
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
				return ErrInvalidLengthCoordinator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoordinator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpaceId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCoordinator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCoordinator
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
func (m *SpaceSignResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoordinator
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
			return fmt.Errorf("proto: SpaceSignResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SpaceSignResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receipt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoordinator
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
				return ErrInvalidLengthCoordinator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCoordinator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Receipt == nil {
				m.Receipt = &SpaceReceiptWithSignature{}
			}
			if err := m.Receipt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCoordinator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCoordinator
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
func (m *SpaceReceiptWithSignature) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoordinator
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
			return fmt.Errorf("proto: SpaceReceiptWithSignature: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SpaceReceiptWithSignature: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpaceReceiptPayload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoordinator
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
				return ErrInvalidLengthCoordinator
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCoordinator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpaceReceiptPayload = append(m.SpaceReceiptPayload[:0], dAtA[iNdEx:postIndex]...)
			if m.SpaceReceiptPayload == nil {
				m.SpaceReceiptPayload = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoordinator
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
				return ErrInvalidLengthCoordinator
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCoordinator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCoordinator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCoordinator
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
func (m *SpaceReceipt) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoordinator
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
			return fmt.Errorf("proto: SpaceReceipt: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SpaceReceipt: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpaceId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoordinator
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
				return ErrInvalidLengthCoordinator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoordinator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpaceId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountIdentity", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoordinator
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
				return ErrInvalidLengthCoordinator
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCoordinator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccountIdentity = append(m.AccountIdentity[:0], dAtA[iNdEx:postIndex]...)
			if m.AccountIdentity == nil {
				m.AccountIdentity = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ControlNodeIdentity", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoordinator
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
				return ErrInvalidLengthCoordinator
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCoordinator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ControlNodeIdentity = append(m.ControlNodeIdentity[:0], dAtA[iNdEx:postIndex]...)
			if m.ControlNodeIdentity == nil {
				m.ControlNodeIdentity = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidUntil", wireType)
			}
			m.ValidUntil = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoordinator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValidUntil |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCoordinator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCoordinator
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
func (m *FileLimitCheckRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoordinator
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
			return fmt.Errorf("proto: FileLimitCheckRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FileLimitCheckRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountIdentity", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoordinator
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
				return ErrInvalidLengthCoordinator
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCoordinator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccountIdentity = append(m.AccountIdentity[:0], dAtA[iNdEx:postIndex]...)
			if m.AccountIdentity == nil {
				m.AccountIdentity = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpaceId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoordinator
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
				return ErrInvalidLengthCoordinator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoordinator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpaceId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCoordinator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCoordinator
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
func (m *FileLimitCheckResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoordinator
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
			return fmt.Errorf("proto: FileLimitCheckResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FileLimitCheckResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoordinator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCoordinator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCoordinator
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
func skipCoordinator(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCoordinator
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
					return 0, ErrIntOverflowCoordinator
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
					return 0, ErrIntOverflowCoordinator
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
				return 0, ErrInvalidLengthCoordinator
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCoordinator
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCoordinator
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCoordinator        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCoordinator          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCoordinator = fmt.Errorf("proto: unexpected end of group")
)