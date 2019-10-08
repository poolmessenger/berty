// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protocolmodel.proto

package protocolmodel

import (
	fmt "fmt"
	math "math"
	reflect "reflect"
	strings "strings"
	time "time"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Contact struct {
	ID        string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primary_key"`
	CreatedAt *time.Time `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at,omitempty"`
	UpdatedAt *time.Time `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3,stdtime" json:"updated_at,omitempty"`
}

func (m *Contact) Reset()      { *m = Contact{} }
func (*Contact) ProtoMessage() {}
func (*Contact) Descriptor() ([]byte, []int) {
	return fileDescriptor_cff888de946c583b, []int{0}
}
func (m *Contact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Contact.Unmarshal(m, b)
}
func (m *Contact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Contact.Marshal(b, m, deterministic)
}
func (m *Contact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contact.Merge(m, src)
}
func (m *Contact) XXX_Size() int {
	return xxx_messageInfo_Contact.Size(m)
}
func (m *Contact) XXX_DiscardUnknown() {
	xxx_messageInfo_Contact.DiscardUnknown(m)
}

var xxx_messageInfo_Contact proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Contact)(nil), "berty.protocolmodel.Contact")
}

func init() { proto.RegisterFile("protocolmodel.proto", fileDescriptor_cff888de946c583b) }

var fileDescriptor_cff888de946c583b = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x2f, 0x6c, 0xdb, 0x40,
	0x14, 0xc6, 0xdf, 0x79, 0xd2, 0xb6, 0x7a, 0xcc, 0x25, 0x55, 0x34, 0x7d, 0xae, 0x82, 0x0a, 0x36,
	0x5b, 0xda, 0xd8, 0x86, 0x9a, 0x8e, 0x8c, 0x56, 0x43, 0x23, 0x95, 0xff, 0xdc, 0x1c, 0x2b, 0x76,
	0xce, 0x72, 0x2e, 0xc0, 0x2c, 0x30, 0x30, 0x70, 0x70, 0x63, 0x81, 0x81, 0x81, 0x19, 0x33, 0x0c,
	0x0c, 0xca, 0xe2, 0x3b, 0x12, 0x18, 0x0d, 0x05, 0x4e, 0x71, 0x62, 0x45, 0x81, 0x65, 0xef, 0xd3,
	0xbd, 0xdf, 0xef, 0x4e, 0xf7, 0x99, 0xd7, 0x59, 0x2e, 0xa4, 0x08, 0x44, 0x92, 0x8a, 0x90, 0x27,
	0x4e, 0x9d, 0xac, 0x6b, 0x9f, 0xe7, 0xb2, 0x70, 0x2e, 0x8e, 0x5a, 0x9f, 0xa3, 0x58, 0x76, 0x87,
	0xbe, 0x13, 0x88, 0xd4, 0x8d, 0x44, 0xe2, 0xf5, 0x23, 0xb7, 0x5e, 0xf0, 0x87, 0x3f, 0xdc, 0x4c,
	0x16, 0x19, 0x1f, 0xb8, 0x32, 0x4e, 0xf9, 0x40, 0x7a, 0x69, 0x76, 0x9e, 0x8e, 0x92, 0xd6, 0xfb,
	0x0b, 0x38, 0x12, 0x67, 0xf4, 0x90, 0xea, 0x50, 0x4f, 0xc7, 0xf5, 0xf6, 0x1f, 0x66, 0xbe, 0x7a,
	0x10, 0x7d, 0xe9, 0x05, 0xd2, 0x7a, 0x67, 0x1a, 0x71, 0x78, 0xc3, 0x6e, 0xd9, 0xdd, 0x55, 0xe7,
	0xad, 0x5a, 0xdb, 0xc6, 0xd7, 0x2f, 0xff, 0xd6, 0xb6, 0x15, 0x89, 0x3c, 0xfd, 0xd4, 0xce, 0xf2,
	0x38, 0xf5, 0xf2, 0xe2, 0xa9, 0xc7, 0x8b, 0xf6, 0xa3, 0x11, 0x87, 0xd6, 0x83, 0x69, 0x06, 0x39,
	0xf7, 0x24, 0x0f, 0x9f, 0x3c, 0x79, 0x63, 0xdc, 0xb2, 0xbb, 0x37, 0x1f, 0x5a, 0x4e, 0x24, 0x44,
	0x94, 0x70, 0xa7, 0xb9, 0xd4, 0xf9, 0xd6, 0x3c, 0xaf, 0xf3, 0xba, 0x5c, 0xdb, 0x6c, 0xf2, 0xd7,
	0x66, 0x8f, 0x57, 0x27, 0xee, 0x5e, 0x1e, 0x24, 0xc3, 0x2c, 0x6c, 0x24, 0x2f, 0x9e, 0x23, 0x39,
	0x71, 0xf7, 0xb2, 0xf3, 0x9b, 0x95, 0x15, 0x68, 0x59, 0x81, 0xad, 0x2a, 0xd0, 0xa6, 0x02, 0x6d,
	0x2b, 0xd0, 0xae, 0x02, 0xed, 0x2b, 0xd0, 0x48, 0x81, 0x8d, 0x15, 0x68, 0xaa, 0x40, 0x33, 0x05,
	0x9a, 0x2b, 0xd0, 0x42, 0x81, 0x4a, 0x05, 0x5a, 0x2a, 0xd0, 0x4a, 0x81, 0x36, 0x0a, 0xb4, 0x55,
	0xa0, 0x9d, 0x02, 0xed, 0x15, 0x68, 0xa4, 0x41, 0x63, 0x0d, 0x9a, 0x68, 0xd0, 0x4f, 0x0d, 0xfa,
	0xa5, 0x41, 0x53, 0x0d, 0x9a, 0x69, 0xb0, 0xb9, 0x06, 0x5b, 0x68, 0x50, 0xa9, 0x41, 0x4b, 0x0d,
	0x5a, 0x69, 0xd0, 0x77, 0xfb, 0xd8, 0xa3, 0xe4, 0x41, 0xd7, 0x3d, 0xfc, 0x79, 0xef, 0x54, 0x59,
	0xd3, 0xa9, 0xff, 0xb2, 0x8e, 0x1f, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x1d, 0xf5, 0xf0,
	0x06, 0x02, 0x00, 0x00,
}

func (this *Contact) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Contact{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`CreatedAt:` + strings.Replace(fmt.Sprintf("%v", this.CreatedAt), "Timestamp", "timestamp.Timestamp", 1) + `,`,
		`UpdatedAt:` + strings.Replace(fmt.Sprintf("%v", this.UpdatedAt), "Timestamp", "timestamp.Timestamp", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringProtocolmodel(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}