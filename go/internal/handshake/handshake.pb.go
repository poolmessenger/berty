// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: internal/handshake.proto

package handshake

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	crypto "berty.tech/go/internal/crypto"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type HandshakeFrame_HandshakeStep int32

const (
	HandshakeFrame_STEP_1_KEY_AGREEMENT              HandshakeFrame_HandshakeStep = 0
	HandshakeFrame_STEP_2_KEY_AGREEMENT              HandshakeFrame_HandshakeStep = 1
	HandshakeFrame_STEP_3_DISPATCH                   HandshakeFrame_HandshakeStep = 10
	HandshakeFrame_STEP_3A_KNOWN_IDENTITY_PROOF      HandshakeFrame_HandshakeStep = 20
	HandshakeFrame_STEP_4A_KNOWN_IDENTITY_DISCLOSURE HandshakeFrame_HandshakeStep = 21
	HandshakeFrame_STEP_5A_KNOWN_IDENTITY_DISCLOSURE HandshakeFrame_HandshakeStep = 22
	HandshakeFrame_STEP_3B_KNOWN_DEVICE_PROOF        HandshakeFrame_HandshakeStep = 30
	HandshakeFrame_STEP_4B_KNOWN_DEVICE_DISCLOSURE   HandshakeFrame_HandshakeStep = 31
	HandshakeFrame_STEP_9_DONE                       HandshakeFrame_HandshakeStep = 999
)

var HandshakeFrame_HandshakeStep_name = map[int32]string{
	0:   "STEP_1_KEY_AGREEMENT",
	1:   "STEP_2_KEY_AGREEMENT",
	10:  "STEP_3_DISPATCH",
	20:  "STEP_3A_KNOWN_IDENTITY_PROOF",
	21:  "STEP_4A_KNOWN_IDENTITY_DISCLOSURE",
	22:  "STEP_5A_KNOWN_IDENTITY_DISCLOSURE",
	30:  "STEP_3B_KNOWN_DEVICE_PROOF",
	31:  "STEP_4B_KNOWN_DEVICE_DISCLOSURE",
	999: "STEP_9_DONE",
}

var HandshakeFrame_HandshakeStep_value = map[string]int32{
	"STEP_1_KEY_AGREEMENT":              0,
	"STEP_2_KEY_AGREEMENT":              1,
	"STEP_3_DISPATCH":                   10,
	"STEP_3A_KNOWN_IDENTITY_PROOF":      20,
	"STEP_4A_KNOWN_IDENTITY_DISCLOSURE": 21,
	"STEP_5A_KNOWN_IDENTITY_DISCLOSURE": 22,
	"STEP_3B_KNOWN_DEVICE_PROOF":        30,
	"STEP_4B_KNOWN_DEVICE_DISCLOSURE":   31,
	"STEP_9_DONE":                       999,
}

func (x HandshakeFrame_HandshakeStep) String() string {
	return proto.EnumName(HandshakeFrame_HandshakeStep_name, int32(x))
}

func (HandshakeFrame_HandshakeStep) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_15a58250da931098, []int{0, 0}
}

type HandshakeFrame struct {
	Step             HandshakeFrame_HandshakeStep `protobuf:"varint,1,opt,name=step,proto3,enum=handshake.HandshakeFrame_HandshakeStep" json:"step,omitempty"`
	SignatureKey     []byte                       `protobuf:"bytes,2,opt,name=signatureKey,proto3" json:"signatureKey,omitempty"`
	EncryptionKey    []byte                       `protobuf:"bytes,3,opt,name=encryptionKey,proto3" json:"encryptionKey,omitempty"`
	EncryptedPayload []byte                       `protobuf:"bytes,4,opt,name=encryptedPayload,proto3" json:"encryptedPayload,omitempty"`
}

func (m *HandshakeFrame) Reset()         { *m = HandshakeFrame{} }
func (m *HandshakeFrame) String() string { return proto.CompactTextString(m) }
func (*HandshakeFrame) ProtoMessage()    {}
func (*HandshakeFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a58250da931098, []int{0}
}
func (m *HandshakeFrame) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HandshakeFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HandshakeFrame.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HandshakeFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandshakeFrame.Merge(m, src)
}
func (m *HandshakeFrame) XXX_Size() int {
	return m.Size()
}
func (m *HandshakeFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_HandshakeFrame.DiscardUnknown(m)
}

var xxx_messageInfo_HandshakeFrame proto.InternalMessageInfo

func (m *HandshakeFrame) GetStep() HandshakeFrame_HandshakeStep {
	if m != nil {
		return m.Step
	}
	return HandshakeFrame_STEP_1_KEY_AGREEMENT
}

func (m *HandshakeFrame) GetSignatureKey() []byte {
	if m != nil {
		return m.SignatureKey
	}
	return nil
}

func (m *HandshakeFrame) GetEncryptionKey() []byte {
	if m != nil {
		return m.EncryptionKey
	}
	return nil
}

func (m *HandshakeFrame) GetEncryptedPayload() []byte {
	if m != nil {
		return m.EncryptedPayload
	}
	return nil
}

type HandshakePayload struct {
	Signature []byte           `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	SigChain  *crypto.SigChain `protobuf:"bytes,2,opt,name=sigChain,proto3" json:"sigChain,omitempty"`
	DeviceKey []byte           `protobuf:"bytes,3,opt,name=deviceKey,proto3" json:"deviceKey,omitempty"`
}

func (m *HandshakePayload) Reset()         { *m = HandshakePayload{} }
func (m *HandshakePayload) String() string { return proto.CompactTextString(m) }
func (*HandshakePayload) ProtoMessage()    {}
func (*HandshakePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a58250da931098, []int{1}
}
func (m *HandshakePayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HandshakePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HandshakePayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HandshakePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandshakePayload.Merge(m, src)
}
func (m *HandshakePayload) XXX_Size() int {
	return m.Size()
}
func (m *HandshakePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_HandshakePayload.DiscardUnknown(m)
}

var xxx_messageInfo_HandshakePayload proto.InternalMessageInfo

func (m *HandshakePayload) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *HandshakePayload) GetSigChain() *crypto.SigChain {
	if m != nil {
		return m.SigChain
	}
	return nil
}

func (m *HandshakePayload) GetDeviceKey() []byte {
	if m != nil {
		return m.DeviceKey
	}
	return nil
}

func init() {
	proto.RegisterEnum("handshake.HandshakeFrame_HandshakeStep", HandshakeFrame_HandshakeStep_name, HandshakeFrame_HandshakeStep_value)
	proto.RegisterType((*HandshakeFrame)(nil), "handshake.HandshakeFrame")
	proto.RegisterType((*HandshakePayload)(nil), "handshake.HandshakePayload")
}

func init() { proto.RegisterFile("internal/handshake.proto", fileDescriptor_15a58250da931098) }

var fileDescriptor_15a58250da931098 = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xe3, 0xb6, 0x02, 0x32, 0x4d, 0x8b, 0xb5, 0x14, 0xb0, 0xa2, 0xca, 0x0d, 0x81, 0x8a,
	0x0a, 0x09, 0x5b, 0xa4, 0x70, 0x00, 0x4e, 0x69, 0xbc, 0xa5, 0x56, 0xc0, 0x8e, 0x6c, 0x03, 0x2a,
	0x17, 0xcb, 0x71, 0x16, 0x7b, 0x45, 0xeb, 0x8d, 0xec, 0x0d, 0x52, 0xde, 0x82, 0x37, 0xe0, 0x75,
	0x38, 0xf6, 0xc8, 0x11, 0x25, 0x42, 0xbc, 0x06, 0xca, 0xda, 0xd8, 0x24, 0x91, 0x7a, 0xf3, 0x7e,
	0xff, 0xe7, 0x99, 0xd5, 0xce, 0x80, 0x42, 0x13, 0x4e, 0xd2, 0x24, 0xb8, 0xd0, 0xe3, 0x20, 0x19,
	0x65, 0x71, 0xf0, 0x85, 0x68, 0xe3, 0x94, 0x71, 0x86, 0xea, 0x25, 0x68, 0xde, 0x2f, 0xa5, 0x8c,
	0x46, 0x61, 0x1c, 0xd0, 0x24, 0x77, 0x9a, 0x4f, 0x23, 0xca, 0xe3, 0xc9, 0x50, 0x0b, 0xd9, 0xa5,
	0x1e, 0xb1, 0x88, 0xe9, 0x02, 0x0f, 0x27, 0x9f, 0xc5, 0x49, 0x1c, 0xc4, 0x57, 0xae, 0xb7, 0x7f,
	0x6f, 0xc2, 0xee, 0xd9, 0xbf, 0xaa, 0xa7, 0x69, 0x70, 0x49, 0xd0, 0x6b, 0xd8, 0xca, 0x38, 0x19,
	0x2b, 0x52, 0x4b, 0x3a, 0xda, 0xed, 0x3c, 0xd6, 0xaa, 0x5b, 0x2c, 0x8b, 0xd5, 0xd1, 0xe5, 0x64,
	0xec, 0x88, 0x9f, 0x50, 0x1b, 0x1a, 0x19, 0x8d, 0x92, 0x80, 0x4f, 0x52, 0xd2, 0x27, 0x53, 0x65,
	0xa3, 0x25, 0x1d, 0x35, 0x9c, 0x25, 0x86, 0x1e, 0xc1, 0x0e, 0x49, 0xc2, 0x74, 0x3a, 0xe6, 0x94,
	0x25, 0x0b, 0x69, 0x53, 0x48, 0xcb, 0x10, 0x3d, 0x01, 0xb9, 0x00, 0x64, 0x34, 0x08, 0xa6, 0x17,
	0x2c, 0x18, 0x29, 0x5b, 0x42, 0x5c, 0xe3, 0xed, 0xef, 0x1b, 0xb0, 0xb3, 0x74, 0x1b, 0xa4, 0xc0,
	0x9e, 0xeb, 0xe1, 0x81, 0xff, 0xcc, 0xef, 0xe3, 0x73, 0xbf, 0xfb, 0xc6, 0xc1, 0xf8, 0x1d, 0xb6,
	0x3c, 0xb9, 0x56, 0x26, 0x9d, 0x95, 0x44, 0x42, 0x77, 0xe0, 0xb6, 0x48, 0x8e, 0x7d, 0xc3, 0x74,
	0x07, 0x5d, 0xaf, 0x77, 0x26, 0x03, 0x6a, 0xc1, 0x7e, 0x0e, 0xbb, 0x7e, 0xdf, 0xb2, 0x3f, 0x5a,
	0xbe, 0x69, 0x60, 0xcb, 0x33, 0xbd, 0x73, 0x7f, 0xe0, 0xd8, 0xf6, 0xa9, 0xbc, 0x87, 0x0e, 0xe1,
	0x81, 0x30, 0x9e, 0xaf, 0x19, 0x86, 0xe9, 0xf6, 0xde, 0xda, 0xee, 0x7b, 0x07, 0xcb, 0x77, 0x4b,
	0xed, 0xc5, 0x75, 0xda, 0x3d, 0xa4, 0x42, 0x33, 0xef, 0x77, 0x52, 0x68, 0x06, 0xfe, 0x60, 0xf6,
	0x70, 0xd1, 0x4d, 0x45, 0x0f, 0xe1, 0x20, 0xef, 0xb6, 0x92, 0xff, 0x57, 0xe4, 0x00, 0xc9, 0xb0,
	0x2d, 0xa4, 0x97, 0xbe, 0x61, 0x5b, 0x58, 0xfe, 0x73, 0xb3, 0x3d, 0x01, 0xb9, 0x7c, 0xa0, 0xe2,
	0xd5, 0xd0, 0x3e, 0xd4, 0xcb, 0xb9, 0x88, 0x69, 0x37, 0x9c, 0x0a, 0xa0, 0x43, 0xb8, 0x95, 0xd1,
	0xa8, 0xb7, 0x58, 0x2d, 0x31, 0xc5, 0xed, 0x4e, 0x5d, 0x73, 0x0b, 0xe0, 0x94, 0xd1, 0xa2, 0xc8,
	0x88, 0x7c, 0xa5, 0x21, 0xa9, 0x06, 0x59, 0x81, 0x93, 0x57, 0x3f, 0x66, 0xaa, 0x74, 0x35, 0x53,
	0xa5, 0x5f, 0x33, 0x55, 0xfa, 0x36, 0x57, 0x6b, 0x57, 0x73, 0xb5, 0xf6, 0x73, 0xae, 0xd6, 0x3e,
	0xb5, 0x86, 0x24, 0xe5, 0x53, 0x8d, 0x93, 0x30, 0xd6, 0x23, 0xa6, 0xaf, 0xef, 0xfc, 0xf0, 0x86,
	0xd8, 0xd0, 0xe3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x49, 0x11, 0x20, 0x10, 0x03, 0x00,
	0x00,
}

func (m *HandshakeFrame) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HandshakeFrame) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HandshakeFrame) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EncryptedPayload) > 0 {
		i -= len(m.EncryptedPayload)
		copy(dAtA[i:], m.EncryptedPayload)
		i = encodeVarintHandshake(dAtA, i, uint64(len(m.EncryptedPayload)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.EncryptionKey) > 0 {
		i -= len(m.EncryptionKey)
		copy(dAtA[i:], m.EncryptionKey)
		i = encodeVarintHandshake(dAtA, i, uint64(len(m.EncryptionKey)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SignatureKey) > 0 {
		i -= len(m.SignatureKey)
		copy(dAtA[i:], m.SignatureKey)
		i = encodeVarintHandshake(dAtA, i, uint64(len(m.SignatureKey)))
		i--
		dAtA[i] = 0x12
	}
	if m.Step != 0 {
		i = encodeVarintHandshake(dAtA, i, uint64(m.Step))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *HandshakePayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HandshakePayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HandshakePayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DeviceKey) > 0 {
		i -= len(m.DeviceKey)
		copy(dAtA[i:], m.DeviceKey)
		i = encodeVarintHandshake(dAtA, i, uint64(len(m.DeviceKey)))
		i--
		dAtA[i] = 0x1a
	}
	if m.SigChain != nil {
		{
			size, err := m.SigChain.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHandshake(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintHandshake(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintHandshake(dAtA []byte, offset int, v uint64) int {
	offset -= sovHandshake(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HandshakeFrame) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Step != 0 {
		n += 1 + sovHandshake(uint64(m.Step))
	}
	l = len(m.SignatureKey)
	if l > 0 {
		n += 1 + l + sovHandshake(uint64(l))
	}
	l = len(m.EncryptionKey)
	if l > 0 {
		n += 1 + l + sovHandshake(uint64(l))
	}
	l = len(m.EncryptedPayload)
	if l > 0 {
		n += 1 + l + sovHandshake(uint64(l))
	}
	return n
}

func (m *HandshakePayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovHandshake(uint64(l))
	}
	if m.SigChain != nil {
		l = m.SigChain.Size()
		n += 1 + l + sovHandshake(uint64(l))
	}
	l = len(m.DeviceKey)
	if l > 0 {
		n += 1 + l + sovHandshake(uint64(l))
	}
	return n
}

func sovHandshake(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHandshake(x uint64) (n int) {
	return sovHandshake(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HandshakeFrame) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHandshake
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
			return fmt.Errorf("proto: HandshakeFrame: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HandshakeFrame: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Step", wireType)
			}
			m.Step = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandshake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Step |= HandshakeFrame_HandshakeStep(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignatureKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandshake
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
				return ErrInvalidLengthHandshake
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHandshake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SignatureKey = append(m.SignatureKey[:0], dAtA[iNdEx:postIndex]...)
			if m.SignatureKey == nil {
				m.SignatureKey = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncryptionKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandshake
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
				return ErrInvalidLengthHandshake
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHandshake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EncryptionKey = append(m.EncryptionKey[:0], dAtA[iNdEx:postIndex]...)
			if m.EncryptionKey == nil {
				m.EncryptionKey = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncryptedPayload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandshake
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
				return ErrInvalidLengthHandshake
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHandshake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EncryptedPayload = append(m.EncryptedPayload[:0], dAtA[iNdEx:postIndex]...)
			if m.EncryptedPayload == nil {
				m.EncryptedPayload = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHandshake(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHandshake
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHandshake
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
func (m *HandshakePayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHandshake
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
			return fmt.Errorf("proto: HandshakePayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HandshakePayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandshake
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
				return ErrInvalidLengthHandshake
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHandshake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SigChain", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandshake
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
				return ErrInvalidLengthHandshake
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHandshake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SigChain == nil {
				m.SigChain = &crypto.SigChain{}
			}
			if err := m.SigChain.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandshake
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
				return ErrInvalidLengthHandshake
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHandshake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeviceKey = append(m.DeviceKey[:0], dAtA[iNdEx:postIndex]...)
			if m.DeviceKey == nil {
				m.DeviceKey = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHandshake(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHandshake
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHandshake
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
func skipHandshake(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHandshake
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
					return 0, ErrIntOverflowHandshake
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHandshake
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
				return 0, ErrInvalidLengthHandshake
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthHandshake
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHandshake
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipHandshake(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthHandshake
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthHandshake = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHandshake   = fmt.Errorf("proto: integer overflow")
)
