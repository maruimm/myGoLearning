// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package myProto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type FOO int32

const (
	FOO_X FOO = 17
)

var FOO_name = map[int32]string{
	17: "X",
}

var FOO_value = map[string]int32{
	"X": 17,
}

func (x FOO) Enum() *FOO {
	p := new(FOO)
	*p = x
	return p
}

func (x FOO) String() string {
	return proto.EnumName(FOO_name, int32(x))
}

func (x *FOO) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FOO_value, data, "FOO")
	if err != nil {
		return err
	}
	*x = FOO(value)
	return nil
}

func (FOO) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

type Test struct {
	Label                *string  `protobuf:"bytes,1,req,name=label" json:"label,omitempty"`
	Type                 *int32   `protobuf:"varint,2,opt,name=type,def=77" json:"type,omitempty"`
	Reps                 []int64  `protobuf:"varint,3,rep,name=reps" json:"reps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Test) Reset()         { *m = Test{} }
func (m *Test) String() string { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()    {}
func (*Test) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

func (m *Test) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test.Unmarshal(m, b)
}
func (m *Test) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test.Marshal(b, m, deterministic)
}
func (m *Test) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test.Merge(m, src)
}
func (m *Test) XXX_Size() int {
	return xxx_messageInfo_Test.Size(m)
}
func (m *Test) XXX_DiscardUnknown() {
	xxx_messageInfo_Test.DiscardUnknown(m)
}

var xxx_messageInfo_Test proto.InternalMessageInfo

const Default_Test_Type int32 = 77

func (m *Test) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *Test) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_Test_Type
}

func (m *Test) GetReps() []int64 {
	if m != nil {
		return m.Reps
	}
	return nil
}

func init() {
	proto.RegisterEnum("myProto.FOO", FOO_name, FOO_value)
	proto.RegisterType((*Test)(nil), "myProto.Test")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 120 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcf, 0xad, 0x0c, 0x00, 0x31, 0x94, 0x3c, 0xb8,
	0x58, 0x42, 0x52, 0x8b, 0x4b, 0x84, 0x44, 0xb8, 0x58, 0x73, 0x12, 0x93, 0x52, 0x73, 0x24, 0x18,
	0x15, 0x98, 0x34, 0x38, 0x83, 0x20, 0x1c, 0x21, 0x31, 0x2e, 0x96, 0x92, 0xca, 0x82, 0x54, 0x09,
	0x26, 0x05, 0x46, 0x0d, 0x56, 0x2b, 0x26, 0x73, 0xf3, 0x20, 0x30, 0x5f, 0x48, 0x88, 0x8b, 0xa5,
	0x28, 0xb5, 0xa0, 0x58, 0x82, 0x59, 0x81, 0x59, 0x83, 0x39, 0x08, 0xcc, 0xd6, 0xe2, 0xe1, 0x62,
	0x76, 0xf3, 0xf7, 0x17, 0x62, 0xe5, 0x62, 0x8c, 0x10, 0x10, 0x04, 0x04, 0x00, 0x00, 0xff, 0xff,
	0xca, 0x25, 0x09, 0x82, 0x6d, 0x00, 0x00, 0x00,
}
