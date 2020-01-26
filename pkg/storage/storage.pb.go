// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage.proto

package storage

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type SpeedInfo struct {
	Time                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Number               string               `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	Speed                float32              `protobuf:"fixed32,3,opt,name=speed,proto3" json:"speed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SpeedInfo) Reset()         { *m = SpeedInfo{} }
func (m *SpeedInfo) String() string { return proto.CompactTextString(m) }
func (*SpeedInfo) ProtoMessage()    {}
func (*SpeedInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{0}
}

func (m *SpeedInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpeedInfo.Unmarshal(m, b)
}
func (m *SpeedInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpeedInfo.Marshal(b, m, deterministic)
}
func (m *SpeedInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpeedInfo.Merge(m, src)
}
func (m *SpeedInfo) XXX_Size() int {
	return xxx_messageInfo_SpeedInfo.Size(m)
}
func (m *SpeedInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SpeedInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SpeedInfo proto.InternalMessageInfo

func (m *SpeedInfo) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *SpeedInfo) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *SpeedInfo) GetSpeed() float32 {
	if m != nil {
		return m.Speed
	}
	return 0
}

func init() {
	proto.RegisterType((*SpeedInfo)(nil), "storage.SpeedInfo")
}

func init() { proto.RegisterFile("storage.proto", fileDescriptor_0d2c4ccf1453ffdb) }

var fileDescriptor_0d2c4ccf1453ffdb = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x2e, 0xc9, 0x2f,
	0x4a, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0xa5, 0xe4, 0xd3,
	0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0xc1, 0xc2, 0x49, 0xa5, 0x69, 0xfa, 0x25, 0x99, 0xb9, 0xa9,
	0xc5, 0x25, 0x89, 0xb9, 0x05, 0x10, 0x95, 0x4a, 0x99, 0x5c, 0x9c, 0xc1, 0x05, 0xa9, 0xa9, 0x29,
	0x9e, 0x79, 0x69, 0xf9, 0x42, 0x7a, 0x5c, 0x2c, 0x20, 0x79, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x6e,
	0x23, 0x29, 0x3d, 0x88, 0x66, 0x3d, 0x98, 0x66, 0xbd, 0x10, 0x98, 0xe6, 0x20, 0xb0, 0x3a, 0x21,
	0x31, 0x2e, 0xb6, 0xbc, 0xd2, 0xdc, 0xa4, 0xd4, 0x22, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20,
	0x28, 0x4f, 0x48, 0x84, 0x8b, 0xb5, 0x18, 0x64, 0xa8, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x53, 0x10,
	0x84, 0x93, 0xc4, 0x06, 0x36, 0xc7, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xad, 0xf5, 0x33, 0xf6,
	0xac, 0x00, 0x00, 0x00,
}
